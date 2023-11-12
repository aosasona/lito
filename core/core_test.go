package core

import (
	"go.trulyao.dev/lito/ext/option"
	"go.trulyao.dev/lito/pkg/types"
)

var (
	String = option.StringValue
	Int    = option.IntValue
	Bool   = option.BoolValue
)

func mockConfig(configPath ...string) types.Config {
	mockConfigPath := "lito.json"
	if len(configPath) > 0 {
		mockConfigPath = configPath[0]
	}

	return types.Config{
		Admin: option.Some(&types.Admin{
			Enabled: Bool(true),
			Port:    Int(9090),
			APIKey:  String("1234567890abcdefghij"),
		}),
		Proxy: option.Some(&types.Proxy{
			Host:                String("0.0.0.0"),
			HTTPPort:            Int(80),
			HTTPSPort:           Int(443),
			EnableTLS:           Bool(true),
			TLSEmail:            String("someone@example.com"),
			EnableHTTPSRedirect: Bool(true),
			ConfigPath:          String(mockConfigPath),
			Storage:             option.Some(types.StorageJSON),
			CNames:              option.Some([]string{"example.com"}),
		}),
		Services: map[string]*types.Service{
			"demo": {
				TargetHost: String("https://something.dp.com"),
				TargetPort: Int(80),
				EnableTLS:  Bool(false),
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
				TargetHost: String("https://anotherdemo.com"),
				TargetPort: Int(80),
				EnableTLS:  Bool(false),
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
