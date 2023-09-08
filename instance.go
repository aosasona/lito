package lito

// Instance interface implementations

import (
	"encoding/json"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

func (l *Lito) String() (string, error) {
	s, err := json.Marshal(l.Config)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func (l *Lito) GetLogHandler() logger.Logger { return l.LogHandler }

func (l *Lito) GetAdminConfig() *types.Admin { return &l.Config.Admin }

func (l *Lito) GetServicesConfig() map[string]types.Service { return l.Config.Services }

func (l *Lito) GetProxyConfig() *types.Proxy { return &l.Config.Proxy }

func (l *Lito) Reload() error { return l.StorageHandler.Load() }

func (l *Lito) Lock() { l.Config.mutex.Lock() }

func (l *Lito) Unlock() { l.Config.mutex.Unlock() }

func (l *Lito) Commit() error { return l.StorageHandler.Persist() }
