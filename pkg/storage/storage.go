package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

type Storage interface {
	Path() string
	Load() error
	Persist() error
}

type Opts struct {
	Config     *types.Config
	LogHandler logger.Logger
}

// This needs to be setup to track the main instances of the Lito struct fields
func New(opts *Opts) (Storage, error) {
	utils.Assert(opts != nil, "opts is nil in storage package - this is probably a bug")
	utils.Assert(opts.Config != nil, "opts.Config is nil in storage package - this is probably a bug")
	utils.Assert(opts.LogHandler != nil, "opt.LogHandler is nil in storage package - this is probably a bug")

	if opts.Config.Proxy.IsNone() {
		return nil, fmt.Errorf("Proxy config is not set, this may be a bug, please investigate")
	}

	if opts.Config.Proxy.Unwrap().Storage.IsNone() {
		return nil, fmt.Errorf("Storage config is not set, must be one of: memory, json")
	}

	switch opts.Config.Proxy.Unwrap().Storage.Unwrap() {
	case types.StorageJSON:
		return NewJSONStorage(opts), nil
	case types.StorageMemory:
		return NewMemoryStorage(opts), nil
	default:
		return nil, fmt.Errorf("Unknown storage type: %v", opts.Config.Proxy.Unwrap().Storage)
	}
}
