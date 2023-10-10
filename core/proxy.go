package core

import "go.trulyao.dev/lito/pkg/types"

func (c *Core) RunProxy() error {
	return nil
}

func (c *Core) lookupServiceByName(name string) *types.Service {
	service, ok := c.config.Services[name]
	if !ok {
		return nil
	}
	return service
}

func (c *Core) lookupServiceByHost(host string) *types.Service {
	for _, service := range c.config.Services {
		for _, serviceHost := range service.Domains {
			if serviceHost.DomainName == host {
				return service
			}
		}
	}
	return nil
}
