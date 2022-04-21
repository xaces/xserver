package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/orm"
)

// Menu
type Menu struct {
}

// ListHandler 列表
func (o *Menu) ListHandler(c *gin.Context) {
	var param service.MenuPage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysMenu
	total, _ := orm.DbByWhere(&model.SysMenu{}, param.Where()).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询详细
func (o *Menu) GetHandler(c *gin.Context) {
	service.QueryById(&model.SysMenu{}, c)
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
	ctx.JSONOk().WriteTo(c)
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
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Menu) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysMenu{}, c)
}

func MenuRouters(r *gin.RouterGroup) {
	sysMenu := Menu{}
	r.GET("/menu/list", sysMenu.ListHandler)
	r.GET("/menu/:id", sysMenu.GetHandler)
	r.POST("/menu", sysMenu.AddHandler)
	r.PUT("/menu", sysMenu.UpdateHandler)
	r.DELETE("/menu/:id", sysMenu.DeleteHandler)
}
