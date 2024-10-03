package commands

import cli "github.com/urfave/cli/v2"

func InitConfigCmd() *cli.Command {
	return &cli.Command{
		Name:  "init <path>",
		Usage: "Create the configuration file using the specified storage format",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "storage",
				Usage: "The storage format/backend to use",
				Value: "toml",
			},
		},
		Action: initConfig,
	}
}

// TODO: infer storage type from filename
func initConfig(c *cli.Context) error {
	panic("not implemented")
}
