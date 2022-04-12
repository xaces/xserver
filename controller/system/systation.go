package system

import (
	"xserver/middleware"
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
	tok := middleware.GetUserToken(c)
	var data []model.SysStation
	toatl, _ := orm.DbFindBy(&data, "organize_guid = ?", tok.OrganizeGuid)
	ctx.JSONOk().Write(gin.H{"total": toatl, "data": data}, c)
}

// GetHandler 详细
func (o *Station) GetHandler(c *gin.Context) {
	var data model.SysStation
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Station) AddHandler(c *gin.Context) {
	var p model.SysStation
	if err := c.ShouldBind(&p.SysStationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	t := middleware.GetUserToken(c)
	p.Guid = util.NUID()
	p.OrganizeGuid = t.OrganizeGuid
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Station) UpdateHandler(c *gin.Context) {
	var p model.SysStation
	if err := c.ShouldBind(&p.SysStationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
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
