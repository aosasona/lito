package types

import "go.trulyao.dev/lito/ext/option"

type Service struct {
	// TargetHost is the host that the service will forward to
	TargetHost option.String `json:"host"`

	// TargetPort is the port that the service will forward to
	TargetPort option.Int `json:"port,omitempty"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS option.Bool `json:"enable_tls"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains"`

	// StripHeaders is a list of headers that will be stripped from the request before forwarding
	StripHeaders option.Option[[]string] `json:"strip_headers"`
}
