package lito

type Storage string

const (
	StorageJSON Storage = "json"
	// Not yet implemented
	StorageSQLite3 Storage = "sqlite"
)

type Service struct {
	// TargetHost is the host that the service will forward to
	TargetHost string `json:"host"`

	// TargetPort is the port that the service will forward to
	TargetPort int `json:"port"`

	// EnableTLS is a flag to enable TLS for the service
	EnableTLS bool `json:"enable_tls"`

	// Domains is a list of domains that the service will respond to
	Domains []Domain `json:"domains"`
}

type Proxy struct {
	// Host is the host that the proxy will listen on
	Host string `json:"host"`

	// HTTPPort is the port that the proxy will listen on for HTTP connections
	HTTPPort int `json:"port"`

	// HTTPSPort is the port that the proxy will listen on for TLS connections
	HTTPSPort int `json:"tls_port"`

	// DataDir is the directory that the proxy will store data in
	DataDir string `json:"data_dir"`

	// Storage is the type of store that the proxy will use
	Storage Storage `json:"storage" ts:"type:'json' | 'sqlite'"`

	// CNames is a list of CNAMEs associated with the proxy/host running Lito
	CNames []string `json:"cnames"`
}

type Config struct {
	Admin    Admin              `json:"admin"`
	Services map[string]Service `json:"services"`
	Proxy    Proxy              `json:"proxy"`
}

func LoadFromDirectory(path string) (*Config, error) {
	return nil, nil
}
