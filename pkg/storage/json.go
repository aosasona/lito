package storage

import "os"

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
	s.instance.Lock()
	defer s.instance.Unlock()

	return nil
}

func (j *JSON) Persist() error {
	s.instance.Lock()
	defer s.instance.Unlock()

	config, err := s.instance.String()
	if err != nil {
		return err
	}

	err = os.WriteFile(j.path, []byte(config), 0644)
	if err != nil {
		return err
	}

	return nil
}
