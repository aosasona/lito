package types

type Service struct {
	// TargetHost is the host that the service will forward to
	TargetHost string `json:"host,omitempty"`

	// TargetPort is the port that the service will forward to
	TargetPort int `json:"port,omitempty"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS *bool `json:"enable_tls,omitempty"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains"`
}
