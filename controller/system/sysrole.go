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
	var p service.RolePage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	where := p.Where()
	if tok.RoleId != model.SysUserRoleId {
		where.Append("created_by like ?", tok.UserName) // 非管理员只能获取当前用户创建的的角色
	}
	var data []model.SysRole
	total, _ := orm.DbByWhere(&model.SysRole{}, where).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
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
	locRoleId := middleware.GetUserToken(c).RoleId
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
	ctx.JSONOk().WriteData(gin.H{"menuIds": p.MenuIds, "menus": menus}, c)
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
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Role) UpdateHandler(c *gin.Context) {
	var p model.SysRole
	if err := c.ShouldBind(&p.SysRoleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	// 更新数据
	if err := orm.DbUpdateById(p, p.Id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// EnableHandler 改变状态
func (o *Role) EnableHandler(c *gin.Context) {
	var p model.SysRole
	//获取参数
	if err := c.ShouldBind(&p.SysRoleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(model.SysRole{}, p.Id, "enable", p.Enable); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
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
