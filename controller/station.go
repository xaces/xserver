package controller

import (
	"net/http/httputil"
	"net/url"
	"strings"
	"xserver/entity/mnger"
	"xserver/middleware"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(uri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "api/device/list") {
			mnger.DevUserList(c)
			return
		}
		pos := strings.Index(c.Request.URL.Path, uri)
		api := new(url.URL)
		api.Host = middleware.GetUserToken(c).Host
		api.Scheme = c.Request.URL.Scheme
		c.Request.URL.Path = c.Request.URL.Path[pos:]
		httputil.NewSingleHostReverseProxy(api).ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
