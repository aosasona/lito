package lito

import (
	"fmt"

	"go.trulyao.dev/lito/pkg/types"
)

func (l *Lito) persistConfig() error {
	if l.Config.Proxy.Storage == types.StorageMemory {
		l.LogHandler.Warn("Storage is set to memory, skipping config persistence - this is NOT recommended for production use")
		return nil
	}

	l.LogHandler.Info(fmt.Sprintf("Persisting config to %s", l.Config.Proxy.Storage))
	l.Config.mutex.Lock()
	defer l.Config.mutex.Unlock()

	return nil
}

func (l *Lito) persistConfigJSON() error {
	return nil
}
