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
	var data []model.SysTation
	toatl, _ := orm.DbFindBy(&data, "organize_guid = ?", tok.OrganizeGUID)
	ctx.JSONWrite(gin.H{"total": toatl, "data": data}, c)
}

// GetHandler 详细
func (o *Station) GetHandler(c *gin.Context) {
	service.QueryById(&model.SysTation{}, c)
}

// AddHandler 新增
func (o *Station) AddHandler(c *gin.Context) {
	var p model.SysTation
	if err := c.ShouldBind(&p.SysTationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	t := middleware.GetUserToken(c)
	p.GUID = util.NUID()
	p.OrganizeGUID = t.OrganizeGUID
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *Station) UpdateHandler(c *gin.Context) {
	var p model.SysTation
	if err := c.ShouldBind(&p.SysTationOpt); err != nil {
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
func (o *Station) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysTation{}, c)
}

func StationRouters(r *gin.RouterGroup) {
	o := Station{}
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
