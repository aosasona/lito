package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

type store struct {
	instance types.Instance
}

var s = store{}

type Storage interface {
	Path() string
	Load() error
	Persist() error
}

// This needs to be setup to track the main instances of the Lito struct fields
func New(instance types.Instance) (Storage, error) {
	s.instance = instance
	defer utils.Assert(s.instance != nil, "Instance is nil in storage package - this is definitely a bug")

	switch instance.GetProxyConfig().Storage {
	case types.StorageJSON:
		return NewJSONStorage(), nil
	case types.StorageMemory:
		return NewMemoryStorage(), nil
	default:
		return nil, fmt.Errorf("Unknown storage type: %s", instance.GetProxyConfig().Storage)
	}
}
