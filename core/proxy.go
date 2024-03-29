package core

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"go.trulyao.dev/lito/ext/option"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

var proxyHttpServer *http.Server

// TODO: drop certmagic in here to use tls if turned on
func (c *Core) startProxy() error {
	if c.config.Proxy.IsNone() {
		return errors.New("no proxy config present")
	}

	reverseProxy := &httputil.ReverseProxy{
		Director: c.proxyDirector,
	}

	httpPort := c.config.Proxy.Unwrap(&types.DefaultProxy).HTTPPort.Unwrap(80)
	proxyHttpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: reverseProxy,
	}

	return proxyHttpServer.ListenAndServe()
}

func (c *Core) stopProxy() error {
	return proxyHttpServer.Shutdown(nil)
}

func (c *Core) proxyDirector(req *http.Request) {
	req.Header.Del("X-Forwarded-For")

	serviceName, optService, ok := c.findServiceByDomainName(req.Host)
	if !ok {
		if c.debug {
			c.logHandler.Debug("no service found for domain", logger.Field("domain", req.Host))
		}
		return
	}

	targetService := optService.Unwrap(&types.DefaultService)

	targetURL, err := url.Parse(targetService.GetTargetHost())
	if err != nil {
		return
	}

	// Modify the request to point to the target URL
	req.URL.Scheme = targetURL.Scheme
	req.URL.Host = targetURL.Host
	req.URL.Path, req.URL.RawPath = joinURLPath(targetURL, req.URL)

	// Copy the URL query to the "new" request
	if targetURL.RawQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetURL.RawQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetURL.RawQuery + "&" + req.URL.RawQuery
	}

	// Strip all headers that are specified in the service
	for _, header := range targetService.StripHeaders.Unwrap([]string{}) {
		req.Header.Del(header)
	}

	// Set the custom X-Service-Name header to the name of the service
	req.Header.Set("X-Service-Name", serviceName)

	// Log the request if debug mode is enabled
	if c.debug {
		c.logHandler.Debug("proxying request",
			logger.Field("from", req.Host),
			logger.Field("to", req.URL.Host),
			logger.Field("path", req.URL.Path),
			logger.Field("query", req.URL.RawQuery),
			logger.Field("headers", req.Header),
		)
	}
}

// findServiceByName finds a service by its name
func (c *Core) findServiceByName(name string) (option.Option[*types.Service], bool) {
	service, ok := c.config.Services[name]
	if !ok {
		return option.None[*types.Service](), false
	}
	return option.Some(service), true
}

// findServiceByHostname finds a service by the domain the request is issued to
func (c *Core) findServiceByDomainName(domainName string) (string, option.Option[*types.Service], bool) {
	for name, service := range c.config.Services {
		for _, serviceHost := range service.Domains {
			if serviceHost.DomainName == domainName {
				return name, option.Some(service), true
			}
		}
	}
	return "", option.None[*types.Service](), false
}

// Source: net/http/httputil/reverseproxy.go
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// Source: net/http/httputil/reverseproxy.go
func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}
