package storage

import (
	"reflect"
	"testing"

	"go.trulyao.dev/lito/pkg/ref"
	"go.trulyao.dev/lito/pkg/types"
)

var (
	String = ref.Ref[string]
	Int    = ref.Ref[int]
	Bool   = ref.Ref[bool]
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
      "path": null,
      "enable_tls": false,
      "domains": [
        {
          "name": "sub.demo.com",
          "status": {
            "dns": {
              "value": "verified",
              "current_retry_count": 0,
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
				Enabled: Bool(true),
				Port:    Int(9090),
				APIKey:  String("1234567890abcdefghij"),
			},
			Proxy: &types.Proxy{
				Host:                String("0.0.0.0"),
				HTTPPort:            Int(80),
				HTTPSPort:           Int(443),
				EnableTLS:           Bool(true),
				TLSEmail:            String("someone@example.com"),
				EnableHTTPSRedirect: Bool(true),
				ConfigPath:          String(mockConfigPath),
				Storage:             ref.Ref(types.StorageJSON),
				CNames:              ref.Ref([]string{"example.com"}),
			},
			Services: map[string]*types.Service{
				"demo": {
					TargetHost: String("demo.com"),
					TargetPort: Int(80),
					EnableTLS:  Bool(false),
					Domains: []types.Domain{
						{
							DomainName: "sub.demo.com",
							Status: types.DomainStatus{
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
	j, err := NewJSONStorage(jsonOpts)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
		return
	}

	// try to init on disk
	err = j.init()
	if err != nil {
		t.Errorf("failed to init on disk: %v", err)
	}

	diskContent, err := j.read()
	if err != nil {
		t.Errorf("failed to read from disk: %v", err)
	}

	dcStr, err := diskContent.ToJSON()
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
	tempJ, err := NewJSONStorage(jsonOpts)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
	}
	if err = tempJ.init(); err != nil {
		t.Errorf("failed to init on disk: %v", err)
	}

	j, err := NewJSONStorage(
		&Opts{Config: &types.Config{Proxy: &types.Proxy{ConfigPath: String(mockConfigPath)}}},
	)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
	}
	defer func() {
		// clean up
		err = j.remove()
		if err != nil {
			t.Errorf("failed to clean up: %v", err)
		}
	}()

	err = j.Load()
	if err != nil {
		t.Errorf("failed to load from disk: %v", err)
	}

	memContentBytes, err := j.config.ToJSON()
	if err != nil {
		t.Errorf("failed to convert config to JSON bytes: %v", err)
	}

	if !reflect.DeepEqual(memContentBytes, mockConfigBytes) {
		t.Errorf(
			"expected config to be %v, got %v",
			string(mockConfigBytes),
			string(memContentBytes),
		)
	}
}

func Test_Persist(t *testing.T) {
	// init mock config on disk so we can load it again and compare
	tempJ, err := NewJSONStorage(jsonOpts)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
	}
	if err = tempJ.init(); err != nil {
		t.Errorf("failed to init on disk: %v", err)
	}

	// simulate updating data
	opts := jsonOpts
	if opts.Config.Admin == nil {
		opts.Config.Admin = &types.DefaultAdmin
	}
	opts.Config.Admin.Port = Int(21000)
	opts.Config.Proxy.EnableTLS = Bool(false)
	opts.Config.Proxy.TLSEmail = String("john@doe.com")
	if s, ok := opts.Config.Services["demo"]; ok {
		s.TargetPort = Int(5000)
	}

	j, err := NewJSONStorage(opts)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
	}

	defer func() {
		// clean up
		err = j.remove()
		if err != nil {
			t.Errorf("failed to clean up: %v", err)
		}
	}()

	err = j.Persist()
	if err != nil {
		t.Errorf("failed to persist: %v", err)
	}

	// read from disk and compare
	diskContent, err := j.read()
	if err != nil {
		t.Errorf("failed to read from disk: %v", err)
	}

	dcStr, err := diskContent.ToJSON()
	if err != nil {
		t.Errorf("failed to convert disk content to string: %v", err)
	}

	optsBytes, err := opts.Config.ToJSON()
	if err != nil {
		t.Errorf("failed to convert opts to JSON bytes: %v", err)
	}

	if !reflect.DeepEqual(dcStr, optsBytes) {
		t.Errorf("expected disk content to be %v, got %v", string(optsBytes), string(dcStr))
	}
}

func Test_Full(t *testing.T) {
	opts := jsonOpts

	j, err := NewJSONStorage(opts)
	if err != nil {
		t.Fatalf("failed to create JSON storage: %v", err)
	}

	defer func() {
		// clean up
		err := j.remove()
		if err != nil {
			t.Errorf("failed to clean up: %v", err)
		}
	}()

	// Load should and will always be called on startup
	err = j.Load()
	if err != nil {
		t.Errorf("failed to load from disk: %v", err)
	}

	if opts.Config.Admin == nil {
		opts.Config.Admin = &types.DefaultAdmin
	}

	// Concurrently update and persist data a couple of times to simulate a real world scenario (e.g. updating config via API)
	// Previous versions had a bug where the config file would completely be blank on persist; main reason for this test and rewrite
	for i := 0; i < 10; i++ {
		go func() {
			opts.Config.Admin.Port = Int(ref.Deref(opts.Config.Admin.Port, 0) + 100)
			opts.Config.Proxy.EnableTLS = Bool(
				!ref.Deref(ref.Deref(opts.Config.Proxy, types.DefaultProxy).EnableTLS, false),
			)
			opts.Config.Proxy.HTTPPort = Int(
				ref.Deref(ref.Deref(opts.Config.Proxy, types.DefaultProxy).HTTPPort, 0) + 1000,
			)
			err := j.Persist()
			if err != nil {
				t.Errorf("failed to persist: %v", err)
				return
			}
		}()
	}

	// Load again to make sure the data is still there
	err = j.Load()
	if err != nil {
		t.Errorf("failed to load from disk: %v", err)
	}

	// Compare
	memContentBytes, err := j.config.ToJSON()
	if err != nil {
		t.Errorf("failed to convert config to JSON bytes: %v", err)
	}

	optsBytes, err := opts.Config.ToJSON()
	if err != nil {
		t.Errorf("failed to convert opts to JSON bytes: %v", err)
	}

	if !reflect.DeepEqual(memContentBytes, optsBytes) {
		t.Errorf("expected config to be %v, got %v", string(optsBytes), string(memContentBytes))
	}
}
