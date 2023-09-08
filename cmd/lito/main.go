package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.trulyao.dev/lito"
)

const version = "1.0.0"

var config = lito.Config{}

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
