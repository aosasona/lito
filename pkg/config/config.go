package config

import (
	"fmt"

	"github.com/aosasona/lito/cmd/lito"
	"github.com/spf13/viper"
)

func Load(path string) error {
	defaultPaths := [...]string{".", "~", "/etc/lito"}

	viper.SetConfigName("lito")
	viper.SetConfigType("toml")

	for _, val := range defaultPaths {
		viper.AddConfigPath(val)
	}
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		msg := fmt.Errorf("Unable to load config: %v", err.Error())
		return msg
	}

	return nil
}

func Unwrap() *lito.Config {
	var config lito.Config

	viper.Unmarshal(config)

	return &config
}
