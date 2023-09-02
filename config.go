package lito

type Service struct {
	// TargetHost is the host that the service will forward to.
	TargetHost string `json:"host"`

	// TargetPort is the port that the service will forward to.
	TargetPort int `json:"port"`

	// EnableTLS is a flag to enable TLS for the service.
	EnableTLS bool `json:"enable_tls"`
}

type Config struct {
	Admin    AdminConfig        `json:"admin"`
	Services map[string]Service `json:"services"`
}

func LoadFromDirectory(path string) (*Config, error) {
	return nil, nil
}
