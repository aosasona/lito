package types

import (
	"strconv"
	"strings"

	"go.trulyao.dev/lito/pkg/ref"
)

type Service struct {
	// TargetHost is the host that the service will forward to, this should be in the format of scheme://host
	TargetHost *string `json:"host,omitempty" toml:"host"`

	// TargetPort is the port that the service will forward to
	TargetPort *int `json:"port,omitempty" toml:"port"`

	// TargetPath is the path that the service will forward to; this will be appended to the request path
	TargetPath *string `json:"path,omitempty" toml:"path"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS *bool `json:"enable_tls,omitempty" toml:"enable_tls"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains,omitempty" toml:"domains,omitempty"`

	// StripHeaders is a list of headers that will be stripped from the request before forwarding
	StripHeaders *[]string `json:"strip_headers,omitempty" toml:"strip_headers"`
}

var DefaultService = Service{
	TargetHost: ref.Ref("http://localhost"),
	TargetPort: ref.Ref(4321),
	TargetPath: nil,
	EnableTLS:  ref.Ref(true),
	Domains: []Domain{
		{DomainName: "example.test"},
	},
	StripHeaders: nil,
}

func (s *Service) GetTargetHost() string {
	if s.TargetHost == nil {
		return ""
	}

	var host, port string

	// Append port if port is not a common port (80, 443)
	if s.TargetPort != nil && *s.TargetPort != 80 && *s.TargetPort != 443 {
		port = ":" + strconv.Itoa(*s.TargetPort)
	}

	host = strings.TrimSuffix(ref.Deref(s.TargetHost, ""), "/") + port

	// Append path to host
	if s.TargetPath != nil {
		path := ref.Deref(s.TargetPath, "/")
		if strings.HasPrefix(path, "/") {
			host += path
		} else {
			host += "/" + path
		}
	}

	// Remove trailing slash
	host = strings.TrimSuffix(host, "/")

	// If it is already a full URL, return it
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") ||
		strings.HasPrefix(host, "ws://") ||
		strings.HasPrefix(host, "wss://") {
		return host
	}

	// Otherwise, assume it is http
	return "http://" + host
}
