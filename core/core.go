package core

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.trulyao.dev/lito/core/api/handlers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type Core struct {
	config       *types.Config
	logHandler   logger.Logger
	errorHandler types.ErrorHandler
}

type Opts struct {
	Config       *types.Config
	LogHandler   logger.Logger
	ErrorHandler types.ErrorHandler
}

func New(opts *Opts) *Core {
	errorHandler := handlers.ErrorHandler
	if opts.ErrorHandler != nil {
		errorHandler = opts.ErrorHandler
	}

	var logHandler logger.Logger
	logHandler = logger.DefaultLogHandler
	if opts.LogHandler != nil {
		logHandler = opts.LogHandler
	}

	return &Core{
		config:       opts.Config,
		logHandler:   logHandler,
		errorHandler: errorHandler,
	}
}

// HandleShutdown handles the shutdown signal and gracefully shutdown the proxy while maintaining the current config and state
// You should call this before you start the proxy - it will block until the shutdown signal is received
func (c *Core) HandleShutdown(sig chan os.Signal) {
	done := make(chan bool, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// TODO: shutdown the proxy and admin API if enabled

	<-done
}

func (c *Core) Run() error {
	wg := sync.WaitGroup{}

	sig := make(chan os.Signal, 1)
	go func() {
		wg.Add(1)
		c.HandleShutdown(sig)
	}()

	// TODO: run the proxy
	// TODO: run the admin API

	wg.Wait()

	return nil
}
