package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/types"
)

type JSON struct {
	config     *types.Config
	logHandler logger.Logger
}

func NewJSONStorage(opts *Opts) (*JSON, error) {
	if opts == nil {
		return nil, fmt.Errorf("opts provided in NewJSONStorage cannot be nil")
	}

	if opts.Config == nil {
		return nil, fmt.Errorf("config provided in NewJSONStorage cannot be nil")
	}

	return &JSON{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}, nil
}

func (j *JSON) IsWatchchable() bool { return true }

// Get the path or set the default path and return that
func (j *JSON) Path() string {
	if j.config.Proxy == nil {
		j.config.Proxy = &types.DefaultProxy
		j.config.Proxy.ConfigPath = ref.Ref("lito.json")
	}

	return ref.Deref(j.config.Proxy.ConfigPath, "lito.json")
}

// Create the JSON config if it doesn't already exist
func (j *JSON) CreateIfNotExists(optPath ...string) error {
	// If an alternative path is provided use, that as our new config path and create the config if it doesn't already exist
	if len(optPath) > 0 && strings.TrimSpace(optPath[0]) != "" {
		if j.config.Proxy == nil {
			j.config.Proxy = &types.DefaultProxy
		}

		j.config.Proxy.ConfigPath = ref.Ref(optPath[0])
	}

	if !j.exists() || j.isEmpty() {
		if err := j.init(); err != nil {
			return err
		}
	}

	return nil
}

// Load reads the config from disk and loads it into memory, creating it if it doesn't exist yet
func (j *JSON) Load() error {
	if err := j.CreateIfNotExists(); err != nil {
		return err
	}

	j.config.Lock()
	defer j.config.Unlock()

	config, err := j.read()
	if err != nil {
		return fmt.Errorf("failed to read config from disk: %w", err)
	}

	j.config.Update(config)
	j.debug("successfully loaded config from disk")

	return nil
}

// Persist writes the current config to disk
func (j *JSON) Persist() error {
	j.config.RLock()
	defer j.config.RUnlock()

	configBytes, err := j.config.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert config to JSON: %s", err.Error())
	}

	if err = os.WriteFile(j.Path(), configBytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write config to disk: %s", err.Error())
	}

	j.debug("successfully persisted config to disk", logger.Param{Key: "path", Value: j.Path()})
	return nil
}

func (j *JSON) read() (*types.Config, error) {
	configBytes, err := os.ReadFile(j.Path())
	if err != nil {
		return &types.Config{}, err
	}

	config := new(types.Config)
	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return &types.Config{}, err
	}

	return config, nil
}

// init() creates the config file on disk using the current config in memory
func (j *JSON) init() error {
	j.config.Lock()
	defer j.config.Unlock()

	if err := os.MkdirAll(filepath.Dir(j.Path()), os.ModePerm); err != nil {
		return err
	}

	configBytes, err := j.config.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert config to JSON: %w", err)
	}

	if err = os.WriteFile(j.Path(), configBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (j *JSON) remove() error {
	return os.Remove(j.Path())
}

func (j *JSON) exists() bool {
	_, err := os.Stat(j.Path())
	return !os.IsNotExist(err)
}

func (j *JSON) isEmpty() bool {
	fileInfo, err := os.Stat(j.Path())
	if err != nil {
		j.debug(
			"failed to open config file, this might be a bug, check the error details to confirm",
			logger.Param{Key: "error", Value: err.Error()},
		)
		return true
	}

	return fileInfo.Size() == 0
}

func (j *JSON) debug(msg string, params ...logger.Param) {
	if j.logHandler != nil {
		j.logHandler.Debug(msg, params...)
	}
}
