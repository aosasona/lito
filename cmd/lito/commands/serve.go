package commands

import (
	cli "github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/core"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/types"
)

var runFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "Enable debug mode",
		Value:   false,
	},
	&cli.BoolFlag{
		Name:    "enable-admin",
		Aliases: []string{"a"},
		Usage:   "Enable the admin API",
		Value:   false,
	},
	&cli.IntFlag{
		Name:  "admin-port",
		Usage: "The port that the admin API will listen on",
		Value: 2023,
	},
	&cli.StringFlag{
		Name:  "admin-key",
		Usage: `The API key that will be used to authenticate with the admin API. If not specified, a random one will be generated.`,
		Value: "",
	},
	// TODO: make this log-dir when rolling logs is implemented
	&cli.StringFlag{
		Name:    "log-file",
		Aliases: []string{"l"},
		Usage:   "The path to the log file",
		Value:   "lito.log",
	},
	&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "The path to the config file",
		Value:   "lito.toml",
	},
	&cli.StringFlag{
		Name:    "host",
		Aliases: []string{"H"},
		Usage:   "The host that the proxy will listen on",
		Value:   "0.0.0.0",
	},
	&cli.IntFlag{
		Name:    "http-port",
		Aliases: []string{"p"},
		Usage:   "The HTTP port that the proxy will listen on",
		Value:   80,
	},
	&cli.IntFlag{
		Name:    "https-port",
		Aliases: []string{"s"},
		Usage:   "The HTTPS port that the proxy will listen on",
		Value:   443,
	},
	&cli.StringFlag{
		Name:  "tls-email",
		Usage: "The email address that will be used to register TLS certificates",
		Value: "",
	},
	&cli.BoolFlag{
		Name:  "enable-tls",
		Usage: "Enable TLS",
		Value: false,
	},
	&cli.BoolFlag{
		Name:    "enable-https-redirect",
		Aliases: []string{"r"},
		Usage:   "Enable HTTP to HTTPS redirects",
		Value:   true,
	},
	&cli.StringFlag{
		Name:  "storage",
		Usage: "The storage backend to use",
		Value: "toml",
	},
}

func RunCmd() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Usage:   "Start the proxy server using the config and flags provided",
		Aliases: []string{"serve", "start"},
		Flags:   runFlags,
		Action:  run,
	}
}

func run(ctx *cli.Context) error {
	c := core.New(&core.Opts{
		Debug: ctx.Bool("debug"),
		Config: &types.Config{
			Admin: &types.Admin{
				Enabled: ref.Ref(ctx.Bool("enable-admin")),
				Port:    ref.Ref(ctx.Int("admin-port")),
				APIKey:  ref.Ref(ctx.String("admin-key")),
			},
			Proxy: &types.Proxy{
				ConfigPath:          ref.Ref(ctx.String("config")),
				Host:                ref.Ref(ctx.String("host")),
				HTTPPort:            ref.Ref(ctx.Int("http-port")),
				HTTPSPort:           ref.Ref(ctx.Int("https-port")),
				TLSEmail:            ref.Ref(ctx.String("tls-email")),
				EnableTLS:           ref.Ref(ctx.Bool("enable-tls")),
				EnableHTTPSRedirect: ref.Ref(ctx.Bool("enable-https-redirect")),
				Storage:             ref.Ref(types.Storage(ctx.String("storage"))),
			},
		},
		LogHandler: logger.DefaultLogHandler,
	})

	return c.Run()
}
