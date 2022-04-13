package controller

import (
	"net/url"
	"strings"
	"xserver/entity/cache"
	"xserver/util"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(uri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := cache.SysTation(c.Query("stationGuid"))
		// TODO 不同设备向不同工作站请求
		pos := strings.Index(c.Request.URL.Path, uri)
		api := &url.URL{
			Scheme: s.Scheme,
			Host:   s.Host,
		}
		util.SingleHostProxy(api, c.Request.URL.Path[pos:], c)
		c.Abort()
	}
}
