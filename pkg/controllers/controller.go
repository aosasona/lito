package controllers

import (
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

type controller struct {
	instance types.Instance
}

var c = controller{}

// This needs to be setup to track the main instances of the Lito struct fields
func Init(i types.Instance) {
	c.instance = i
}

func FillDefaults() {
	fillAdminDefaults()
	fillProxyDefaults()
}

func (controller *controller) admin() *types.Admin {
	return controller.instance.GetAdminConfig()
}

func (controller *controller) services() map[string]types.Service {
	return controller.instance.GetServicesConfig()
}

func (controller *controller) proxy() *types.Proxy {
	return controller.instance.GetProxyConfig()
}

func (controller *controller) logger() logger.Logger {
	return controller.instance.GetLogHandler()
}
