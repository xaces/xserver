package util

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SingleHostProxy(api *url.URL, path string, c *gin.Context) {
	c.Request.URL.Path = path
	httputil.NewSingleHostReverseProxy(api).ServeHTTP(c.Writer, c.Request)
}
