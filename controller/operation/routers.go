package operation

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.RouterGroup) {
	CompanyRouters(r.Group("/company"))
	FleetRouters(r.Group("/fleet"))
	VehicleRouter(r.Group("/vehicle"))
}
