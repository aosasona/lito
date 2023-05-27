package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type LitoConfig struct {
	Admin AdminConfig `mapstructure:"admin"`
	Proxy ProxyConfig `mapstructure:"proxy"`
}

type AdminConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Path     string `mapstructure:"path"`
}

type ProxyConfig struct {
	Host        string          `mapstructure:"host"`
	Ports       []uint          `mapstructure:"ports"`
	LogDir      string          `mapstructure:"log_dir"`
	EnableTLS   bool            `mapstructure:"enable_tls"`
	WatchConfig bool            `mapstructure:"watch_config"`
	Services    ServicesConfig  `mapstructure:"services"`
	Discovery   DiscoveryConfig `mapstructure:"discovery"`
}

type ServicesConfig struct {
	Data            string         `mapstructure:"data"`
	StorageDriver   string         `mapstructure:"storage_driver"`
	RefreshInterval uint           `mapstructure:"refresh_interval"`
	NotFound        NotFoundConfig `mapstructure:"not_found"`
}

type NotFoundConfig struct {
	Type    string `mapstructure:"type"`
	Content string `mapstructure:"content"`
}

type DiscoveryConfig struct {
	Path            string `mapstructure:"path"`
	AllowExternal   bool   `mapstructure:"allow_external"`
	RefreshInterval uint   `mapstructure:"refresh_interval"`
}

func loadLitoConfig(configPath string) (*LitoConfig, error) {
	config := new(LitoConfig)

	v := viper.New()
	v.SetConfigName("lito")
	v.SetConfigType("toml")
	v.AddConfigPath(strings.TrimSuffix(configPath, "/"))
	v.AddConfigPath(".")

	// TODO: Add hot-reloading config file on change

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("unable to read config from file: %v", err.Error())
		}
		return nil, err
	}

	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to load marshal into struct: %v", err.Error())
	}

	return config, nil
}
