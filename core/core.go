package core

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.trulyao.dev/lito/core/api/handlers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
	"golang.org/x/sync/errgroup"
)

type Core struct {
	debug          bool
	config         *types.Config
	storageHandler storage.Storage
	logHandler     logger.Logger
	errorHandler   types.ErrorHandler
}

type Opts struct {
	Debug        bool
	Config       *types.Config
	LogHandler   logger.Logger
	ErrorHandler types.ErrorHandler
}

func New(opts *Opts) *Core {
	if opts == nil {
		panic("opts cannot be nil")
	}

	errorHandler := handlers.ErrorHandler
	if opts.ErrorHandler != nil {
		errorHandler = opts.ErrorHandler
	}

	var logHandler logger.Logger
	logHandler = logger.DefaultLogHandler
	if opts.LogHandler != nil {
		logHandler = opts.LogHandler
	}

	var storageHandler storage.Storage
	if !opts.Config.Proxy.IsNone() {
		if !opts.Config.Proxy.Unwrap().Storage.IsNone() {
			logHandler.Info("loading storage handler", logger.Field("type", opts.Config.Proxy.Unwrap().Storage.Unwrap()))
			storageHandler, _ = storage.New(&storage.Opts{
				Config:     opts.Config,
				LogHandler: logHandler,
			})
		}
	}

	return &Core{
		debug:          opts.Debug,
		config:         opts.Config,
		storageHandler: storageHandler,
		logHandler:     logHandler,
		errorHandler:   errorHandler,
	}
}

// HandleShutdown handles the shutdown signal and gracefully shutdown the proxy while maintaining the current config and state
// You should call this before you start the proxy - it will block until the shutdown signal is received
func (c *Core) HandleShutdown(sig chan os.Signal) {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	c.logHandler.Info("shutting down proxy")
	c.stopProxy()
	c.logHandler.Info("proxy shutdown complete")

	if err := c.storageHandler.Persist(); err != nil {
		c.logHandler.Error("failed to persist config", logger.Field("error", err))
	}

	if err := c.logHandler.Sync(); err != nil {
		c.logHandler.Error("failed to sync log handler", logger.Field("error", err))
	}
}

func (c *Core) Run() error {
	sig := make(chan os.Signal, 1)
	go func() {
		c.HandleShutdown(sig)
	}()

	if err := c.storageHandler.Load(); err != nil {
		c.logHandler.Error("failed to load config", logger.Field("error", err))
		return err
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		c.logHandler.Info("starting proxy")
		if err := c.startProxy(); err != nil {
			if errors.Is(http.ErrServerClosed, err) {
				return nil
			}

			return err
		}

		return nil
	})

	adminApiEnabled := c.config.Admin.IsSome() && c.config.Admin.Unwrap().Enabled.IsSome() && c.config.Admin.Unwrap().Enabled.Unwrap() == true
	if !adminApiEnabled && c.storageHandler.IsWatchchable() {
		eg.Go(func() error {
			c.watchConfig(sig)
			return nil
		})
	}

	// TODO: run the admin API

	if err := eg.Wait(); err != nil {
		c.logHandler.Error("Something went horribly wrong", logger.Field("error", err.Error()))
		return err
	}

	return nil
}
