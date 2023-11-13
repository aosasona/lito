package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.trulyao.dev/lito/pkg/logger"
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

// We can unwrap here because it is guaranteed that the config is not nil
func (j *JSON) Path() string {
	return j.config.
		Proxy.Unwrap().
		ConfigPath.Unwrap()
}

// Load reads the config from disk and loads it into memory, creating it if it doesn't exist yet
func (j *JSON) Load() error {
	if !j.exists() || j.isEmpty() {
		if err := j.init(); err != nil {
			return err
		}
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

	configBytes, err := j.config.ToJson()
	if err != nil {
		return fmt.Errorf("failed to convert config to JSON: %s", err.Error())
	}

	var file *os.File

	if file, err = os.OpenFile(j.Path(), os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
		return fmt.Errorf("failed to open config file: %s", err.Error())
	}

	if _, err = file.Write(configBytes); err != nil {
		return fmt.Errorf("failed to write config to disk: %s", err.Error())
	}

	j.debug("successfully persisted config to disk")

	return nil
}

func (j *JSON) read() (*types.Config, error) {
	configBytes, err := os.ReadFile(j.Path())
	if err != nil {
		return nil, err
	}

	config := new(types.Config)
	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// init() creates the config file on disk using the current config in memory
//
// This function should only be used when the config file doesn't exist yet or is empty (e.g. on first run)
func (j *JSON) init() error {
	j.config.Lock()
	defer j.config.Unlock()

	err := os.MkdirAll(filepath.Dir(j.Path()), 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(j.Path()); os.IsNotExist(err) {
		file, err := os.Create(j.Path())
		if err != nil {
			return err
		}
		defer file.Close()

		configBytes, err := j.config.ToJson()
		if err != nil {
			return fmt.Errorf("failed to convert config to JSON: %w", err)
		}

		_, err = file.Write(configBytes)
		if err != nil {
			return err
		}
	}

	return nil
}

// remove() deletes the config file on disk
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
		j.debug("failed to open config file, this might be a bug, check the error details to confirm", logger.Param{Key: "error", Value: err.Error()})
		return true
	}

	return fileInfo.Size() == 0
}

func (j *JSON) debug(msg string, params ...logger.Param) {
	if j.logHandler != nil {
		j.logHandler.Debug(msg, params...)
	}
}
