package router

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"xserver/controller"
	"xserver/controller/system"
	"xserver/middleware"

	"github.com/gin-gonic/gin"
)

func loadStaticResources(r *gin.Engine, o *Option) {
	distPath := o.View
	r.Static("admin", distPath+"/admin")
	r.Static("component", distPath+"/component")
	r.Static("config", distPath+"/config")
	r.Static("view", distPath+"/view")
	r.LoadHTMLGlob(distPath + "/*.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET(o.Root, controller.RootHandler)
	r.GET("/index", controller.IndexHandler)
}

func initRouter(o *Option) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	loadStaticResources(r, o)
	r.Use(gin.Logger()) // 日志
	r.Use(middleware.Cors())
	root := r.Group(o.Root)
	root.GET("/captcha", controller.CaptchaHandler)
	root.POST("/login", controller.LoginHandler)
	root.POST("/logout", controller.LogoutHandler)
	root.POST("/file/upload", controller.FileUploadHandler)
	jwt := root.Group("")
	jwt.Use(middleware.JWTAuth())
	// jwt.GET("/getRouters", controller.GetRoutesHandler)
	sys := jwt.Group("/system")
	system.MenuRouters(sys)
	system.RoleRouters(sys)
	system.UserRouters(sys)
	system.DeptRouters(sys)
	// system.PostRouters(sys)
	system.DictDataRouters(sys)
	system.DictTypeRouters(sys)
	system.FileRouters(sys)
	system.NoticeRouters(sys)
	return r
}

type Option struct {
	Timeout time.Duration
	Port    uint16
	Root    string
	View    string
}

var (
	s *http.Server
)

// New
func Run(o *Option) *http.Server {
	r := initRouter(o)
	address := fmt.Sprintf(":%d", o.Port)
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    o.Timeout * time.Second,
		WriteTimeout:   o.Timeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()
	return s
}

func Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	return nil
}
