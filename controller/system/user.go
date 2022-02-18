package system

import (
	"errors"
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"xserver/middleware"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

const (
	defaultpwd = "123456"
)

// User 系统管理用户
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
	var users []model.SysUser
	totalCount, _ := orm.DbPage(&model.SysUser{}, where).Find(param.Page, param.Limit, &users)
	ctx.JSONOk().Write(gin.H{"data": users, "count": totalCount}, c)
}

// GetHandler 查询详细
func (o *User) GetHandler(c *gin.Context) {
	userId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var user model.SysUser
	if err := orm.DbFirstById(&user, userId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(&user, c)
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
	var user model.SysUser
	//获取参数
	if err := c.ShouldBind(&user); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := service.CheckAddUser(&user); err == nil {
		ctx.JSONWriteError(errors.New("user already exists"), c)
		return
	}
	user.CreatedBy = middleware.GetUserToken(c).UserName
	user.Password = service.NewSysPassword(&user, defaultpwd)
	if err := orm.DbCreate(&user); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// EnableHandler 改变状态
func (o *User) EnableHandler(c *gin.Context) {
	var user model.SysUser
	if err := c.ShouldBind(&user); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateColById(&user, user.Id, "enable", user.Enable); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// AddPageHandler 新增界面
func (o *User) AddPageHandler(c *gin.Context) {
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
	var user model.SysUser
	if err := orm.DbFirstById(&user, tok.UserId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(user, c)
}

// UpdatePwdHandler 重置密码
func (o *User) UpdatePwdHandler(c *gin.Context) {
	var param updatePwd
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	var user model.SysUser
	if err := orm.DbFirstById(&user, tok.UserId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	oldPassword := service.NewSysPassword(&user, param.OldPassword)
	if oldPassword != user.Password {
		ctx.JSONWriteError(errors.New("old password error"), c)
		return
	}
	newPassword := service.NewSysPassword(&user, param.NewPassword)
	if err := orm.DbUpdateColById(&user, user.Id, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// ProfileAuatarHandler 上传头像
func (o *User) ProfileAuatarHandler(c *gin.Context) {
	// claims, err := middleware.GetUserOfToken(c)
	// if err != nil {
	// 	ctx.JSONWriteError(err, c)
	// 	return
	// }
	// path := "/public/upload/"
	// saveDir := configs.Default.Local.ServeRootPath + path
	// fileHead, err := c.FormFile("avatarfile")
	// if err != nil {
	// 	ctx.JSONWriteError(err, c)
	// 	return
	// }
	// curdate := time.Now().UnixNano()
	// filename := claims.Username + strconv.FormatInt(curdate, 10) + ".png"
	// dts := saveDir + filename
	// if err := c.SaveUploadedFile(fileHead, dts); err != nil {
	// 	ctx.JSONWriteError(err, c)
	// 	return
	// }
	// avatarURL := configs.Default.Local.ServeRoot + path + filename
	// if err := orm.DbUpdateColById(&model.SysUser{}, claims.UserID, "avatar", avatarURL); err != nil {
	// 	ctx.JSONWriteError(err, c)
	// 	return
	// }
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *User) UpdateHandler(c *gin.Context) {
	var user model.SysUser
	if err := c.ShouldBind(&user); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if user.UserName == model.SysUserName {
		user.RoleId = model.SysUserRoleId
	}
	if err := orm.DbUpdateModel(user); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// ResetPwdHandler 修改密码
func (o *User) ResetPwdHandler(c *gin.Context) {
	userId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var user model.SysUser
	if err := orm.DbFirstById(&user, userId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	newPassword := service.NewSysPassword(&user, defaultpwd)
	if err := orm.DbUpdateColById(&user, userId, "password", newPassword); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *User) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysUser{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
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
	r.PUT("/user/profile/updatePwd", sysUser.UpdatePwdHandler)
	r.POST("/user/profile/avatar", sysUser.ProfileAuatarHandler)
}
