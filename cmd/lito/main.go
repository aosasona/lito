package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito"
)

const version = "1.0.0"

var (
	config *lito.Config

	flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "enable-admin",
			Aliases:     []string{"a"},
			Usage:       "Enable the admin API",
			Value:       true,
			Destination: &config.Admin.Enabled,
		},
		&cli.IntFlag{
			Name:        "admin-port",
			Usage:       "The port that the admin API will listen on",
			Value:       2080,
			Destination: &config.Admin.Port,
		},
	}
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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
