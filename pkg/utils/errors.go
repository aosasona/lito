package utils

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownStorageType = fmt.Errorf("unknown storage type")
	ErrNotImplemented     = fmt.Errorf("not implemented")
	ErrNoConfigSpecified  = errors.New("no config specified")
)
