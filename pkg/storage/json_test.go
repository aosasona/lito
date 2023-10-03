package storage

import (
	"reflect"
	"testing"

	"go.trulyao.dev/lito/pkg/types"
)

var (
	mockConfigPath = ".mock.config.json"

	mockConfigBytes = []byte(`{
  "admin": {
    "enabled": true,
    "port": 9090,
    "api_key": "1234567890abcdefghij"
  },
  "services": {
    "demo": {
      "host": "demo.com",
      "port": 80,
      "enable_tls": false,
      "domains": [
        {
          "name": "sub.demo.com",
          "status": {
            "cert": {
              "value": "pending",
              "last_updated": 1696336327
            },
            "dns": {
              "value": "verified",
              "retry": {
                "max_tries": 0,
                "interval": 0,
                "current_tries": 0
              },
              "last_updated": 1696336327
            }
          }
        }
      ],
      "strip_headers": null
    }
  },
  "proxy": {
    "host": "0.0.0.0",
    "http_port": 80,
    "https_port": 443,
    "enable_tls": true,
    "tls_email": "someone@example.com",
    "enable_https_redirect": true,
    "config_path": ".mock.config.json",
    "storage": "json",
    "cnames": [
      "example.com"
    ]
  }
}`)

	jsonOpts = &Opts{
		Config: &types.Config{
			Admin: &types.Admin{
				Enabled: true,
				Port:    9090,
				APIKey:  "1234567890abcdefghij",
			},
			Proxy: &types.Proxy{
				Host:                "0.0.0.0",
				HTTPPort:            80,
				HTTPSPort:           443,
				EnableTLS:           true,
				TLSEmail:            "someone@example.com",
				EnableHTTPSRedirect: true,
				ConfigPath:          mockConfigPath,
				Storage:             types.StorageJSON,
				CNames:              []string{"example.com"},
			},
			Services: map[string]*types.Service{
				"demo": {
					TargetHost: "demo.com",
					TargetPort: 80,
					EnableTLS:  false,
					Domains: []types.Domain{
						{
							DomainName: "sub.demo.com",
							Status: types.DomainStatus{
								Cert: types.DomainStatusCert{
									Value:       types.CertStatusPending,
									LastUpdated: 1696336327,
								},
								DNS: types.DomainStatusDNS{
									Value:       types.DNSStatusVerified,
									LastUpdated: 1696336327,
								},
							},
						},
					},
				},
			},
		},
	}
)

func Test_InitOnDisk(t *testing.T) {
	j := NewJSONStorage(jsonOpts)

	// try to init on disk
	err := j.init()
	if err != nil {
		t.Errorf("failed to init on disk: %v", err)
	}

	diskContent, err := j.read()
	if err != nil {
		t.Errorf("failed to read from disk: %v", err)
	}

	dcStr, err := diskContent.ToJson()
	if err != nil {
		t.Errorf("failed to convert disk content to string: %v", err)
	}

	if !reflect.DeepEqual(dcStr, mockConfigBytes) {
		t.Errorf("expected disk content to be %v, got %v", string(mockConfigBytes), string(dcStr))
	}

	// clean up
	err = j.remove()
	if err != nil {
		t.Errorf("failed to clean up: %v", err)
	}
}

func Test_Load(t *testing.T) {
	// init mock config on disk so we can load it again and compare
	tempJ := NewJSONStorage(jsonOpts)
	err := tempJ.init()
	if err != nil {
		t.Errorf("failed to init on disk: %v", err)
	}

	j := NewJSONStorage(&Opts{Config: &types.Config{Proxy: &types.Proxy{ConfigPath: mockConfigPath}}})

	err = j.Load()
	if err != nil {
		t.Errorf("failed to load from disk: %v", err)
	}

	memContentBytes, err := j.config.ToJson()
	if err != nil {
		t.Errorf("failed to convert config to JSON bytes: %v", err)
	}

	if !reflect.DeepEqual(memContentBytes, mockConfigBytes) {
		t.Errorf("expected config to be %v, got %v", string(mockConfigBytes), string(memContentBytes))
	}

	// clean up
	err = j.remove()
	if err != nil {
		t.Errorf("failed to clean up: %v", err)
	}
}
