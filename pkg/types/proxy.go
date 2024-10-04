package types

import (
	"go.trulyao.dev/lito/pkg/ref"
)

type Storage string

const (
	StorageMemory Storage = "memory"
	StorageJSON   Storage = "json"
	StorageTOML   Storage = "toml"
)

type Proxy struct {
	// Host is the host that the proxy will listen on
	Host *string `json:"host,omitempty" toml:"host,omitempty"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort *int `json:"http_port,omitempty" toml:"http_port,omitempty"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort *int `json:"https_port,omitempty" toml:"https_port,omitempty"`

	// EnableTLS is a flag that determines whether or not the proxy will listen for TLS connections
	EnableTLS *bool `json:"enable_tls,omitempty" toml:"enable_tls,omitempty"`

	// TLSEmail is the email address that will be used to register TLS certificates
	TLSEmail *string `json:"tls_email,omitempty" toml:"tls_email,omitempty"`

	// EnableHTTPSRedirect is a flag that determines whether or not the proxy will automatically redirect HTTP requests to HTTPS
	EnableHTTPSRedirect *bool `json:"enable_https_redirect,omitempty" toml:"enable_https_redirect,omitempty"`

	// ConfigPath is the path to the file that the proxy will use to store its configuration - create if not exists or load
	ConfigPath *string `json:"config_path,omitempty" toml:"config_path,omitempty"`

	// Storage is the type of store that the proxy will use
	Storage *Storage `json:"storage,omitempty" mirror:"type:'json' | 'sqlite' | 'memory'" toml:"storage,omitempty"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames *[]string `json:"cnames,omitempty" toml:"cnames,omitempty"`
}

var DefaultProxy = Proxy{
	Host:                ref.Ref("0.0.0.0"),
	HTTPPort:            ref.Ref(80),
	HTTPSPort:           ref.Ref(443),
	EnableTLS:           ref.Ref(false),
	TLSEmail:            ref.Ref(""),
	EnableHTTPSRedirect: ref.Ref(false),
	ConfigPath:          ref.Ref("lito.toml"),
	Storage:             ref.Ref(StorageTOML),
	CNames:              ref.Ref([]string{}),
}
