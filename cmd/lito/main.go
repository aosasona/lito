package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/core"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
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
			Admin: &types.Admin{
				Enabled: ref.Ref(enableAdmin),
				Port:    ref.Ref(adminPort),
				APIKey:  ref.Ref(adminApiKey),
			},
			Proxy: &types.Proxy{
				ConfigPath:          ref.Ref(configPath),
				Host:                ref.Ref(proxyHost),
				HTTPPort:            ref.Ref(httpPort),
				HTTPSPort:           ref.Ref(httpsPort),
				TLSEmail:            ref.Ref(tlsEmail),
				EnableTLS:           ref.Ref(enableTLS),
				EnableHTTPSRedirect: ref.Ref(enableHTTPSRedirect),
				Storage:             ref.Ref(storageType),
			},
		},
		LogHandler: logger.DefaultLogHandler,
	})

	return c.Run()
}
