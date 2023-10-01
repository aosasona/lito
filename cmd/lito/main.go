package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/pkg/types"
)

const version = "1.0.0"

var (
	config              = types.Config{}
	overwriteDiskConfig bool
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
			return runLito()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runLito() error {
	panic("not implemented")
}
