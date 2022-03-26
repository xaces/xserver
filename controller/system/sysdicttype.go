package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// DictType
type DictType struct {
}

// ListHandler 列表
func (o *DictType) ListHandler(c *gin.Context) {
	var param service.DictTypePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysDictType
	total, _ := orm.DbByWhere(&model.SysDictType{}, param.Where()).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 详细
func (o *DictType) GetHandler(c *gin.Context) {
	var data model.SysDictType
	service.QueryById(&data, c)
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
	service.Deletes(&model.SysDictType{}, c)
}

func DictTypeRouters(r *gin.RouterGroup) {
	sysDictType := DictType{}
	r.GET("/dict/type/list", sysDictType.ListHandler)
	r.GET("/dict/type/:id", sysDictType.GetHandler)
	r.POST("/dict/type", sysDictType.AddHandler)
	r.PUT("/dict/type", sysDictType.UpdateHandler)
	r.DELETE("/dict/type/:id", sysDictType.DeleteHandler)
}
