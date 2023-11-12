package core

import (
	"testing"
)

func Test_FindServiceByName(t *testing.T) {
	tests := []struct {
		ServiceName        string
		ExpectedTargetHost string
		Want               bool
	}{
		{"demo", "https://something.dp.com", true},
		{"demo2", "https://anotherdemo.com", true},
		{"demo23", "", false},
	}

	config := mockConfig()
	core := New(&Opts{Config: &config})

	for _, test := range tests {
		service, got := core.findServiceByName(test.ServiceName)
		if got != test.Want {
			t.Errorf("findServiceByName(%s) = %v, want %v", test.ServiceName, got, test.Want)
		}

		if service != nil && service.TargetHost.Unwrap() != test.ExpectedTargetHost {
			t.Errorf("findServiceByName(%s) = %v, want %v", test.ServiceName, service.TargetHost, test.ExpectedTargetHost)
		}

		if service == nil && test.ExpectedTargetHost != "" {
			if service.TargetHost.Unwrap() != test.ExpectedTargetHost {
				t.Errorf("findServiceByName(%s) = %v, want %v", test.ServiceName, service.TargetHost, test.ExpectedTargetHost)
			}
		}
	}
}

func Test_FindServiceByHostName(t *testing.T) {
	tests := []struct {
		DomainName          string
		ExpectedServiceName string
		ExpectedTargetHost  string
		Want                bool
	}{
		{"demo.com", "demo", "https://something.dp.com", true},
		{"addftp.com", "demo2", "https://anotherdemo.com", true},
		{"seconddemo.ai", "demo2", "https://anotherdemo.com", true},
		{"something.com", "", "", false},
	}

	config := mockConfig()
	core := New(&Opts{Config: &config})

	for _, test := range tests {
		name, service, got := core.findServiceByDomainName(test.DomainName)
		if got != test.Want {
			t.Errorf("findServiceByDomainName(%s) = %v, want %v", test.DomainName, got, test.Want)
		}

		if name != test.ExpectedServiceName {
			t.Errorf("findServiceByDomainName(%s) = %v, want %v", test.DomainName, name, test.ExpectedServiceName)
		}

		if service != nil && service.TargetHost.Unwrap() != test.ExpectedTargetHost {
			t.Errorf("findServiceByDomainName(%s) = %v, want %v", test.DomainName, service.TargetHost, test.ExpectedTargetHost)
		}
	}
}
