package lito

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.trulyao.dev/lito/api"
	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

type Lito struct {
	Config         *types.Config
	LogHandler     logger.Logger
	StorageHandler storage.Storage
}

type Opts = Lito // alias

type RunOpts struct {
	OverrideDiskConfig bool
}

func New(opts *Opts) (*Lito, error) {
	if opts.Config == nil {
		return nil, utils.ErrNoConfigSpecified
	}

	if opts.LogHandler == nil {
		opts.LogHandler = &logger.DefaultLogHandler
	}

	return &Lito{
		Config:     opts.Config,
		LogHandler: opts.LogHandler,
	}, nil
}

func (l *Lito) setup() {
	var err error

	controllers.Init(l)

	if l.StorageHandler, err = storage.New(l); err != nil {
		l.LogHandler.Fatal(fmt.Sprintf("Failed to initialize storage handler: %s", err.Error()))
	}

	if err = l.StorageHandler.Load(); err != nil {
		l.LogHandler.Fatal(fmt.Sprintf("Failed to load config: %s", err.Error()))
	}
}

func (l *Lito) handleInterrupt() {
	sig := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		done <- true
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// commit config before exiting
		// TODO: find out why this commits a blank config
		if err := l.Commit(); err != nil {
			l.LogHandler.Fatal(fmt.Sprintf("Failed to commit config: %s", err.Error()))
		}

		// shutdown admin API
		if err := api.Shutdown(ctx); err != nil {
			l.LogHandler.Fatal(fmt.Sprintf("Failed to shutdown admin API: %s", err.Error()))
		}

		l.LogHandler.Info("Received interrupt, shutting down...")
	}()

	<-done
}

func (l *Lito) runAdminServer() {
	l.LogHandler.Info(fmt.Sprintf("Starting admin API on port :%d", l.Config.Admin.Port))
	if err := api.ServeAdminAPI(l.Config.Admin.Port); err != nil {
		l.LogHandler.Fatal(fmt.Sprintf("Failed to start admin API: %s", err.Error()))
	}
}

func (l *Lito) runProxy() {
	l.LogHandler.Info(
		fmt.Sprintf("Starting Lito on port :%d", l.Config.Proxy.HTTPPort),
		logger.Param{Key: "http", Value: l.Config.Proxy.HTTPPort},
		logger.Param{Key: "https", Value: l.Config.Proxy.HTTPSPort},
	)
}

func (l *Lito) Run(opts RunOpts) error {
	l.setup()

	if opts.OverrideDiskConfig {
		l.LogHandler.Warn("Overriding disk config is enabled, this may cause unexpected behavior")
		if err := l.Commit(); err != nil {
			return err
		}
	}

	var wg sync.WaitGroup

	if l.Config.Admin.Enabled {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.runAdminServer()
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.runProxy()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.handleInterrupt()
	}()

	wg.Wait()

	return nil
}
