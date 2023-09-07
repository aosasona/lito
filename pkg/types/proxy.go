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
	Host string `json:"host,omitempty"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort int `json:"http_port,omitempty"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort int `json:"https_port,omitempty"`

	// DataDir is the directory that the proxy will store data in
	DataDir string `json:"data_dir,omitempty"`

	// Storage is the type of store that the proxy will use
	Storage Storage `json:"storage" ts:"type:'json' | 'sqlite'"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames []string `json:"cnames"`
}
