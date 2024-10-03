package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/types"
)

type Storage interface {
	// Get the config's path or initialize it if empty
	Path() string

	// Load the configuration from the appropriate source
	Load() error

	// Save the config back back to the appropriate destination; usually the disk
	Persist() error

	// Whether a storage type is watchable or not
	IsWatchchable() bool

	// Create the configuration file if required, this method should be safe to call multiple times
	CreateIfNotExists(...string) error
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

	if opts.Config.Proxy == nil {
		return nil, fmt.Errorf(
			"proxy config is not set, this might be a bug, please investigate and/or report",
		)
	}

	if opts.Config.Proxy.Storage == nil {
		return nil, fmt.Errorf("storage config is not set, must be one of: memory, json")
	}

	switch ref.Deref(opts.Config.Proxy.Storage, types.StorageMemory) {
	case types.StorageJSON:
		return NewJSONStorage(opts)
	case types.StorageMemory:
		return NewMemoryStorage(opts)
	case types.StorageTOML:
		return NewTOMLStorage(opts)
	default:
		return nil, fmt.Errorf("Unknown storage type: %v", opts.Config.Proxy.Storage)
	}
}

var (
	_ Storage = (*JSON)(nil)
	_ Storage = (*TOML)(nil)
	_ Storage = (*Memory)(nil)
)
