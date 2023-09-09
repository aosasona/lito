package types

type Admin struct {
	// EnableAPI is a flag to enable the admin API
	Enabled bool `json:"enabled" ts:"type:boolean"`

	// Port is the port that the admin API will listen on
	Port int `json:"port" ts:"type:number"`

	// APIKey is the key that the admin API will use for authentication
	APIKey string `json:"api_key" ts:"type:string"`

	// ConfirmDisable is a flag to require confirmation when disabling the admin API so that it cannot be disabled by accident (Go sets bools to false by default)
	ConfirmDisable bool `json:"confirm_disable,omitempty" ts:"type:boolean"`
}
