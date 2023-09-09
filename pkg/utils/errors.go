package utils

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownStorageType        = fmt.Errorf("unknown storage type")
	ErrNotImplemented            = fmt.Errorf("not implemented")
	ErrNoConfigSpecified         = errors.New("no config specified")
	ErrUnableToReadConfig        = errors.New("unable to read config")
	ErrFailedToParseJSONToConfig = errors.New("failed to parse JSON to config")
)
