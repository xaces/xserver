package service

import (
	"xserver/util"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"
)

type BasePage struct {
	PageNum  int `form:"pageNum"`  // 当前页码
	PageSize int `form:"pageSize"` // 每页数
}

// Where 初始化
func (s *BasePage) Where() *orm.DbWhere {
	var where orm.DbWhere
	return &where
}

func Deletes(v interface{}, c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(v, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func QueryById(v interface{}, c *gin.Context) {
	queryId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	err = orm.DbFirstById(v, queryId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(v, c)
}
