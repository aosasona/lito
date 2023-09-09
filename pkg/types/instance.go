package types

import (
	"go.trulyao.dev/lito/pkg/logger"
)

type Instance interface {
	// Getters
	GetLogHandler() logger.Logger
	GetAdminConfig() *Admin
	GetServicesConfig() map[string]Service
	GetProxyConfig() *Proxy

	// Get config as JSON string
	MarshalConfig() (string, error)
	// Unmarshal JSON string to config
	UnmarshalConfig([]byte) (*Config, error)
	// Save current config to disk
	Commit() error
	// Update config
	UpdateConfig(*Config) error
	// Reload config from disk
	Reload() error

	// Mutex functions
	Lock()
	Unlock()
}
