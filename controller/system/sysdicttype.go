package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

	"github.com/gin-gonic/gin"
)

// DictType
type DictType struct {
}

// ListHandler 列表
func (o *DictType) ListHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysDictType
	total, _ := p.Where().Model(&model.SysDictType{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// GetHandler 详细
func (o *DictType) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysDictType{}, c)
}

// AddHandler 新增
func (o *DictType) AddHandler(c *gin.Context) {
	var p model.SysDictType
	//获取参数
	if err := c.ShouldBind(&p.SysDictTypeOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *DictType) UpdateHandler(c *gin.Context) {
	var p model.SysDictType
	//获取参数
	if err := c.ShouldBind(&p.SysDictTypeOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *DictType) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysDictType{}, c)
}

func (o DictType) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
