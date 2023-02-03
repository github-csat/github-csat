package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func setUpDevProxy(handlers *Handlers, router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			proxy(handlers.ProxyFrontendURL)(c)
		}
		//default 404 page not found
	})
}

// proxy was stolen from...somewhere. stackoverflow maybe?
func proxy(remote *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
