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
	Host option.String `json:"host,omitempty" ts:"type:string"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort option.Int `json:"http_port,omitempty" ts:"type:number"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort option.Int `json:"https_port,omitempty" ts:"type:number"`

	// EnableTLS is a flag that determines whether or not the proxy will listen for TLS connections
	EnableTLS option.Bool `json:"enable_tls,omitempty" ts:"type:boolean"`

	// TLSEmail is the email address that will be used to register TLS certificates
	TLSEmail option.String `json:"tls_email,omitempty" ts:"type:string"`

	// EnableHTTPSRedirect is a flag that determines whether or not the proxy will automatically redirect HTTP requests to HTTPS
	EnableHTTPSRedirect option.Bool `json:"enable_https_redirect,omitempty" ts:"type:boolean"`

	// ConfigPath is the path to the file that the proxy will use to store its configuration - create if not exists or load
	ConfigPath option.String `json:"config_path,omitempty" ts:"type:string"`

	// Storage is the type of store that the proxy will use
	Storage option.Option[Storage] `json:"storage,omitempty" ts:"type:'json' | 'sqlite' | 'memory'"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames option.Option[[]string] `json:"cnames,omitempty" ts:"type:string[]"`
}
