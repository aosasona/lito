package lito

import (
	"sync"

	"go.trulyao.dev/lito/pkg/types"
)

type Config struct {
	Admin    types.Admin              `json:"admin"`
	Services map[string]types.Service `json:"services"`
	Proxy    types.Proxy              `json:"proxy"`
	mutex    sync.RWMutex
}

func LoadFromDirectory(path string) (*Config, error) {
	return nil, nil
}
