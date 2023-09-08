package lito

import (
	"fmt"
	"sync"

	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

type Config struct {
	Admin    types.Admin              `json:"admin"`
	Services map[string]types.Service `json:"services"`
	Proxy    types.Proxy              `json:"proxy"`
	mutex    sync.RWMutex
}

type Lito struct {
	Config         *Config
	LogHandler     logger.Logger
	StorageHandler storage.Storage
}

type Opts = Lito // alias

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

func (l *Lito) Run() error {
	l.setup()

	l.LogHandler.Info(fmt.Sprintf("Starting Lito on port :%d", l.Config.Proxy.HTTPPort), logger.Param{Key: "http", Value: l.Config.Proxy.HTTPPort}, logger.Param{Key: "https", Value: l.Config.Proxy.HTTPSPort})

	return nil
}
