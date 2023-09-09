package lito

// Instance interface implementations

import (
	"encoding/json"
	"fmt"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

// String returns the JSON representation of the Lito config
func (l *Lito) MarshalConfig() (string, error) {
	s, err := json.Marshal(l.Config)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func (l *Lito) UnmarshalConfig(data []byte) (*types.Config, error) {
	c := new(types.Config)
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (l *Lito) UpdateConfig(c *types.Config) error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}

	// We keep the old struct and the current lock
	l.Config.Admin = c.Admin
	l.Config.Proxy = c.Proxy
	l.Config.Services = c.Services

	return nil
}

func (l *Lito) GetLogHandler() logger.Logger { return l.LogHandler }

func (l *Lito) GetAdminConfig() *types.Admin { return &l.Config.Admin }

func (l *Lito) GetServicesConfig() map[string]types.Service { return l.Config.Services }

func (l *Lito) GetProxyConfig() *types.Proxy { return &l.Config.Proxy }

func (l *Lito) Reload() error { return l.StorageHandler.Load() }

func (l *Lito) Lock() { l.Config.Lock() }

func (l *Lito) Unlock() { l.Config.Unlock() }

func (l *Lito) Commit() error {
	return l.StorageHandler.Persist()
}
