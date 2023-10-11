package core

import "testing"

func Test_FindServiceByName(t *testing.T) {
	tests := []struct {
		ServiceName        string
		ExpectedTargetHost string
		Want               bool
	}{
		{"demo", "something.dp.com", true},
		{"demo2", "anotherdemo.com", true},
		{"demo23", "", false},
	}

	config := mockConfig()
	core := New(&Opts{Config: &config})

	for _, test := range tests {
		service, got := core.findServiceByName(test.ServiceName)
		if got != test.Want {
			t.Errorf("lookupServiceByName(%s) = %v, want %v", test.ServiceName, got, test.Want)
		}

		if service != nil && service.TargetHost != test.ExpectedTargetHost {
			t.Errorf("lookupServiceByName(%s) = %v, want %v", test.ServiceName, service.TargetHost, test.ExpectedTargetHost)
		}

		if service == nil && test.ExpectedTargetHost != "" {
			if service.TargetHost != test.ExpectedTargetHost {
				t.Errorf("lookupServiceByName(%s) = %v, want %v", test.ServiceName, service.TargetHost, test.ExpectedTargetHost)
			}
		}
	}
}

func Test_FindServiceByProxiedDomain(t *testing.T) {
	tests := []struct {
		DomainName          string
		ExpectedServiceName string
		ExpectedTargetHost  string
		Want                bool
	}{
		{"demo.com", "demo", "something.dp.com", true},
		{"addftp.com", "demo2", "anotherdemo.com", true},
		{"seconddemo.ai", "demo2", "anotherdemo.com", true},
		{"something.com", "", "", false},
	}

	config := mockConfig()
	core := New(&Opts{Config: &config})

	for _, test := range tests {
		name, service, got := core.findServiceByProxiedDomain(test.DomainName)
		if got != test.Want {
			t.Errorf("lookupServiceByDomainName(%s) = %v, want %v", test.DomainName, got, test.Want)
		}

		if name != test.ExpectedServiceName {
			t.Errorf("lookupServiceByDomainName(%s) = %v, want %v", test.DomainName, name, test.ExpectedServiceName)
		}

		if service != nil && service.TargetHost != test.ExpectedTargetHost {
			t.Errorf("lookupServiceByDomainName(%s) = %v, want %v", test.DomainName, service.TargetHost, test.ExpectedTargetHost)
		}
	}
}
