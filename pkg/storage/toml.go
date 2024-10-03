package storage

import (
	"fmt"
	"os"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/types"
)

type TOML struct {
	config     *types.Config
	logHandler logger.Logger
}

// Create a new instance of the TOML storage container
func NewTOMLStorage(opts *Opts) (*TOML, error) {
	if opts == nil {
		return nil, fmt.Errorf("opts provided in NewTOMLStorage cannot be nil")
	}

	if opts.Config == nil {
		return nil, fmt.Errorf("config provided in NewTOMLStorage cannot be nil")
	}

	return &TOML{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}, nil
}

func (t *TOML) IsWatchchable() bool { return true }

// Get the path to the config
func (t *TOML) Path() string {
	if t.config.Proxy == nil {
		t.config.Proxy = &types.DefaultProxy
		t.config.Proxy.ConfigPath = ref.Ref("lito.toml")
	}

	return ref.Deref(t.config.Proxy.ConfigPath, "lito.toml")
}

func (t *TOML) Load() error {
	panic("unimplemented")
}

func (t *TOML) CreateIfNotExists(...string) error {
	panic("unimplemented")
}

func (t *TOML) Persist() error {
	panic("unimplemented")
}

func (t *TOML) remove() error {
	return os.Remove(t.Path())
}

func (t *TOML) exists() bool {
	_, err := os.Stat(t.Path())
	return !os.IsNotExist(err)
}

func (t *TOML) isEmpty() bool {
	fileInfo, err := os.Stat(t.Path())
	if err != nil {
		t.debug(
			"failed to open config file, this might be a bug, check the error details to confirm",
			logger.Param{Key: "error", Value: err.Error()},
		)
		return true
	}

	return fileInfo.Size() == 0
}

func (t *TOML) debug(msg string, params ...logger.Param) {
	if t.logHandler != nil {
		t.logHandler.Debug(msg, params...)
	}
}
