package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"

	"github.com/gin-gonic/gin"
	"github.com/xaces/xutils/orm"
)

// Menu
type Menu struct {
}

// ListHandler 列表
func (o *Menu) ListHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysMenu
	total, _ := p.Where().Model(&model.SysMenu{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询详细
func (o *Menu) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysMenu{}, c)
}

// AddHandler 新增
func (o *Menu) AddHandler(c *gin.Context) {
	var p model.SysMenu
	if err := c.ShouldBind(&p.SysMenuOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *Menu) UpdateHandler(c *gin.Context) {
	var p model.SysMenu
	if err := c.ShouldBind(&p.SysMenuOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除oo
func (o *Menu) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysMenu{}, c)
}
func (o Menu) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
