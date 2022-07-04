package system

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup) {
	Menu{}.Routers(r.Group("/menu"))
	Role{}.Routers(r.Group("/role"))
	User{}.Routers(r.Group("/user"))
	// DeptRouters(r.Group("/dept"))
	Dept{}.Routers(r.Group("/dept"))
	// PostRouters(sys)
	Dict{}.Routers(r.Group("/dict/data"))
	DictType{}.Routers(r.Group("/dict/type"))
	File{}.Routers(r.Group("/file"))
	Notice{}.Routers(r.Group("/notice"))
	Station{}.Routers(r.Group("/station"))
}
