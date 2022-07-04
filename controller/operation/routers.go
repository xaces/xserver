package operation

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup) {
	Company{}.Routers(r.Group("/company"))
	Fleet{}.Routers(r.Group("/fleet"))
	Vehicle{}.Routers(r.Group("/vehicle"))
}
