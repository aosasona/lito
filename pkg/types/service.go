package types

type Service struct {
	// TargetHost is the host that the service will forward to
	TargetHost string `json:"host"`

	// TargetPort is the port that the service will forward to
	TargetPort int `json:"port,omitempty"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS bool `json:"enable_tls"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains"`

	// StripHeaders is a list of headers that will be stripped from the request before forwarding
	StripHeaders []string `json:"strip_headers"`
}
