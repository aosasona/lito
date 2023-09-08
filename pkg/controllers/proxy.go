package controllers

import (
	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

const (
	HostLocalhost = "localhost"
	HostAll       = "0.0.0.0"
	Host127001    = "127.0.0.1"
)

func fillProxyDefaults() {
	if c.proxy().Host == "" {
		c.proxy().Host = HostAll
	}

	if !utils.Contains([]string{"localhost", "127.0.0.1", "0.0.0.0"}, c.proxy().Host) {
		c.logger().Warn("Proxy host is not valid, defaulting to 0.0.0.0")
		c.proxy().Host = HostAll
	}

	if c.proxy().HTTPPort <= 0 {
		c.proxy().HTTPPort = 80
	}

	if c.proxy().EnableTLS {
		if c.proxy().HTTPSPort <= 0 {
			c.proxy().HTTPSPort = 443
		}

		if c.proxy().TLSEmail == "" {
			c.logger().Warn("TLS email not specified, this is required to register TLS certificates")
			c.proxy().EnableTLS = false
		}
	}

	if c.proxy().ConfigPath == "" {
		c.proxy().ConfigPath = "./lito.json"
	}

	if !utils.Contains([]types.Storage{types.StorageMemory, types.StorageJSON}, c.proxy().Storage) {
		c.logger().Warn("Storage type is not valid, defaulting to memory")
		c.proxy().Storage = types.StorageMemory
	}
}
