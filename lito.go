package lito

type Lito struct {
	config *Config
}

type Opts struct {
	// Config is the configuration for the proxy, if not loading from the filesystem.
	Config *Config

	// DataDir is the directory to load and store configuration from, will be created if it does not exist.
	ProxyDir string

	// RootFile is the file to load and store admin configuration from.
	RootFilePath string
}

func New(opts *Opts) (*Lito, error) {
	if opts.Config == nil && opts.ProxyDir == "" {
		return nil, ErrNoConfigSpecified
	}
	return &Lito{}, nil
}

func (l *Lito) Run() error {
	return nil
}
