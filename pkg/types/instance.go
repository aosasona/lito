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
	String() (string, error)
	// Save current config to disk
	Commit() error
	// Reload config from disk
	Reload() error

	// Mutex functions
	Lock()
	Unlock()
}
