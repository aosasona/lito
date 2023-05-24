package config

type LitoConfig struct {
	Admin     AdminConfig     `mapstructure:"admin"`
	Proxy     ProxyConfig     `mapstructure:"proxy"`
	Routes    RoutesConfig    `mapstructure:"proxy.routes"`
	Discovery DiscoveryConfig `mapstructure:"proxy.discovery"`
}

type AdminConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Path     string `mapstructure:"path"`
}

type ProxyConfig struct {
	Host      string `mapstructure:"host"`
	Ports     []int  `mapstructure:"ports"`
	LogDir    string `mapstructure:"log_dir"`
	EnableTLS bool   `mapstructure:"enable_tls"`
}

type RoutesConfig struct {
	Data            string         `mapstructure:"data"`
	StorageDriver   string         `mapstructure:"storage_driver"`
	RefreshInterval int            `mapstructure:"refresh_interval"`
	NotFound        NotFoundConfig `mapstructure:"not_found"`
}

type NotFoundConfig struct {
	Type    string `mapstructure:"type"`
	Content string `mapstructure:"content"`
}

type DiscoveryConfig struct {
	Path            string `mapstructure:"path"`
	AllowExternal   bool   `mapstructure:"allow_external"`
	RefreshInterval int    `mapstructure:"refresh_interval"`
}

func parseLitoConfig() (*LitoConfig, error) {
	var config LitoConfig

	return &config, nil
}
