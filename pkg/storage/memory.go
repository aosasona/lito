package storage

import (
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type Memory struct {
	config     *types.Config
	logHandler logger.Logger
}

func NewMemoryStorage(opts *Opts) *Memory {
	return &Memory{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}
}

func (m *Memory) IsWatchchable() bool { return false }

func (m *Memory) Path() string { return ":memory:" }

func (m *Memory) Load() error { return nil }

func (m *Memory) Persist() error {
	m.logHandler.Warn("Storage is set to memory, skipping config persistence - this is NOT recommended for production use")
	return nil
}
