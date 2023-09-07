package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito"
)

const version = "1.0.0"

var (
	config = lito.Config{}

	flags = []cli.Flag{
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
		Action: func(c *cli.Context) error {
			return runLito()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runLito() error {
	l, err := lito.New(&lito.Opts{Config: &config})
	if err != nil {
		return err
	}

	defer func() {
		if err := l.LogHandler.Sync(); err != nil {
			log.Fatalf("Failed to sync log handler: %v", err)
		}
	}()

	return l.Run()
}
