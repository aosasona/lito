package main

import (
	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

var (
	debug = false

	enableAdmin = false
	adminPort   = 0
	adminApiKey = ""
	configPath  = ""

	proxyHost           = ""
	httpPort            = 0
	httpsPort           = 0
	tlsEmail            = ""
	enableTLS           = true
	enableHTTPSRedirect = true
	storageType         = types.StorageMemory
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "debug",
		Aliases:     []string{"d"},
		Usage:       "Enable debug mode",
		Value:       false,
		Destination: &debug,
	},
	&cli.BoolFlag{
		Name:        "enable-admin",
		Aliases:     []string{"a"},
		Usage:       "Enable the admin API",
		Value:       false,
		Destination: &enableAdmin,
	},
	&cli.IntFlag{
		Name:        "admin-port",
		Usage:       "The port that the admin API will listen on",
		Value:       2023,
		Destination: &adminPort,
	},
	&cli.StringFlag{
		Name:        "admin-key",
		Usage:       `The API key that will be used to authenticate with the admin API. If not specified, a random one will be generated.`,
		Value:       "",
		Destination: &adminApiKey,
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
		Destination: &configPath,
	},
	&cli.StringFlag{
		Name:        "host",
		Aliases:     []string{"H"},
		Usage:       "The host that the proxy will listen on",
		Value:       "0.0.0.0",
		Destination: &proxyHost,
	},
	&cli.IntFlag{
		Name:        "http-port",
		Aliases:     []string{"p"},
		Usage:       "The HTTP port that the proxy will listen on",
		Value:       80,
		Destination: &httpPort,
	},
	&cli.IntFlag{
		Name:        "https-port",
		Aliases:     []string{"s"},
		Usage:       "The HTTPS port that the proxy will listen on",
		Value:       443,
		Destination: &httpsPort,
	},
	&cli.StringFlag{
		Name:        "tls-email",
		Usage:       "The email address that will be used to register TLS certificates",
		Value:       "",
		Destination: &tlsEmail,
	},
	&cli.BoolFlag{
		Name:        "enable-tls",
		Usage:       "Enable TLS",
		Value:       true,
		Destination: &enableTLS,
	},
	&cli.BoolFlag{
		Name:        "enable-https-redirect",
		Aliases:     []string{"r"},
		Usage:       "Enable HTTP to HTTPS redirects",
		Value:       true,
		Destination: &enableHTTPSRedirect,
	},
	&cli.StringFlag{
		Name:        "storage",
		Usage:       "The storage backend to use",
		Value:       "json",
		Destination: (*string)(&storageType),
	},
}
