package commands

import (
	"fmt"
	"os"
	"path"

	cli "github.com/urfave/cli/v2"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/storage"
	"go.trulyao.dev/lito/pkg/types"
)

func InitConfigCmd() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Create the configuration file using the specified storage format",
		Args:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Usage: "The storage format/backend to use",
				Value: "toml",
			},
		},
		Action: initConfig,
	}
}

func initConfig(c *cli.Context) error {
	filename := ""
	targetPath := c.Args().First()
	format := c.String("format")

	opts := &storage.Opts{
		Config: &types.Config{
			Admin: &types.DefaultAdmin,
			Proxy: &types.DefaultProxy,
			Services: map[string]*types.Service{
				"example": &types.DefaultService,
			},
		},
		LogHandler: logger.DefaultLogHandler,
	}

	var (
		storageFormat storage.Storage
		err           error
	)

	switch types.Storage(format) {
	case types.StorageTOML:
		filename = "lito.toml"
		storageFormat, err = storage.NewTOMLStorage(opts)
	case types.StorageJSON:
		filename = "lito.json"
		storageFormat, err = storage.NewJSONStorage(opts)
	default:
		return cli.Exit(fmt.Sprintf("unsupported storage format: %s", format), 1)
	}

	if err != nil {
		return cli.Exit(err, 1)
	}

	// Ensure we have a target path and filename (just in case)
	if targetPath == "" {
		targetPath = "."
	}
	if filename == "" {
		return cli.Exit("failed to determine filename", 1)
	}

	// If the target path is a directory, append the filename to the path
	if isDir(targetPath) {
		targetPath = path.Join(targetPath, filename)
	}

	opts.Config.Proxy.ConfigPath = ref.Ref(targetPath)
	if err := storageFormat.Persist(); err != nil {
		return cli.Exit(err, 1)
	}

	return nil
}

func isDir(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}
