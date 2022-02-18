package controller

import (
	"net/http"
	"xserver/middleware"

	"github.com/gin-gonic/gin"
)

// IndexHandler 首页
func IndexHandler(c *gin.Context) {
	claims := middleware.GetUserToken(c)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"avatar":    "avatar",
		"loginname": claims.UserName,
		"username":  claims.UserName,
	})
}

// IndexHandler 首页
func RootHandler(c *gin.Context) {
	claims := middleware.GetUserToken(c)
	if claims.UserName == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"avatar":    "avatar",
		"loginname": claims.UserName,
		"username":  claims.UserName,
	})
}
