package core

import (
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type Core struct {
	config     *types.Config
	logHandler logger.Logger
}

type Opts struct {
	Config     *types.Config
	LogHandler logger.Logger
}

func New(opts *Opts) *Core {
	return &Core{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}
}

func (c *Core) Run() error {
	return nil
}
