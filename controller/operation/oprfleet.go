package operation

import (
	"errors"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Fleet
type Fleet struct {
}

// ListHandler 列表
func (o *Fleet) ListHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var data []model.OprOrganization
	toatl, _ := orm.DbFindBy(&data, "guid = ?", tok.OrganizeGuid)
	ctx.JSONOk().Write(gin.H{"total": toatl, "data": data}, c)
}

// LisTreeHandler 列表
func (o *Fleet) LisTreeHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	data := service.OprOrganizeTree(tok.OrganizeGuid, nil)
	if data == nil {
		ctx.JSONWriteError(errors.New("no data"), c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// GetHandler 详细
func (o *Fleet) GetHandler(c *gin.Context) {
	var data model.OprOrganization
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Fleet) AddHandler(c *gin.Context) {
	var data model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&data.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	// 获取组织信息, 从数据库
	data.Guid = middleware.GetUserToken(c).OrganizeGuid
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Fleet) UpdateHandler(c *gin.Context) {
	var data model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&data.OprOrganizationOpt); err != nil {
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
func (o *Fleet) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.OprOrganization{}, c)
}

// DevicesHandler 列表
func (o *Fleet) DevicesHandler(c *gin.Context) {
}

func FleetRouters(r *gin.RouterGroup) {
	o := Fleet{}
	r.GET("/fleet/list", o.ListHandler)
	r.GET("/fleet/listree", o.LisTreeHandler)
	r.GET("/fleet/:id", o.GetHandler)
	r.POST("/fleet", o.AddHandler)
	r.PUT("/fleet", o.UpdateHandler)
	r.DELETE("/fleet/:id", o.DeleteHandler)
	r.GET("/fleet/devices", o.DevicesHandler)
}
