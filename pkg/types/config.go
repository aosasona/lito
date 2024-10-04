package types

import (
	"encoding/json"
	"fmt"
	"sync"

	toml "github.com/pelletier/go-toml"
	"go.trulyao.dev/lito/pkg/ref"
)

type Config struct {
	Admin    *Admin              `json:"admin,omitempty"    toml:"admin,omitempty"`
	Services map[string]*Service `json:"services,omitempty" toml:"services,omitempty"`
	Proxy    *Proxy              `json:"proxy,omitempty"    toml:"proxy,omitempty"`
	rmu      sync.RWMutex        `json:"-"                  toml:"-"                  mirror:"-"`
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

// String converts the config to the appropriate storage format
func (c *Config) String() string {
	storage := "json"
	if c.Proxy != nil {
		storage = string(ref.Deref(c.Proxy.Storage, "json"))
	}

	var (
		b   []byte
		err error
	)

	switch storage {
	case "json", "memory":
		b, err = c.ToJSON()
	case "toml":
		b, err = c.ToTOML()
	}

	if err != nil {
		fmt.Printf("error converting config to %s format: %s\n", storage, err)
		return ""
	}

	return string(b)
}

// ToJSON converts the config to a JSON byte array
func (c *Config) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c *Config) ToTOML() ([]byte, error) {
	return toml.Marshal(c)
}
