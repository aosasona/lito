package controllers

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/types"
	"go.trulyao.dev/lito/pkg/utils"
)

func fillAdminDefaults() {
	admin := c.admin()

	if admin.Enabled {
		if admin.Port <= 0 {
			admin.Port = 2023
		}

		if admin.APIKey == "" {
			c.logger().Warn("Admin API key not specified, generating a random one")
			admin.APIKey = utils.GenerateAPIKey()
			c.logger().Info(fmt.Sprintf("Admin API key: %s - this will not be shown again, please store it somewhere safe or change it via the admin API", admin.APIKey))
		}

		if admin.Port > 0 {
			c.logger().Info(fmt.Sprintf("Admin API listening on port %d", admin.Port))
		}
	}
}

func GetConfig() *types.Config {
	return c.instance.GetConfig()
}

func ValidateAPIKey(key string) bool {
	// return c.instance.ValidateAPIKey(key)
	return false
}
