package lito

import (
	"errors"

	"github.com/aosasona/lito/pkg/http_func"
)

type Proxy struct {
	Host       string
	Port       int
	ConfigPath string
}

func NewProxy(host string, port int, configPath string) (*Proxy, error) {
	if configPath == "" {
		configPath = "."
	}

	err := validateProxyInfo(host, port)
	if err != nil {
		return nil, err
	}

	portIsAvailable := http_func.IsPortAvailable(host, port)
	if !portIsAvailable {
		return nil, errors.New("port is not available")
	}

	return &Proxy{
		Host:       host,
		Port:       port,
		ConfigPath: configPath,
	}, nil
}

func validateProxyInfo(host string, port int) error {
	if host == "" {
		return errors.New("host is empty")
	}

	if port == 0 {
		return errors.New("port is empty")
	}

	return nil
}
