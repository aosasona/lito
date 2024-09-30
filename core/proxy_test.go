package core

import (
	"net/http"
	"testing"

	"go.trulyao.dev/lito/pkg/ref"
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
		if service != nil && ref.Deref(service.TargetHost, "") != test.ExpectedTargetHost {
			t.Errorf("findServiceByName(%s) = %v, want %v", test.ServiceName, service.TargetHost, test.ExpectedTargetHost)
		}

		if service == nil && test.ExpectedTargetHost != "" {
			if ref.Deref(service.TargetHost, "") != test.ExpectedTargetHost {
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

		if service != nil && ref.Deref(service.TargetHost, "") != test.ExpectedTargetHost {
			t.Errorf("findServiceByDomainName(%s) = %v, want %v", test.DomainName, service.TargetHost, test.ExpectedTargetHost)
		}
	}
}

func Test_ProxyDirector(t *testing.T) {
	request, err := http.NewRequest("GET", "https://demo.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	config := mockConfig()
	New(&Opts{Config: &config}).proxyDirector(request)

	if request.URL.Scheme != "https" {
		t.Errorf("proxyDirector() = %v, want %v", request.URL.Scheme, "https")
	}

	if request.URL.Host != "something.dp.com" {
		t.Errorf("proxyDirector() = %v, want %v", request.URL.Host, "something.dp.com")
	}

	if request.URL.Path != "/demo/" {
		t.Errorf("proxyDirector() = %v, want %v", request.URL.Path, "/demo/")
	}

	if request.URL.RawQuery != "" {
		t.Errorf("proxyDirector() = %v, want %v", request.URL.RawQuery, "")
	}

	if request.Header.Get("X-Service-Name") != "demo" {
		t.Errorf("proxyDirector() = %v, want %v", request.Header.Get("X-Service-Name"), "demo")
	}

	// Test full URL
	if request.URL.String() != "https://something.dp.com/demo/" {
		t.Errorf("proxyDirector() = %v, want %v", request.URL.String(), "https://something.dp.com/demo")
	}
}
