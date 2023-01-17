package lito

import (
	"errors"
	"strings"

	"github.com/aosasona/lito/pkg/errtype"
	"github.com/aosasona/lito/pkg/httpfunc"
)

type Proxy struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Target struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Service struct {
	Name          string   `toml:"name"`
	Targets       []Target `toml:"targets"`
	AccessControl struct {
		Private  bool   `toml:"private"`
		User     string `toml:"user"`
		Password string `toml:"user"`
	} `toml:"access_control"`
}

type Config struct {
	Proxy    Proxy     `toml:"proxy"`
	Services []Service `toml:"services"`
}

func New(config *Config) error {
	err := validateProxyInfo(config.Proxy.Host, config.Proxy.Port)
	if err != nil {
		return err
	}

	portIsAvailable := httpfunc.IsPortAvailable(config.Proxy.Host, config.Proxy.Port)
	if !portIsAvailable {
		return errors.New(errtype.PORT_NOT_AVAILABLE)
	}

	return nil
}

func validateProxyInfo(host string, port int) error {
	if host == "" {
		return errors.New(errtype.NO_HOST_PROVIDED)
	}

	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		return errors.New(errtype.NO_HOST_PROTOCOL_PROVIDED)
	}

	if port == 0 {
		return errors.New(errtype.PORT_ZERO)
	}

	freePort := httpfunc.IsPortAvailable(host, port)
	if !freePort {
		return errors.New(errtype.PORT_NOT_AVAILABLE)
	}

	return nil
}
