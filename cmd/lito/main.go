package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/core"
	"go.trulyao.dev/lito/ext/option"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

const (
	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Name:                   "lito",
		Version:                version,
		Suggest:                true,
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Usage:                  "A flexible and lightweight reverse proxy - with automatic TLS (powered by Caddy's CertMagic)",
		Flags:                  flags,
		Action: func(c *cli.Context) error {
			return run()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := core.New(&core.Opts{
		Debug: debug,
		Config: &types.Config{
			Admin: option.Some(&types.Admin{
				Enabled: option.Some(enableAdmin),
				Port:    option.IntValue(adminPort),
				APIKey:  option.StringValue(adminApiKey),
			}),
			Proxy: option.Some(&types.Proxy{
				ConfigPath:          option.StringValue(configPath),
				Host:                option.StringValue(proxyHost),
				HTTPPort:            option.IntValue(httpPort),
				HTTPSPort:           option.IntValue(httpsPort),
				TLSEmail:            option.StringValue(tlsEmail),
				EnableTLS:           option.BoolValue(enableTLS),
				EnableHTTPSRedirect: option.BoolValue(enableHTTPSRedirect),
				Storage:             option.Some(storageType),
			}),
		},
		LogHandler: logger.DefaultLogHandler,
	})

	return c.Run()
}
