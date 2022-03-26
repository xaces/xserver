package system

import (
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Station
type Station struct {
}

// ListHandler 列表
func (o *Station) ListHandler(c *gin.Context) {
	var p orm.DbPage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysStation
	totalCount, _ := orm.DbByWhere(&model.SysStation{}, p.DbWhere()).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": totalCount, "data": data}, c)
}

// GetHandler 详细
func (o *Station) GetHandler(c *gin.Context) {
	var data model.SysStation
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Station) AddHandler(c *gin.Context) {
	var data model.SysStation
	//获取参数
	if err := c.ShouldBind(&data.SysStationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	data.Guid = util.NUID()
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Station) UpdateHandler(c *gin.Context) {
	var data model.SysStation
	//获取参数
	if err := c.ShouldBind(&data.SysStationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Station) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysStation{}, c)
}

func StationRouters(r *gin.RouterGroup) {
	o := Station{}
	r.GET("/station/list", o.ListHandler)
	r.GET("/station/:id", o.GetHandler)
	r.POST("/station", o.AddHandler)
	r.PUT("/station", o.UpdateHandler)
	r.DELETE("/station/:id", o.DeleteHandler)
}
