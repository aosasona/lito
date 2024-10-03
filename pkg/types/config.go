package types

import (
	"encoding/json"
	"sync"
)

type Config struct {
	Admin    *Admin              `json:"admin,omitempty"`
	Services map[string]*Service `json:"services,omitempty"`
	Proxy    *Proxy              `json:"proxy,omitempty"`
	rmu      sync.RWMutex        `json:"-"                  mirror:"-"`
}

func (c *Config) Lock() { c.rmu.Lock() }

func (c *Config) Unlock() { c.rmu.Unlock() }

func (c *Config) RLock() { c.rmu.RLock() }

func (c *Config) RUnlock() { c.rmu.RUnlock() }

func (c *Config) WithLock(f func()) {
	c.Lock()
	defer c.Unlock()

	f()
}

// Update updates the config with the new config without overwriting the mutex - only use this if you know what you're doing
func (c *Config) Update(config *Config) {
	c.Admin = config.Admin
	c.Services = config.Services
	c.Proxy = config.Proxy
}

// String converts the config to a JSON string - should only be used for debugging, handle errors properly when persisting
func (c *Config) String() string {
	b, err := c.ToJSON()
	if err != nil {
		return ""
	}

	return string(b)
}

// ToJSON converts the config to a JSON byte array
func (c *Config) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
