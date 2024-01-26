package types

import (
	"strings"

	"go.trulyao.dev/lito/ext/option"
)

type Service struct {
	// TargetHost is the host that the service will forward to, this should be in the format of scheme://host
	TargetHost option.String `json:"host,omitempty" ts:"type:string"`

	// TargetPort is the port that the service will forward to
	TargetPort option.Int `json:"port,omitempty" ts:"type:number"`

	// TargetPath is the path that the service will forward to; this will be appended to the request path
	TargetPath option.String `json:"path,omitempty" ts:"type:string"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS option.Bool `json:"enable_tls,omitempty" ts:"type:boolean"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains,omitempty" ts:"type:Domain[]"`

	// StripHeaders is a list of headers that will be stripped from the request before forwarding
	StripHeaders option.Option[[]string] `json:"strip_headers,omitempty" ts:"type:string[]"`
}

var DefaultService = Service{
	TargetHost:   option.None[string](),
	TargetPort:   option.None[int](),
	TargetPath:   option.None[string](),
	EnableTLS:    option.None[bool](),
	Domains:      []Domain{},
	StripHeaders: option.None[[]string](),
}

func (s *Service) GetTargetHost() string {
	if s.TargetHost.IsNone() {
		return ""
	}

	var host, port string

	// Append port if port is not a common port (80, 443)
	if s.TargetPort.Unwrap(80) != 80 && s.TargetPort.Unwrap(80) != 443 {
		port = ":" + s.TargetPort.StringWithDefault("")
	}

	host = strings.TrimSuffix(s.TargetHost.Unwrap(""), "/") + port

	// Append path to host
	if !s.TargetPath.IsNone() {
		path := s.TargetPath.Unwrap("/")
		if strings.HasPrefix(path, "/") {
			host += path
		} else {
			host += "/" + path
		}
	}

	// Remove trailing slash
	if strings.HasSuffix(host, "/") {
		host = strings.TrimSuffix(host, "/")
	}

	// If it is already a full URL, return it
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "ws://") || strings.HasPrefix(host, "wss://") {
		return host
	}

	// Otherwise, assume it is http
	return "http://" + host
}
