package controller

import (
	"net/url"
	"strings"
	"xserver/middleware"
	"xserver/util"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(uri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		pos := strings.Index(c.Request.URL.Path, uri)
		t := middleware.GetUserToken(c)
		api := &url.URL{
			Scheme: t.Scheme,
			Host:   t.Host,
		}
		if strings.Contains(c.Request.URL.Path, "") {
			api.RawQuery = "organizeGuid=" + t.OrganizeGuid
		}
		util.SingleHostProxy(api, c.Request.URL.Path[pos:], c)
		c.Abort()
	}
}
