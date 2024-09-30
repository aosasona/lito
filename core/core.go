package core

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"go.trulyao.dev/lito/core/api/handlers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
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

	var (
		storageHandler storage.Storage
		err            error
	)
	if opts.Config.Proxy != nil && opts.Config.Proxy.Storage != nil {
		logHandler.Info("loading storage handler", logger.Field("type", opts.Config.Proxy.Storage))

		if storageHandler, err = storage.New(&storage.Opts{Config: opts.Config, LogHandler: logHandler}); err != nil {
			logHandler.Error("failed to load storage handler", logger.Field("error", err))
			os.Exit(1)
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
	_ = c.stopProxy()
	c.logHandler.Info("proxy shutdown complete")

	if err := c.storageHandler.Persist(); err != nil {
		c.logHandler.Error("failed to persist config", logger.Field("error", err))
	}

	if err := c.logHandler.Sync(); err != nil {
		c.logHandler.Error("failed to sync log handler", logger.Field("error", err))
	}
}

func (c *Core) performSanityCheck() error {
	if c.config.Proxy != nil {
		if ref.Deref(c.config.Proxy.EnableTLS, false) &&
			ref.Deref(c.config.Proxy.TLSEmail, "") == "" {
			return errors.New("TLS email is not set but TLS is enabled!")
		}
	}

	return nil
}

func (c *Core) Run() error {
	// Sanity check
	c.ensureAllFielsdSet()

	sig := make(chan os.Signal, 1)
	go func() {
		c.HandleShutdown(sig)
	}()

	if err := c.storageHandler.Load(); err != nil {
		c.logHandler.Error("failed to load config", logger.Field("error", err))
		return err
	}

	if err := c.performSanityCheck(); err != nil {
		return err
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		if c.config.Proxy == nil {
			return errors.New("no proxy config present")
		}

		c.logHandler.Info(
			"starting proxy server",
			logger.Field("port", ref.Deref(c.config.Proxy.HTTPPort, 80)),
		)
		if err := c.startProxy(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			return err
		}

		return nil
	})

	adminAPIEnabled := c.config.Admin != nil && ref.Deref(c.config.Admin.Enabled, false)
	if !adminAPIEnabled && c.storageHandler.IsWatchchable() {
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

func (c *Core) ensureAllFielsdSet() {
	var err error
	if c.config == nil {
		log.Fatal("config cannot be nil")
		return
	}

	if c.logHandler == nil {
		c.logHandler = logger.DefaultLogHandler
	}

	if c.errorHandler == nil {
		c.errorHandler = handlers.ErrorHandler
	}

	if c.storageHandler == nil {
		c.logHandler.Warn("no storage handler set, using memory storage")
		if c.storageHandler, err = storage.NewMemoryStorage(&storage.Opts{Config: c.config, LogHandler: c.logHandler}); err != nil {
			c.logHandler.Error("failed to load storage handler", logger.Field("error", err))
			os.Exit(1)
		}
	}
}
