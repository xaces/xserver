package controller

import (
	"errors"
	"net/http"
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
	errLogin          = errors.New("account or pssword error")
	errCaptachaVerify = errors.New("captcha verify error")
	errAuth           = errors.New("authentication error")
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
			user.Password = lo.Password
			user.DeviceIds = "*"
			err = service.SysUserCreate(user)
		}
	}
	if err != nil {
		ctx.JSONWriteError(errLogin, c)
		return
	}
	if user.Enable == 0 {
		ctx.JSONWriteError(errAuth, c)
		return
	}
	verifyPassword := service.SysUserPassword(user, lo.Password)
	if strings.Compare(user.Password, verifyPassword) != 0 {
		ctx.JSONWriteError(errLogin, c)
		return
	}
	tokenNext(c, user)
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, u *model.SysUser) {
	// 获取主账号站点地址
	token, err := middleware.GenerateToken(u.SysUserToken)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"tokenName": "Token", "tokenValue": token}, c)
}

// LogoutHandler 注册
func LogoutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
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
	if err := orm.DbFirstById(&user, tok.Id); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(user, c)
}
