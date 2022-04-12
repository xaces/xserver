package controller

import (
	"net/url"
	"strings"
	"xserver/util"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(uri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 不同设备向不同工作站请求
		pos := strings.Index(c.Request.URL.Path, uri)
		api := &url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:12100",
		}
		util.SingleHostProxy(api, c.Request.URL.Path[pos:], c)
		c.Abort()
	}
}
