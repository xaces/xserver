package operation

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.RouterGroup) {
	oper := r.Group("/operation")
	CompanyRouters(oper)
	FleetRouters(oper)
	VehicleRouter(oper)
}
