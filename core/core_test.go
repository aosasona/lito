package core

import "go.trulyao.dev/lito/pkg/types"

func mockConfig(configPath ...string) types.Config {
	mockConfigPath := "lito.json"
	if len(configPath) > 0 {
		mockConfigPath = configPath[0]
	}

	return types.Config{
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
				TargetHost: "something.dp.com",
				TargetPort: 80,
				EnableTLS:  false,
				Domains: []types.Domain{
					{
						DomainName: "demo.com",
						Status: types.DomainStatus{
							DNS: types.DomainStatusDNS{Value: types.DNSStatusVerified, LastUpdated: 1696336327},
						},
					},
				},
			},

			"demo2": {
				TargetHost: "anotherdemo.com",
				TargetPort: 80,
				EnableTLS:  false,
				Domains: []types.Domain{
					{
						DomainName: "addftp.com",
						Status: types.DomainStatus{
							DNS: types.DomainStatusDNS{Value: types.DNSStatusVerified, CurrentRetryCount: 0, LastUpdated: 1696336327},
						},
					},
					{
						DomainName: "seconddemo.ai",
						Status: types.DomainStatus{
							DNS: types.DomainStatusDNS{Value: types.DNSStatusPending, CurrentRetryCount: 0, LastUpdated: 1696336327},
						},
					},
				},
			},
		},
	}
}
