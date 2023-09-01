package lito

import (
	"fmt"

	"github.com/aosasona/lito/internal/core"
)

type Lito struct {
	core *core.Core
}

type Config struct {
	ConfigDir string
}

var DefaultConfig = Config{
	ConfigDir: ".",
}

func New(c Config) (*Lito, error) {
	litoCore, err := core.New(c.ConfigDir)
	if err != nil {
		return nil, err
	}
	return &Lito{
		core: litoCore,
	}, nil
}

func (l *Lito) watchConfig() error {
	fmt.Println("watching config")
	return nil
}

func (l *Lito) Run() error {
	if l.core.Config().Lito().Proxy.WatchConfig {
		err := l.watchConfig()
		if err != nil {
			fmt.Println("failed to watch config: ", err)
		}
	}
	return nil
}
