package types

import (
	"encoding/json"
	"sync"
)

type Config struct {
	Admin    *Admin              `json:"admin"`
	Services map[string]*Service `json:"services"`
	Proxy    *Proxy              `json:"proxy"`
	sync.RWMutex
}

func (c *Config) Lock() {
	c.Lock()
}

func (c *Config) Unlock() {
	c.Unlock()
}

// ToJson converts the config to a JSON byte array
func (c *Config) ToJson() ([]byte, error) {
	return json.Marshal(c)
}
