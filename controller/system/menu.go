package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/orm"
)

// Menu 系统管理菜单
type Menu struct {
}

// ListHandler 列表
func (o *Menu) ListHandler(c *gin.Context) {
	var param service.MenuPage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var menus []model.SysMenu
	totalCount, _ := orm.DbPage(&model.SysMenu{}, param.Where()).Find(param.PageNum, param.PageSize, &menus)
	ctx.JSONOk().Write(gin.H{"count": totalCount, "data": menus}, c)
}

// GetHandler 查询菜单详细
func (o *Menu) GetHandler(c *gin.Context) {
	getId, _ := ctx.ParamInt(c, "id")
	var data model.SysMenu
	if err := orm.DbFirstById(&data, getId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// AddHandler 新增
func (o *Menu) AddHandler(c *gin.Context) {
	var data model.SysMenu
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Menu) UpdateHandler(c *gin.Context) {
	var data model.SysMenu
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除菜单
func (o *Menu) DeleteHandler(c *gin.Context) {
	id, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if _, err := orm.DbDeleteBy(model.SysMenu{}, "id = ?", id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func MenuRouters(r *gin.RouterGroup) {
	sysMenu := Menu{}
	r.GET("/menu/list", sysMenu.ListHandler)
	r.GET("/menu/:id", sysMenu.GetHandler)
	r.POST("/menu", sysMenu.AddHandler)
	r.PUT("/menu", sysMenu.UpdateHandler)
	r.DELETE("/menu/:id", sysMenu.DeleteHandler)
}
