package storage

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/types"
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
	switch instance.GetProxyConfig().Storage {
	case types.StorageJSON:
		s.instance = instance
		return NewJSONStorage(), nil
	default:
		return nil, fmt.Errorf("Unknown storage type: %s", instance.GetProxyConfig().Storage)
	}
}
