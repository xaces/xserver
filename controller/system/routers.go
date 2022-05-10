package system

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.RouterGroup) {
	MenuRouters(r.Group("/menu"))
	RoleRouters(r.Group("/role"))
	UserRouters(r.Group("/user"))
	DeptRouters(r.Group("/dept"))
	// PostRouters(sys)
	DictDataRouters(r.Group("/dict/data"))
	DictTypeRouters(r.Group("/dict/type"))
	FileRouters(r.Group("/file"))
	NoticeRouters(r.Group("/notice"))
	StationRouters(r.Group("/station"))
}