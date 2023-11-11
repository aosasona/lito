package types

import "go.trulyao.dev/lito/ext/option"

type Storage string

const (
	StorageMemory Storage = "memory"
	StorageJSON   Storage = "json"

	// IMPORTANT: not yet implemented, do not use!
	StorageSQLite3 Storage = "sqlite"
)

type Proxy struct {
	// Host is the host that the proxy will listen on
	Host option.String `json:"host"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort option.Int `json:"http_port"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort option.Int `json:"https_port,omitempty"`

	// EnableTLS is a flag that determines whether or not the proxy will listen for TLS connections
	EnableTLS option.Bool `json:"enable_tls"`

	// TLSEmail is the email address that will be used to register TLS certificates
	TLSEmail option.String `json:"tls_email"`

	// EnableHTTPSRedirect is a flag that determines whether or not the proxy will automatically redirect HTTP requests to HTTPS
	EnableHTTPSRedirect option.Bool `json:"enable_https_redirect"`

	// ConfigPath is the path to the file that the proxy will use to store its configuration - create if not exists or load
	ConfigPath option.String `json:"config_path"`

	// Storage is the type of store that the proxy will use
	Storage option.Option[Storage] `json:"storage" ts:"type:'json' | 'sqlite' | 'memory'"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames option.Option[[]string] `json:"cnames"`
}
