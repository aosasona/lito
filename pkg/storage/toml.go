package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type TOML struct {
	config     *types.Config
	logHandler logger.Logger
}

func NewTOMLStorage(opts *Opts) (*TOML, error) {
	if opts == nil {
		return nil, fmt.Errorf("opts provided in NewJSONStorage cannot be nil")
	}

	if opts.Config == nil {
		return nil, fmt.Errorf("config provided in NewJSONStorage cannot be nil")
	}

	return &TOML{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}, nil
}

func (t *TOML) IsWatchchable() bool { return true }
