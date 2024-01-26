package core

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

func (c *Core) watchConfig(waitChan chan os.Signal) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		c.logHandler.Error("failed to create watcher", logger.Field("error", err))
		return
	}
	defer watcher.Close()

	if c.config.Proxy.IsNone() {
		c.logHandler.Info("no proxy config found, skipping config watcher")
		return
	}

	if c.config.Proxy.Unwrap(&types.DefaultProxy).ConfigPath.IsNone() {
		c.logHandler.Info("no config path found, skipping config watcher")
		return
	}

	path := c.config.Proxy.Unwrap(&types.DefaultProxy).ConfigPath.Unwrap("config.json")

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				c.logHandler.Info(fmt.Sprintf("config file changed, reloading config from %s", path), logger.Field("event", event))
				if err := c.storageHandler.Load(); err != nil {
					c.logHandler.Error(fmt.Sprintf("failed to reload config from %s", path), logger.Field("error", err))
					return
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				c.logHandler.Error(fmt.Sprintf("failed to watch config file at %s", path), logger.Field("error", err))
			}
		}
	}()

	if err = watcher.Add(path); err != nil {
		c.logHandler.Error(fmt.Sprintf("failed to watch config file at %s", path), logger.Field("error", err))
		return
	}

	c.logHandler.Info(fmt.Sprintf("watching config file at %s", path))

	<-waitChan
}
