package router

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"xserver/configs"
	"xserver/controller"
	"xserver/controller/operation"
	"xserver/controller/system"
	"xserver/middleware"

	"github.com/gin-gonic/gin"
)

type options struct {
	Timeout int64
	Port    uint16
	Root    string
	View    string
}

var (
	s *http.Server
	o options
)

func loadStaticResources(r *gin.Engine) {
	r.Static("admin", o.View+"/admin")
	r.Static("component", o.View+"/component")
	r.Static("config", o.View+"/config")
	r.Static("view", o.View+"/view")
	r.LoadHTMLGlob(o.View + "/*.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET(o.Root, controller.RootHandler)
	r.GET("/", controller.RootHandler)
	r.GET("/index", controller.IndexHandler)
}

func initRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	loadStaticResources(r)
	r.Use(gin.Logger()) // 日志
	r.Use(middleware.Cors())
	root := r.Group(o.Root + "/api")
	root.GET("/captcha", controller.CaptchaHandler)
	root.POST("/login", controller.LoginHandler)
	root.POST("/logout", controller.LogoutHandler)
	root.GET("/devices/:guid", controller.DevicesHandler)
	root.GET("/download/:file", controller.DownloadHandler)
	root.GET("/ws", controller.WsHandler)
	jwt := root.Group("")
	jwt.Use(middleware.JWTAuth())
	jwt.Any("/station/*api", controller.ProxyHandler("/station/api"))
	system.Routers(jwt.Group("/system"))
	operation.Routers(jwt.Group("/operation"))
	return r
}

// New
func Run() error {
	if err := configs.GViper.UnmarshalKey("http", &o); err != nil {
		return err
	}
	r := initRouter()
	address := fmt.Sprintf(":%d", o.Port)
	s = &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    time.Duration(o.Timeout) * time.Second,
		WriteTimeout:   time.Duration(o.Timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()
	return nil
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	return nil
}
