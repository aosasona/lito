package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type Memory struct {
	config     *types.Config
	logHandler logger.Logger
}

func NewMemoryStorage(opts *Opts) (*Memory, error) {
	if opts == nil {
		return nil, fmt.Errorf("opts provided in NewMemoryStorage cannot be nil")
	}

	if opts.Config == nil {
		return nil, fmt.Errorf("config provided in NewMemoryStorage cannot be nil")
	}

	return &Memory{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}, nil
}

func (m *Memory) IsWatchchable() bool { return false }

func (m *Memory) Path() string { return ":memory:" }

func (m *Memory) Load() error { return nil }

func (m *Memory) Persist() error {
	m.warn("Storage is set to memory, skipping config persistence - this is NOT recommended for production use")
	return nil
}

func (m *Memory) warn(msg string, params ...logger.Param) {
	m.logHandler.Warn(msg, params...)
}
