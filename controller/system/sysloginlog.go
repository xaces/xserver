package system

import (
	"xserver/model"

	"github.com/xaces/xutils/ctx"

	"github.com/gin-gonic/gin"
)

// LoginLog
type LoginLog struct {
}

// ListHandler 列表
func (o *LoginLog) ListHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysLoginLog
	total, _ := p.DbWhere().Model(&model.SysLoginLog{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

func (o LoginLog) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
}
