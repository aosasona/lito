package config

type InitConfig struct {
	ConfigDir string // the directory where the lito.toml config file is located not including the file name itself eg. etc/lito
	DataDir   string // the directory where the data should be stored (e.g. /var/lib/lito)
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

	if initConfig.DataDir == "" {
		initConfig.ConfigDir = "./data"
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
