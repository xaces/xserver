package controller

import (
	"net/http/httputil"
	"net/url"
	"strings"
	"xserver/entity/mnger"
	"xserver/middleware"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(api string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pos := strings.Index(c.Request.URL.Path, api)
		if pos < 0 {
			c.Next()
			return
		}
		org := mnger.Company(middleware.GetUserToken(c).OrganizeGuid)
		if org == nil {
			c.Next()
			return
		}
		proxy := new(url.URL)
		proxy.Host = org.SysStation.Host
		proxy.Scheme = org.SysStation.Scheme
		c.Request.URL.Path = c.Request.URL.Path[pos:]
		if c.Request.URL.Path == "/station/api/device/list" {
			proxy.RawQuery = "organizeGuid=" + org.OrganizeGuid
		}
		httputil.NewSingleHostReverseProxy(proxy).ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
