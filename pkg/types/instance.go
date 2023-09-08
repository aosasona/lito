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

	// Actions
	String() (string, error)
	Commit() error
	Reload() error

	// Mutex functions
	Lock()
	Unlock()
}
