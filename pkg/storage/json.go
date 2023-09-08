package storage

type JSON struct {
	FilePath string
}

func NewJSONStorage(path string) *JSON {
	return &JSON{
		FilePath: path,
	}
}

func (j *JSON) Path() string { return j.FilePath }

func (j *JSON) Load() error {
	s.instance.Lock()
	defer s.instance.Unlock()
	return nil
}

func (j *JSON) Persist() error {
	return nil
}
