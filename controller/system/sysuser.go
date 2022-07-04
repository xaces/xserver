package system

import (
	"errors"
	"xserver/entity/cache"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

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
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	if tok.UserName != model.SysUserName {
		p.createdBy = tok.UserName // 非管理员用户只能查看自己创建的用户
	}
	var data []model.SysUser
	total, _ := p.User().Model(&model.SysUser{}).Find(&data)
	ctx.JSONWrite(gin.H{"data": data, "total": total}, c)
}

// GetHandler 查询详细
func (o *User) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysUser{}, c)
}

// GetRolesHandler
func (o *User) GetRolesHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var roles []model.SysRole
	if _, err := orm.DbFindBy(&roles, "created_by = ?", tok.UserName); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONWriteData(roles, c)
}

// AddHandler 新增用户
func (o *User) AddHandler(c *gin.Context) {
	var p model.SysUser
	//获取参数
	if err := c.ShouldBind(&p.SysUserOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	p.OrganizeGUID = tok.OrganizeGUID
	p.OrganizeName = tok.OrganizeName
	p.CreatedBy = tok.UserName
	p.Type = model.SysUserTypeComm
	if err := service.SysUserCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// EnableHandler 改变状态
func (o *User) EnableHandler(c *gin.Context) {
	var p model.SysUser
	if err := c.ShouldBind(&p.SysUserOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(&p, p.ID, "enable", p.Enable); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// ExportHandler 导出
func (o *User) ExportHandler(c *gin.Context) {
	// var users []model.SysUser
	// totalCount, err := orm.DbFindPage(param.PageNum, param.PageSize, param.Where(sysUser), &model.SysUser{}, &users)
	ctx.JSONError(c)
}

type updatePwd struct {
	OldPassword string `form:"oldPassword" binding:"required,min=0,max=30"`
	NewPassword string `form:"newPassword" binding:"required,min=0,max=30"`
}

// ProfileHandler profile
func (o *User) ProfileHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var data model.SysUser
	if err := orm.DbFirstById(&data, tok.ID); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	u := cache.NewDevUser(data.ID, data.DeviceIds)
	u.Val.Clear()
	ctx.JSONWriteData(data, c)
}

// UpdatePwdHandler 重置密码
func (o *User) UpdatePwdHandler(c *gin.Context) {
	var p updatePwd
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	var data model.SysUser
	if err := orm.DbFirstById(&data, tok.ID); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	oldPassword := service.SysUserPassword(&data, p.OldPassword)
	if oldPassword != data.Password {
		ctx.JSONWriteError(errors.New("old password error"), c)
		return
	}
	newPassword := service.SysUserPassword(&data, p.NewPassword)
	if err := orm.DbUpdateColById(&data, data.ID, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
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
	if err := orm.DbUpdateColById(&model.SysUser{}, tok.ID, "avatar", avatarURL); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *User) UpdateHandler(c *gin.Context) {
	var data model.SysUser
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if data.UserName == model.SysUserName {
		data.RoleID = model.SysUserRoleId
	}
	if err := orm.DbUpdateModel(data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// ResetPwdHandler 修改密码
func (o *User) ResetPwdHandler(c *gin.Context) {
	id := ctx.ParamUInt(c, "id")
	var data model.SysUser
	if err := orm.DbFirstById(&data, id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	newPassword := service.SysUserPassword(&data, defaultpwd)
	if err := orm.DbUpdateColById(&data, id, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *User) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysUser{}, c)
}

type authDevice struct {
	UserId    uint   `form:"userId"`
	DeviceIds string `json:"deviceIds"`
}

// AuthDevicesHandler 授权
func (o *User) AuthDevicesHandler(c *gin.Context) {
	var p authDevice
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	v := cache.UserDevs(p.UserId)
	if v == nil {
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
	v := cache.UserDevs(p.UserId)
	if v == nil {
		ctx.JSONWriteError(nil, c)
		return
	}
	v.Dels(p.DeviceIds)
	orm.DbUpdateColById(&model.SysUser{}, p.UserId, "deviceIds", v.DeviceIds)
}

// DeviceTreeHandler 设备树
func (o *User) DeviceTreeHandler(c *gin.Context) {
	t := middleware.GetUserToken(c)
	devs := service.SysUserDevice(t)
	// 过滤用户数据
	var data []service.Vehicle
	u := cache.UserDevs(t.ID)
	if u != nil && devs != nil {
		for _, v := range devs {
			if u.DeviceIds == "*" {
				u.Update(v.Id)
			} else if !u.Include(v.Id) {
				continue
			}
			data = append(data, v)
		}
	}
	tree := service.OprOrganizeTree(t.OrganizeGUID, data)
	ctx.JSONWriteData(tree, c)
}

type devicelist struct {
	UserId uint64 `json:"userId"`
	Permis bool   `json:"permis"`
}

// DevicesHandler 用户
func (o *User) DevicesHandler(c *gin.Context) {
	var p devicelist
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	t := middleware.GetUserToken(c)
	devs := service.SysUserDevice(t)
	// 过滤用户数据
	u := cache.UserDevs(t.ID)
	if u == nil || devs == nil {
		ctx.JSONOk(c)
		return
	}
	if u.DeviceIds == "*" {
		tree := service.OprOrganizeTree(t.OrganizeGUID, devs)
		ctx.JSONWriteData(tree, c)
		return
	}
	var data []service.Vehicle
	for _, v := range devs {
		if p.Permis && !u.Include(v.Id) {
			continue
		} else if !p.Permis && u.Include(v.Id) {
			continue
		}
		data = append(data, v)
	}
	ctx.JSONWriteData(data, c)
}

func (o User) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.PageHandler)
	r.GET("/:id", o.GetHandler)
	r.GET("/getRoles", o.GetRolesHandler)
	r.POST("", o.AddHandler)
	r.GET("/export", o.ExportHandler)
	r.PUT("", o.UpdateHandler)
	r.PUT("/resetPwd/:id", o.ResetPwdHandler)
	r.DELETE("/:id", o.DeleteHandler)
	r.PUT("/enable", o.EnableHandler)
	r.GET("/profile", o.ProfileHandler)
	r.PUT("/profile", o.UpdateHandler)
	r.PUT("/profile/updatePwd", o.UpdatePwdHandler)
	r.POST("/profile/avatar", o.ProfileAuatarHandler)
	r.PUT("/authDevices", o.AuthDevicesHandler)
	r.PUT("/cancelAuthDevices", o.CancelAuthDevicesHandler)
	r.GET("/deviceTree", o.DeviceTreeHandler)
	r.GET("/devices", o.DevicesHandler)
}
