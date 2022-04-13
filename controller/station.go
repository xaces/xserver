package controller

import (
	"net/url"
	"strings"
	"xserver/entity/cache"
	"xserver/model"
	"xserver/util"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"
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

func DevicesHandler(c *gin.Context) {
	guid := ctx.ParamString(c, "guid")
	var vehis []model.OprVehicle
	if guid != "" {
		orm.DbFindBy(&vehis, "station_guid = ?", guid)
	}
	ctx.JSONOk().WriteData(vehis, c)
}
