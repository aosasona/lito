package storage

import (
	"fmt"
	"os"

	"go.trulyao.dev/lito/pkg/utils"
)

type JSON struct {
	path string
}

func NewJSONStorage() *JSON {
	return &JSON{
		path: s.instance.GetProxyConfig().ConfigPath,
	}
}

func (j *JSON) Path() string { return j.path }

func (j *JSON) Load() error {
	if !j.Exists() || j.Empty() {
		err := j.CreateFile()
		if err != nil {
			return err
		}
	}

	s.instance.Lock()
	defer s.instance.Unlock()

	cByte, err := os.ReadFile(j.path)
	if err != nil {
		return utils.ErrUnableToReadConfig
	}

	config, err := s.instance.UnmarshalConfig(cByte)
	if err != nil {
		return utils.ErrFailedToParseJSONToConfig
	}

	err = s.instance.UpdateConfig(config)
	if err != nil {
		return err
	}

	return nil
}

func (j *JSON) Empty() bool {
	f, err := os.ReadFile(j.path)
	if err != nil {
		return true
	}

	if len(f) == 0 {
		return true
	}

	return false
}

func (j *JSON) Exists() bool {
	f, err := os.Stat(j.path)
	if err != nil {
		return false
	}

	if f.IsDir() {
		return false
	}

	return true
}

func (j *JSON) CreateFile() error {
	s.instance.Lock()
	defer s.instance.Unlock()

	config, err := s.instance.MarshalConfig()
	if err != nil {
		return err
	}

	if err := os.WriteFile(j.path, []byte(config), 0644); err != nil {
		s.instance.GetLogHandler().Fatal(fmt.Sprintf("failed to create config file: %s", err.Error()))
	}

	return nil
}

func (j *JSON) Persist() error {
	s.instance.Lock()
	defer s.instance.Unlock()

	s.instance.GetLogHandler().Info(fmt.Sprintf("Persisting config to %s", j.path))
	config, err := s.instance.MarshalConfig()
	if err != nil {
		return err
	}

	if config == "" {
		return utils.ErrEmptyConfig
	}

	if err := os.WriteFile(j.path, []byte(config), 0644); err != nil {
		return err
	}

	return nil
}
