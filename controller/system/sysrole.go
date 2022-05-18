package system

import (
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/orm"
)

// Role
type Role struct {
}

// PageHandler 列表
func (o *Role) PageHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	if tok.RoleID != model.SysUserRoleId {
		p.createdBy = tok.UserName
	}
	var data []model.SysRole
	total, _ := orm.DbByWhere(&model.SysRole{}, p.Role()).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询
func (o *Role) GetHandler(c *gin.Context) {
	service.QueryById(&model.SysRole{}, c)
}

// GetRolePowerHandler 查询
func (o *Role) GetRolePowerHandler(c *gin.Context) {
	getId, err := ctx.QueryUInt(c, "roleId")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var p model.SysRole
	if err := orm.DbFirstById(&p, getId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	locRoleId := middleware.GetUserToken(c).RoleID
	var menus []model.SysMenu
	if locRoleId != model.SysUserRoleId {
		var locRole model.SysRole
		if err := orm.DbFirstById(&locRole, locRoleId); err != nil {
			ctx.JSONWriteError(err, c)
			return
		}
		orm.DbFindBy(&menus, "id in (?)", locRole.MenuIds)
	} else {
		orm.DbFind(&menus)
	}
	// roles 查询用户权限 menus当前登录用户全面
	ctx.JSONWriteData(gin.H{"menuIds": p.MenuIds, "menus": menus}, c)
}

// AddHandler 新增
func (o *Role) AddHandler(c *gin.Context) {
	var p model.SysRole
	//获取参数
	if err := c.ShouldBind(&p.SysRoleOpt); err != nil {
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
func (o *Role) UpdateHandler(c *gin.Context) {
	var p model.SysRole
	if err := c.ShouldBind(&p.SysRoleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	// 更新数据
	if err := orm.DbUpdateById(p, p.ID); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// EnableHandler 改变状态
func (o *Role) EnableHandler(c *gin.Context) {
	var p model.SysRole
	//获取参数
	if err := c.ShouldBind(&p.SysRoleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(model.SysRole{}, p.ID, "enable", p.Enable); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *Role) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysRole{}, c)
}

func RoleRouters(r *gin.RouterGroup) {
	o := Role{}
	r.GET("/list", o.PageHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.GET("/getRolePower", o.GetRolePowerHandler)
	r.PUT("/enable", o.EnableHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
