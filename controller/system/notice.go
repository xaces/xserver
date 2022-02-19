package system

import (
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Notice 系统管理字典类型
type Notice struct {
}

// ListHandler 字典类型列表
func (o *Notice) ListHandler(c *gin.Context) {
	var param service.NoticePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysNotice
	totalCount, err := orm.DbPage(&model.SysNotice{}, param.Where()).Find(param.PageNum, param.PageSize, &data)
	if err == nil {
		ctx.JSONOk().Write(gin.H{"total": totalCount, "rows": data}, c)
		return
	}
	ctx.JSONWriteError(err, c)
}

// GetHandler 查询字典详细
func (o *Notice) GetHandler(c *gin.Context) {
	getId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data model.SysNotice
	err = orm.DbFirstById(&data, getId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// AddHandler 新增
func (o *Notice) AddHandler(c *gin.Context) {
	var data model.SysNotice
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Notice) UpdateHandler(c *gin.Context) {
	var data model.SysNotice
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Notice) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysNotice{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func NoticeRouters(r *gin.RouterGroup) {
	sysNotice := Notice{}
	r.GET("/notice/list", sysNotice.ListHandler)
	r.GET("/notice/:id", sysNotice.GetHandler)
	r.POST("/notice", sysNotice.AddHandler)
	r.PUT("/notice", sysNotice.UpdateHandler)
	r.DELETE("/notice/:id", sysNotice.DeleteHandler)
}
