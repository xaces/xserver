package system

import (
	"errors"
	"xserver/entity/mnger"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

const (
	defaultpwd = "123456"
)

// User
type User struct {
}

// PageHandler 列表
func (o *User) PageHandler(c *gin.Context) {
	var param service.UserPage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	where := param.Where()
	if tok.UserName != model.SysUserName {
		where.String("created_by like ?", tok.UserName) // 非管理员用户只能查看自己创建的用户
	}
	var data []model.SysUser
	total, _ := orm.DbByWhere(&model.SysUser{}, where).Find(&data)
	ctx.JSONOk().Write(gin.H{"data": data, "total": total}, c)
}

// GetHandler 查询详细
func (o *User) GetHandler(c *gin.Context) {
	var data model.SysUser
	service.QueryById(&data, c)
}

// GetRolesHandler
func (o *User) GetRolesHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var roles []model.SysRole
	if _, err := orm.DbFindBy(&roles, "created_by like ?", tok.UserName); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(roles, c)
}

// AddHandler 新增用户
func (o *User) AddHandler(c *gin.Context) {
	var data model.SysUser
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	data.OrganizeGuid = tok.OrganizeGuid
	data.OrganizeName = tok.OrganizeName
	data.CreatedBy = tok.UserName
	data.UserType = model.SysUserTypeComm
	if err := service.SysUserCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// EnableHandler 改变状态
func (o *User) EnableHandler(c *gin.Context) {
	var data model.SysUser
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(&data, data.Id, "enable", data.Enable); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// ExportHandler 导出
func (o *User) ExportHandler(c *gin.Context) {
	// var users []model.SysUser
	// totalCount, err := orm.DbFindPage(param.PageNum, param.PageSize, param.Where(sysUser), &model.SysUser{}, &users)
	ctx.JSON(ctx.StatusError).WriteTo(c)
}

type updatePwd struct {
	OldPassword string `form:"oldPassword" binding:"required,min=0,max=30"`
	NewPassword string `form:"newPassword" binding:"required,min=0,max=30"`
}

// ProfileHandler profile
func (o *User) ProfileHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var data model.SysUser
	if err := orm.DbFirstById(&data, tok.Id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// UpdatePwdHandler 重置密码
func (o *User) UpdatePwdHandler(c *gin.Context) {
	var param updatePwd
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	var data model.SysUser
	if err := orm.DbFirstById(&data, tok.Id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	oldPassword := service.SysUserPassword(&data, param.OldPassword)
	if oldPassword != data.Password {
		ctx.JSONWriteError(errors.New("old password error"), c)
		return
	}
	newPassword := service.SysUserPassword(&data, param.NewPassword)
	if err := orm.DbUpdateColById(&data, data.Id, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// ProfileAuatarHandler 上传头像
func (o *User) ProfileAuatarHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	path := "/public/upload/"
	fileHead, err := c.FormFile("avatarfile")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	filename := tok.UserName + ".png"
	dts := path + filename
	if err := c.SaveUploadedFile(fileHead, dts); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	avatarURL := dts
	if err := orm.DbUpdateColById(&model.SysUser{}, tok.Id, "avatar", avatarURL); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *User) UpdateHandler(c *gin.Context) {
	var data model.SysUser
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if data.UserName == model.SysUserName {
		data.RoleId = model.SysUserRoleId
	}
	if err := orm.DbUpdateModel(data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// ResetPwdHandler 修改密码
func (o *User) ResetPwdHandler(c *gin.Context) {
	getId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data model.SysUser
	if err := orm.DbFirstById(&data, getId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	newPassword := service.SysUserPassword(&data, defaultpwd)
	if err := orm.DbUpdateColById(&data, getId, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *User) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysUser{}, c)
}

type authDevice struct {
	UserId    uint64 `form:"userId"`
	DeviceIds string `json:"deviceIds"`
}

// AuthDevicesHandler 授权
func (o *User) AuthDevicesHandler(c *gin.Context) {
	var p authDevice
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	v, ok := mnger.UserDevs[p.UserId]
	if !ok {
		ctx.JSONWriteError(nil, c)
		return
	}
	v.Set(p.DeviceIds)
	orm.DbUpdateColById(&model.SysUser{}, p.UserId, "deviceIds", v.DeviceIds)
}

// CancelAuthDevicesHandler 取消授权
func (o *User) CancelAuthDevicesHandler(c *gin.Context) {
	var p authDevice
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	v, ok := mnger.UserDevs[p.UserId]
	if !ok {
		ctx.JSONWriteError(nil, c)
		return
	}
	v.Dels(p.DeviceIds)
	orm.DbUpdateColById(&model.SysUser{}, p.UserId, "deviceIds", v.DeviceIds)
}

func UserRouters(r *gin.RouterGroup) {
	sysUser := User{}
	r.GET("/user/list", sysUser.PageHandler)
	r.GET("/user/:id", sysUser.GetHandler)
	r.GET("/user/getRoles", sysUser.GetRolesHandler)
	r.POST("/user", sysUser.AddHandler)
	r.GET("/user/export", sysUser.ExportHandler)
	r.PUT("/user", sysUser.UpdateHandler)
	r.PUT("/user/resetPwd/:id", sysUser.ResetPwdHandler)
	r.DELETE("/user/:id", sysUser.DeleteHandler)
	r.PUT("/user/enable", sysUser.EnableHandler)
	r.GET("/user/profile", sysUser.ProfileHandler)
	r.PUT("/user/profile", sysUser.UpdateHandler)
	r.PUT("/user/authDevices", sysUser.AuthDevicesHandler)
	r.PUT("/user/cancelAuthDevices", sysUser.CancelAuthDevicesHandler)
	r.PUT("/user/profile/updatePwd", sysUser.UpdatePwdHandler)
	r.POST("/user/profile/avatar", sysUser.ProfileAuatarHandler)
}
