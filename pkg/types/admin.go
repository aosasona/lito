package types

import (
	"go.trulyao.dev/lito/pkg/ref"
)

type Admin struct {
	// EnableAPI is a flag to enable the admin API
	Enabled *bool `json:"enabled,omitempty" mirror:"type:boolean"`

	// Port is the port that the admin API will listen on
	Port *int `json:"port,omitempty" mirror:"type:number"`

	// APIKey is the key that the admin API will use for authentication
	APIKey *string `json:"api_key,omitempty" mirror:"type:string"`
}

var DefaultAdmin = Admin{
	Enabled: ref.Ref(false),
	Port:    nil,
	APIKey:  nil,
}
