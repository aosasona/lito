package lito

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/controllers"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

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
	controllers.Init(l)
	controllers.FillDefaults()

	storage.Init(l)

	switch l.Config.Proxy.Storage {
	case types.StorageJSON:
		l.StorageHandler = storage.NewJSONStorage(l.Config.Proxy.ConfigPath)
	default:
		l.LogHandler.Fatal(fmt.Sprintf("Unknown storage type: %s", l.Config.Proxy.Storage))
	}

	if err := l.StorageHandler.Load(); err != nil {
		l.LogHandler.Fatal(fmt.Sprintf("Failed to load config: %s", err.Error()))
	}
}

func (l *Lito) Run() error {
	l.setup()

	l.LogHandler.Info(fmt.Sprintf("Starting Lito on port :%d", l.Config.Proxy.HTTPPort), logger.Param{Key: "http", Value: l.Config.Proxy.HTTPPort}, logger.Param{Key: "https", Value: l.Config.Proxy.HTTPSPort})

	return nil
}
