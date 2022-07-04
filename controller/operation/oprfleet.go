package operation

import (
	"errors"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Fleet
type Fleet struct {
}

// ListHandler 列表
func (o *Fleet) ListHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var data []model.OprOrganization
	toatl, _ := orm.DbFindBy(&data, "guid = ?", tok.OrganizeGUID)
	ctx.JSONWrite(gin.H{"total": toatl, "data": data}, c)
}

// LisTreeHandler 列表
func (o *Fleet) LisTreeHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	data := service.OprOrganizeTree(tok.OrganizeGUID, nil)
	if data == nil {
		ctx.JSONWriteError(errors.New("no data"), c)
		return
	}
	ctx.JSONWriteData(data, c)
}

// GetHandler 详细
func (o *Fleet) GetHandler(c *gin.Context) {
	service.QueryByID(&model.OprOrganization{}, c)
}

// AddHandler 新增
func (o *Fleet) AddHandler(c *gin.Context) {
	var p model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&p.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	// 获取组织信息, 从数据库
	p.GUID = middleware.GetUserToken(c).OrganizeGUID
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *Fleet) UpdateHandler(c *gin.Context) {
	var p model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&p.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *Fleet) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.OprOrganization{}, c)
}

// DevicesHandler 列表
func (o *Fleet) DevicesHandler(c *gin.Context) {
}

func (o Fleet) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
	r.GET("/listree", o.LisTreeHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
	r.GET("/devices", o.DevicesHandler)
}
