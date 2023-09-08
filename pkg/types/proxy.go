package types

type Storage string

const (
	StorageMemory Storage = "memory"
	StorageJSON   Storage = "json"
	// Not yet implemented
	StorageSQLite3 Storage = "sqlite"
)

type Proxy struct {
	// Host is the host that the proxy will listen on
	Host string `json:"host"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort int `json:"http_port"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort int `json:"https_port,omitempty"`

	// EnableTLS is a flag that determines whether or not the proxy will listen for TLS connections
	EnableTLS bool `json:"enable_tls,omitempty"`

	// TLSEmail is the email address that will be used to register TLS certificates
	TLSEmail string `json:"tls_email"`

	// ConfigPath is the path to the file that the proxy will use to store its configuration - create if not exists or load
	ConfigPath string `json:"data_dir"`

	// Storage is the type of store that the proxy will use
	Storage Storage `json:"storage" ts:"type:'json' | 'sqlite'"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames []string `json:"cnames"`
}
