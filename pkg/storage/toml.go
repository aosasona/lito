package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
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
	if err := t.CreateIfNotExists(); err != nil {
		return err
	}

	t.config.Lock()
	defer t.config.Unlock()

	config, err := t.read()
	if err != nil {
		return fmt.Errorf("failed to read config from disk: %w", err)
	}

	t.config.Update(config)
	t.debug("successfully loaded config from disk", logger.Param{Key: "path", Value: t.Path()})

	return nil
}

func (t *TOML) CreateIfNotExists(optPath ...string) error {
	// If an alternative path is provided use, that as our new config path and create the config if it doesn't already exist
	if len(optPath) > 0 && strings.TrimSpace(optPath[0]) != "" {
		if t.config.Proxy == nil {
			t.config.Proxy = &types.DefaultProxy
		}

		t.config.Proxy.ConfigPath = ref.Ref(optPath[0])
	}

	if !t.exists() || t.isEmpty() {
		if err := t.init(); err != nil {
			return err
		}
	}

	return nil
}

func (t *TOML) Persist() error {
	t.config.Lock()
	defer t.config.Unlock()

	configBytes, err := t.config.ToTOML()
	if err != nil {
		return fmt.Errorf("failed to convert config to JSON: %s", err.Error())
	}

	if err = os.WriteFile(t.Path(), configBytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write config to disk: %s", err.Error())
	}

	t.debug("successfully persisted config to disk", logger.Param{Key: "path", Value: t.Path()})
	return nil
}

func (t *TOML) read() (*types.Config, error) {
	configBytes, err := os.ReadFile(t.Path())
	if err != nil {
		return &types.Config{}, err
	}

	config := new(types.Config)
	if err = toml.Unmarshal(configBytes, config); err != nil {
		return &types.Config{}, err
	}

	return config, nil
}

// init() creates the config file on disk using the current config in memory
func (t *TOML) init() error {
	t.config.Lock()
	defer t.config.Unlock()

	if err := os.MkdirAll(filepath.Dir(t.Path()), os.ModePerm); err != nil {
		return err
	}

	// Protection	against overwriting existing config files
	var (
		configBytes []byte
		err         error
	)

	if configBytes, err = t.config.ToTOML(); err != nil {
		return err
	}

	if err = os.WriteFile(t.Path(), configBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
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
