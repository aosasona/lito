package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/core"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

const version = "1.0.0"

var config = types.Config{}

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
		Config:     &config,
		LogHandler: logger.DefaultLogHandler,
	})

	return c.Run()
}
