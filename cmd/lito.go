package main

import (
	"errors"
	"strings"

	"github.com/aosasona/lito/pkg/httpfunc"
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

	portIsAvailable := httpfunc.IsPortAvailable(host, port)
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

	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		return errors.New("host protocol not defined (add http:// or https://)")
	}

	if port == 0 {
		return errors.New("port is empty")
	}

	freePort := httpfunc.IsPortAvailable(host, port)
	if !freePort {
		return errors.New("port is not available")
	}

	return nil
}
