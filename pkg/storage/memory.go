package storage

type Memory struct{}

func NewMemoryStorage() *Memory {
	return &Memory{}
}

func (m *Memory) Path() string { return "" }

func (m *Memory) Load() error { return nil }

func (m *Memory) Persist() error {
	s.instance.GetLogHandler().Warn("Storage is set to memory, skipping config persistence - this is NOT recommended for production use")
	return nil
}
