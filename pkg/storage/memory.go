package storage

type Memory struct{}

func NewMemoryStorage() *Memory {
	return &Memory{}
}

func (m *Memory) Path() string { return "" }

func (m *Memory) Load() error { return nil }

func (m *Memory) Persist() error { return nil }
