package system

import (
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// DictType 系统管理字典类型
type DictType struct {
}

// ListHandler 字典类型列表
func (o *DictType) ListHandler(c *gin.Context) {
	var param service.DictTypePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysDictType
	totalCount, err := orm.DbPage(&model.SysDictType{}, param.Where()).Find(param.PageNum, param.PageSize, &data)
	if err == nil {
		ctx.JSONOk().Write(gin.H{"total": totalCount, "rows": data}, c)
		return
	}
	ctx.JSONWriteError(err, c)
}

// GetHandler 查询字典详细
func (o *DictType) GetHandler(c *gin.Context) {
	getId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data model.SysDictType
	err = orm.DbFirstById(&data, getId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// AddHandler 新增
func (o *DictType) AddHandler(c *gin.Context) {
	var data model.SysDictType
	//获取参数
	if err := c.ShouldBind(&data.SysDictTypeOpt); err != nil {
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
func (o *DictType) UpdateHandler(c *gin.Context) {
	var data model.SysDictType
	//获取参数
	if err := c.ShouldBind(&data.SysDictTypeOpt); err != nil {
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
func (o *DictType) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysDictType{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func DictTypeRouters(r *gin.RouterGroup) {
	sysDictType := DictType{}
	r.GET("/dict/type/list", sysDictType.ListHandler)
	r.GET("/dict/type/:id", sysDictType.GetHandler)
	r.POST("/dict/type", sysDictType.AddHandler)
	r.PUT("/dict/type", sysDictType.UpdateHandler)
	r.DELETE("/dict/type/:id", sysDictType.DeleteHandler)
}