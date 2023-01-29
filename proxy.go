package ginkit

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RemoveHeaders(headers ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		for _, header := range headers {
			c.Request.Header.Del(header)
		}
	}
}

func Proxy(target string) func(*gin.Context) {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		// TODO: Replace the transport with something configurable.  Ignore backend SSL, TTL, etc
		// proxy.Transport = ...
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
