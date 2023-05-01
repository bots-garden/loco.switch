package handlers

import (
	"loco-switch/data"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

/*
TODO:

  - manage errors with the appropriate code (not panic)
  - add an error counter:
    if for a same function/endpoint, there are to many errors,
    then unregister the function/endpoint
*/
func Proxy(c *gin.Context) {

	key := c.Param("function_name") + ":default"

	endpointUrl, err := data.GetEndpointUrl(key)

	// Call of function that is not registered
	if err != nil { 
		log.Println("🔴 handlers.Proxy [data.GetEndpointUrl]", err)
		//return
	}

	remote, err := url.Parse(endpointUrl)
	if err != nil {
		log.Println("🔴 handlers.Proxy [url.Parse]", err)
		//return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {	
		// Logging requests body here	
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

/*
TODO:

  - manage errors with the appropriate code (not panic)
  - add an error counter:
    if for a same function/endpoint, there are to many errors,
    then unregister the function/endpoint
*/
func ProxyRevision(c *gin.Context) {
	key := c.Param("function_name") + ":" + c.Param("function_revision")

	endpointUrl, err := data.GetEndpointUrl(key)

	// Call of function that is not registered
	if err != nil {
		log.Println("🔴 handlers.ProxyRevision [data.GetEndpointUrl]", err)
		//return
	}

	remote, err := url.Parse(endpointUrl)
	if err != nil {
		log.Println("🔴 handlers.ProxyRevision [url.Parse]", err)
		//return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		// Logging requests body here		

		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
