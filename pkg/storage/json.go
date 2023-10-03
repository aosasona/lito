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

func NewJSONStorage(opts *Opts) *JSON {
	return &JSON{
		config:     opts.Config,
		logHandler: opts.LogHandler,
	}
}

func (j *JSON) Path() string {
	return j.config.Proxy.ConfigPath
}

// Load reads the config from disk and loads it into memory, creating it if it doesn't exist yet
func (j *JSON) Load() error {
	if !j.exists() || j.isEmpty() {
		if err := j.init(); err != nil {
			return err
		}
	}

	config, err := j.read()
	if err != nil {
		return fmt.Errorf("failed to read config from disk: %w", err)
	}

	j.config.Update(config)

	return nil
}

// Persist writes the current config to disk
func (j *JSON) Persist() error {
	panic("not implemented")
}

func (j *JSON) read() (*types.Config, error) {
	configBytes, err := os.ReadFile(j.config.Proxy.ConfigPath)
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
	err := os.MkdirAll(filepath.Dir(j.config.Proxy.ConfigPath), 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(j.config.Proxy.ConfigPath); os.IsNotExist(err) {
		file, err := os.Create(j.config.Proxy.ConfigPath)
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
	return os.Remove(j.config.Proxy.ConfigPath)
}

func (j *JSON) exists() bool {
	_, err := os.Stat(j.config.Proxy.ConfigPath)
	return !os.IsNotExist(err)
}

func (j *JSON) isEmpty() bool {
	fileInfo, err := os.Stat(j.config.Proxy.ConfigPath)
	if err != nil {
		j.logHandler.Debug("failed to open config file, this might be a bug, check the error details to confirm", logger.Param{Key: "error", Value: err.Error()})
		return true
	}

	return fileInfo.Size() == 0
}
