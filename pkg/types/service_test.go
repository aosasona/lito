package types

import (
	"testing"

	"go.trulyao.dev/lito/ext/option"
)

func Test_GetHostName(t *testing.T) {
	tests := []struct {
		name    string
		service *Service
		want    string
	}{
		{
			name:    "empty service",
			service: &Service{},
			want:    "",
		},
		{
			name: "service with scheme",
			service: &Service{
				TargetHost: option.StringValue("http://example.com"),
			},
			want: "http://example.com",
		},
		{
			name: "service with scheme and port",
			service: &Service{
				TargetHost: option.StringValue("https://example.com"),
				TargetPort: option.IntValue(8080),
			},
			want: "https://example.com:8080",
		},
		{
			name: "service with scheme, port, and path",
			service: &Service{
				TargetHost: option.StringValue("https://example.com"),
				TargetPort: option.IntValue(8080),
				TargetPath: option.StringValue("/some/path"),
			},
			want: "https://example.com:8080/some/path",
		},
		{
			name: "service with scheme, port, and path with leading slash and trailing slash in host",
			service: &Service{
				TargetHost: option.StringValue("https://example.com/"),
				TargetPort: option.IntValue(8080),
				TargetPath: option.StringValue("/some/path"),
			},
			want: "https://example.com:8080/some/path",
		},
	}

	for _, tt := range tests {
		if got := tt.service.GetTargetHost(); got != tt.want {
			t.Errorf("%s: got %s, want %s", tt.name, got, tt.want)
		}
	}
}
