package core

import "github.com/aosasona/lito/internal/config"

type Core struct {
	config *config.Config
}

func New(configDir string) (*Core, error) {
	core := new(Core)
	conf, err := config.Load(config.InitConfig{ConfigDir: configDir})
	if err != nil {
		return nil, err
	}
	core.config = conf
	return core, nil
}

func (c *Core) Config() *config.Config {
	return c.config
}
