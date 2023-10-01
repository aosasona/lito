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
	defer utils.Assert(opts != nil, "opts is nil in storage package - this is definitely a bug")
	defer utils.Assert(opts.Config != nil, "opts.Config is nil in storage package - this is definitely a bug")
	defer utils.Assert(opts.LogHandler != nil, "opt.LogHandler is nil in storage package - this is definitely a bug")

	switch opts.Config.Proxy.Storage {
	case types.StorageJSON:
		return NewJSONStorage(opts), nil
	case types.StorageMemory:
		return NewMemoryStorage(opts), nil
	default:
		return nil, fmt.Errorf("Unknown storage type: %s", opts.Config.Proxy.Storage)
	}
}
