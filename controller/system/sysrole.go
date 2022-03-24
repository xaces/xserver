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
	var param service.RolePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	where := param.Where()
	if tok.RoleId != model.SysUserRoleId {
		where.Append("created_by like ?", tok.UserName) // 非管理员只能获取当前用户创建的的角色
	}
	var data []model.SysRole
	total, _ := orm.DbPage(&model.SysRole{}, where).Find(param.PageNum, param.PageSize, &data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询
func (o *Role) GetHandler(c *gin.Context) {
	var data model.SysRole
	service.QueryById(&data, c)
}

// GetRolePowerHandler 查询
func (o *Role) GetRolePowerHandler(c *gin.Context) {
	getId, err := ctx.QueryUInt64(c, "roleId")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data model.SysRole
	if err := orm.DbFirstById(&data, getId); err != nil {
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
	ctx.JSONOk().WriteData(gin.H{"menuIds": data.MenuIds, "menus": menus}, c)
}

// AddHandler 新增
func (o *Role) AddHandler(c *gin.Context) {
	var data model.SysRole
	//获取参数
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
func (o *Role) UpdateHandler(c *gin.Context) {
	var data model.SysRole
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	// 更新数据
	if err := orm.DbUpdateById(data, data.Id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// EnableHandler 改变状态
func (o *Role) EnableHandler(c *gin.Context) {
	var data model.SysRole
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(model.SysRole{}, data.Id, "enable", data.Enable); err != nil {
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
	sysRole := Role{}
	r.GET("/role/list", sysRole.PageHandler)
	r.GET("/role/:id", sysRole.GetHandler)
	r.POST("/role", sysRole.AddHandler)
	r.GET("/role/getRolePower", sysRole.GetRolePowerHandler)
	r.PUT("/role/enable", sysRole.EnableHandler)
	r.PUT("/role", sysRole.UpdateHandler)
	r.DELETE("/role/:id", sysRole.DeleteHandler)
}