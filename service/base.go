package service

import (
	"xserver/util"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"
)

func Deletes(v interface{}, c *gin.Context) {
	idstr := c.Param("id")
	if idstr != "" {
		ids := util.StringToIntSlice(idstr, ",")
		if err := orm.DbDeleteByIds(v, ids); err != nil {
			ctx.JSONWriteError(err, c)
			return
		}
	}
	ctx.JSONOk(c)
}

func QueryById(v interface{}, c *gin.Context) {
	queryId := ctx.ParamUInt(c, "id")
	if err := orm.DbFirstById(v, queryId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONWriteData(v, c)
}
