package core

import "go.trulyao.dev/lito/pkg/types"

func (c *Core) RunProxy() error {
	return nil
}

func (c *Core) findServiceByName(name string) (*types.Service, bool) {
	service, ok := c.config.Services[name]
	if !ok {
		return nil, false
	}
	return service, true
}

func (c *Core) findServiceByProxiedDomain(domainName string) (string, *types.Service, bool) {
	for name, service := range c.config.Services {
		for _, serviceHost := range service.Domains {
			if serviceHost.DomainName == domainName {
				return name, service, true
			}
		}
	}
	return "", nil, false
}
