package lito

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/utils"
)

type Lito struct {
	Config     *Config
	LogHandler logger.Logger
}

type Opts = Lito

func New(opts *Opts) (*Lito, error) {
	if opts.Config == nil {
		return nil, ErrNoConfigSpecified
	}

	if opts.LogHandler == nil {
		opts.LogHandler = &logger.DefaultLogHandler
	}

	return &Lito{
		Config:     opts.Config,
		LogHandler: opts.LogHandler,
	}, nil
}

func (l *Lito) Run() error {
	if l.Config.Proxy.HTTPPort <= 0 {
		l.Config.Proxy.HTTPPort = 80
	}

	if l.Config.Proxy.HTTPSPort <= 0 {
		l.Config.Proxy.HTTPSPort = 443
	}

	l.LogHandler.Info("Starting Lito", logger.Param{Key: "http", Value: l.Config.Proxy.HTTPPort}, logger.Param{Key: "https", Value: l.Config.Proxy.HTTPSPort})

	if l.Config.Admin.Enabled {
		if l.Config.Admin.Port <= 0 {
			l.Config.Admin.Port = 2023
		}

		if l.Config.Admin.APIKey == "" {
			l.LogHandler.Warn("Admin API key not specified, generating a random one")
			l.Config.Admin.APIKey = utils.GenerateAPIKey()
			l.LogHandler.Info(fmt.Sprintf("Admin API key: %s - this will not be shown again, please store it somewhere safe or change it via the admin API", l.Config.Admin.APIKey))
		}

		if l.Config.Admin.Port > 0 {
			l.LogHandler.Info(fmt.Sprintf("Admin API listening on port %d", l.Config.Admin.Port))
		}
	}

	return nil
}
