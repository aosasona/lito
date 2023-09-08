package storage

import "go.trulyao.dev/lito/pkg/types"

type store struct {
	instance types.Instance
}

var s = store{}

type Storage interface {
	Path() string
	Load() error
	Persist() error
}

// This needs to be setup to track the main instances of the Lito struct fields
func Init(i types.Instance) {
	s.instance = i
}
