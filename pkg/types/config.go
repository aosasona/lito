package types

import "sync"

type Config struct {
	Admin    Admin              `json:"admin"`
	Services map[string]Service `json:"services"`
	Proxy    Proxy              `json:"proxy"`
	mutex    sync.RWMutex
}

func (c *Config) Lock() {
	c.mutex.Lock()
}

func (c *Config) Unlock() {
	c.mutex.Unlock()
}
