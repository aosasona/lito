package storage

import (
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type JSON struct {
	config     *types.Config
	logHandler logger.Logger
}

func NewJSONStorage(opts *Opts) *JSON {
	return &JSON{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}
}

func (j *JSON) Path() string {
	return j.config.Proxy.ConfigPath
}

func (j *JSON) Load() error {
	panic("not implemented")
}

func (j *JSON) Persist() error {
	panic("not implemented")
}
