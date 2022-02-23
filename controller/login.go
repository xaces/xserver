package controller

import (
	"errors"
	"strings"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

var (
	errLogin          = errors.New("Account or pssword error")
	errCaptachaVerify = errors.New("Captcha verify error")
	errAuth           = errors.New("Authentication error")
)

type loginReq struct {
	Username string `json:"username" binding:"required,max=30"`
	Password string `json:"password" binding:"required,max=128"`
	Captcha  string `json:"captcha" binding:"required,max=10"`
	UID      string `json:"uid" binding:"required,min=5,max=30"`
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	var lo *loginReq
	err := c.ShouldBindJSON(&lo)
	if err != nil || lo == nil {
		ctx.JSONWriteError(errLogin, c)
		return
	}
	if !store.Verify(lo.UID, lo.Captcha, true) {
		ctx.JSONWriteError(errCaptachaVerify, c)
		return
	}
	user := &model.SysUser{}
	user.UserName = lo.Username
	err = orm.DbFirstWhere(user, user)
	if err != nil {
		if lo.Username == model.SysUserName {
			user.Enable = 1
			user.RoleId = model.SysUserRoleId
			user.NickName = "administrator"
			user.Password = service.NewSysPassword(user, lo.Password)
			user.CreatedBy = "default"
			err = orm.DbCreate(user)
		}
	}
	if err != nil {
		ctx.JSONWriteError(errLogin, c)
		return
	}
	verifyPassword := service.NewSysPassword(user, lo.Password)
	if strings.Compare(user.Password, verifyPassword) != 0 {
		ctx.JSONWriteError(errLogin, c)
		return
	}
	tokenNext(c, user)
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, u *model.SysUser) {
	// proxy := middleware.ProxyStation{Guid: user.ProxyGuid}
	// if s := mnger.Station.Get(user.ProxyGuid); s != nil {
	// 	proxy.Address = s.Address
	// }
	token, err := middleware.GenerateToken(middleware.UserToken{UserId: u.Id, RoleId: u.RoleId, UserName: u.UserName})
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"tokenName": "Token", "tokenValue": token}, c)
}

// LogoutHandler 注册
func LogoutHandler(c *gin.Context) {
	// c.Redirect(http.StatusFound, pkg.Conf.Local.ServeRoot)
	c.Abort()
}

// CaptchaHandler 图形验证码
func CaptchaHandler(c *gin.Context) {
	driver := base64Captcha.DefaultDriverDigit
	capC := base64Captcha.NewCaptcha(driver, store)
	//以base64编码
	id, b64s, err := capC.Generate()
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"image": b64s, "uid": id}, c)
}

// UserInfoHandler 获取用户信息
func UserInfoHandler(c *gin.Context) {
	tok := middleware.GetUserToken(c)
	var user model.SysUser
	if err := orm.DbFirstById(&user, tok.UserId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(user, c)
}
