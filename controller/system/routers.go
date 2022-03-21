package system

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.RouterGroup) {
	sys := r.Group("/system")
	MenuRouters(sys)
	RoleRouters(sys)
	UserRouters(sys)
	DeptRouters(sys)
	// PostRouters(sys)
	DictDataRouters(sys)
	DictTypeRouters(sys)
	FileRouters(sys)
	NoticeRouters(sys)
	StationRouters(sys)
}