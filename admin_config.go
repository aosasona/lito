package lito

type AdminConfig struct {
	// EnableAPI is a flag to enable the admin API.
	EnableAPI bool `json:"enable_api"`

	// Port is the port that the admin API will listen on.
	Port int `json:"port"`
}
