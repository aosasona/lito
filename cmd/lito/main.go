package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/cmd/lito/commands"
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
		Usage:                  "A flexible and lightweight reverse proxy",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
	}

	app.Commands = []*cli.Command{
		commands.RunCmd(),
		commands.InitConfigCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
