package types

import "go.trulyao.dev/lito/ext/option"

// type Admin struct {
// 	// EnableAPI is a flag to enable the admin API
// 	Enabled bool `json:"enabled" ts:"type:boolean"`
//
// 	// Port is the port that the admin API will listen on
// 	Port int `json:"port" ts:"type:number"`
//
// 	// APIKey is the key that the admin API will use for authentication
// 	APIKey string `json:"api_key" ts:"type:string"`
// }

type Admin struct {
	// EnableAPI is a flag to enable the admin API
	Enabled option.Bool `json:"enabled,omitempty" ts:"type:boolean"`

	// Port is the port that the admin API will listen on
	Port option.Int `json:"port,omitempty" ts:"type:number"`

	// APIKey is the key that the admin API will use for authentication
	APIKey option.String `json:"api_key,omitempty" ts:"type:string"`
}
