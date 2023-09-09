package main

import (
	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/pkg/logger"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "overwrite-config",
		Aliases:     []string{"o"},
		Usage:       "Overwrite the config file with the current config",
		Value:       false,
		Destination: &overwriteDiskConfig,
	},
	&cli.BoolFlag{
		Name:        "enable-admin",
		Aliases:     []string{"a"},
		Usage:       "Enable the admin API",
		Value:       false,
		Destination: &config.Admin.Enabled,
	},
	&cli.IntFlag{
		Name:        "admin-port",
		Usage:       "The port that the admin API will listen on",
		Value:       2023,
		Destination: &config.Admin.Port,
	},
	&cli.StringFlag{
		Name:        "admin-key",
		Usage:       `The API key that will be used to authenticate with the admin API. If not specified, a random one will be generated.`,
		Value:       "",
		Destination: &config.Admin.APIKey,
	},
	&cli.StringFlag{
		Name:        "log-file",
		Aliases:     []string{"l"},
		Usage:       "The path to the log file",
		Value:       "lito.log",
		Destination: &logger.DefaultLogHandler.Path,
	},
	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "The path to the config file",
		Value:       "lito.json",
		Destination: &config.Proxy.ConfigPath,
	},
	&cli.StringFlag{
		Name:        "host",
		Aliases:     []string{"H"},
		Usage:       "The host that the proxy will listen on",
		Value:       "0.0.0.0",
		Destination: &config.Proxy.Host,
	},
	&cli.IntFlag{
		Name:        "http-port",
		Aliases:     []string{"p"},
		Usage:       "The HTTP port that the proxy will listen on",
		Value:       80,
		Destination: &config.Proxy.HTTPPort,
	},
	&cli.IntFlag{
		Name:        "https-port",
		Aliases:     []string{"s"},
		Usage:       "The HTTPS port that the proxy will listen on",
		Value:       443,
		Destination: &config.Proxy.HTTPSPort,
	},
	&cli.StringFlag{
		Name:        "tls-email",
		Usage:       "The email address that will be used to register TLS certificates",
		Value:       "",
		Destination: &config.Proxy.TLSEmail,
	},
	&cli.BoolFlag{
		Name:        "enable-tls",
		Usage:       "Enable TLS",
		Value:       true,
		Destination: &config.Proxy.EnableTLS,
	},
	&cli.BoolFlag{
		Name:        "enable-https-redirect",
		Aliases:     []string{"r"},
		Usage:       "Enable HTTP to HTTPS redirects",
		Value:       true,
		Destination: &config.Proxy.EnableHTTPSRedirect,
	},
	&cli.StringFlag{
		Name:        "storage",
		Usage:       "The storage backend to use",
		Value:       "json",
		Destination: (*string)(&config.Proxy.Storage),
	},
}
