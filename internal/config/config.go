package config

type InitConfig struct {
	ConfigDir string // the directory where the lito.toml config file is located not including the file name itself eg. etc/lito
}

type Config struct {
	lito     *LitoConfig
	services *[]ServicesConfig
}

func Load(initConfig InitConfig) (*Config, error) {
	config := new(Config)

	if initConfig.ConfigDir == "" {
		initConfig.ConfigDir = "."
	}

	litoConfig, err := loadLitoConfig(initConfig.ConfigDir)
	if err != nil {
		return nil, err
	}

	config.lito = litoConfig

	return config, nil
}

func (c *Config) Lito() *LitoConfig {
	return c.lito
}

func (c *Config) Services() *[]ServicesConfig {
	return c.services
}
