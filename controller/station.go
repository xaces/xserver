package controller

import (
	"errors"
	"net/url"
	"strings"
	"xserver/entity/cache"
	"xserver/model"
	"xserver/util"

	"github.com/gin-gonic/gin"
	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"
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
	guid := c.Param("guid")
	if cache.SysTation(guid) == nil {
		ctx.JSONWriteError(errors.New("invalid station"), c)
		return
	}
	var vehis []model.OprVehicle
	orm.DbFindBy(&vehis, "station_guid = ?", guid)
	ctx.JSONWriteData(vehis, c)
}
