package lito

type Admin struct {
	// EnableAPI is a flag to enable the admin API
	Enabled bool `json:"enabled"`

	// Port is the port that the admin API will listen on
	Port int `json:"port"`

	// APIKey is the key that the admin API will use for authentication
	APIKey string `json:"api_key"`
}
