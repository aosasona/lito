package types

import (
	"strings"

	"go.trulyao.dev/lito/ext/option"
)

type Service struct {
	// TargetHost is the host that the service will forward to, this should be in the format of scheme://host
	TargetHost option.String `json:"host"`

	// TargetPort is the port that the service will forward to
	TargetPort option.Int `json:"port,omitempty"`

	// TargetPath is the path that the service will forward to; this will be appended to the request path
	TargetPath option.String `json:"path,omitempty"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS option.Bool `json:"enable_tls"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains"`

	// StripHeaders is a list of headers that will be stripped from the request before forwarding
	StripHeaders option.Option[[]string] `json:"strip_headers"`
}

func (s *Service) GetTargetHost() string {
	if s.TargetHost.IsNone() {
		return ""
	}

	var port string
	if !s.TargetPort.IsNone() {
		port = ":" + s.TargetPort.String()
	}

	host := strings.TrimSuffix(s.TargetHost.Unwrap(), "/") + port
	if !s.TargetPath.IsNone() {
		path := s.TargetPath.Unwrap()
		if strings.HasPrefix(path, "/") {
			host += path
		} else {
			host += "/" + path
		}
	}

	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "ws://") || strings.HasPrefix(host, "wss://") {
		return host
	}

	return "http://" + host
}
