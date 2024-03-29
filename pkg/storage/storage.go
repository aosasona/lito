package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type Storage interface {
	Path() string
	Load() error
	Persist() error
	IsWatchchable() bool
}

type Opts struct {
	Config     *types.Config
	LogHandler logger.Logger
}

// This needs to be setup to track the main instances of the Lito struct fields
func New(opts *Opts) (Storage, error) {
	if opts == nil {
		return nil, fmt.Errorf("opts provided in storage.New cannot be nil")
	}

	if opts.Config == nil {
		return nil, fmt.Errorf("config provided in storage.New cannot be nil")
	}

	if opts.Config.Proxy.IsNone() {
		return nil, fmt.Errorf("Proxy config is not set, this may be a bug, please investigate")
	}

	if opts.Config.Proxy.Unwrap(&types.DefaultProxy).Storage.IsNone() {
		return nil, fmt.Errorf("Storage config is not set, must be one of: memory, json")
	}

	switch opts.Config.Proxy.Unwrap(&types.DefaultProxy).Storage.Unwrap(types.StorageMemory) {
	case types.StorageJSON:
		return NewJSONStorage(opts)
	case types.StorageMemory:
		return NewMemoryStorage(opts)
	default:
		return nil, fmt.Errorf("Unknown storage type: %v", opts.Config.Proxy.Unwrap(&types.DefaultProxy).Storage)
	}
}
