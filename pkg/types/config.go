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

// String converts the config to a JSON string - should only be used for debugging, handle errors properly when persisting
func (c *Config) String() string {
	b, err := c.ToJson()
	if err != nil {
		return ""
	}

	return string(b)
}

// ToJson converts the config to a JSON byte array
func (c *Config) ToJson() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
