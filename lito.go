package lito

import "go.trulyao.dev/lito/pkg/logger"

type Lito struct {
	Config     *Config
	LogHandler logger.Logger
}

type Opts struct {
	Lito
}

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
	return nil
}
