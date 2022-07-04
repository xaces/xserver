package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Dict
type Dict struct {
}

// ListHandler 列表
func (o *Dict) ListHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysDictData
	total, _ := p.Where().Model(&model.SysDictData{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// ListExcludeHandler 列表（排除节点）
func (o *Dict) ListExcludeHandler(c *gin.Context) {
	// id, err := ctxQueryInt(c, "id")
	// if err != nil {
	// 	JSONP(StatusError).WriteTo(c)
	// }
	// where := fmt.Sprintf("id != %d", id)
	// var depts []model.Dept
	// orm.DbFindAll(where, depts, "order_num asc")
	ctx.JSONOk(c)
}

// GetHandler 查询详细
func (o *Dict) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysDictData{}, c)
}

// DictTypeHandler
func (o *Dict) DictTypeHandler(c *gin.Context) {
	dtype := c.Param("code")
	var data []model.SysDictData
	total, _ := orm.DbFindBy(&data, "type_code = ?", dtype)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// RoleDeptTreeselectHandler 根据角色ID查询树结构
func (o *Dict) RoleDeptTreeselectHandler(c *gin.Context) {
	ctx.JSONOk(c)
}

// AddHandler 新增字典
func (o *Dict) AddHandler(c *gin.Context) {
	var p model.SysDictData
	//获取参数
	if err := c.ShouldBind(&p.SysDictDataOpt); err != nil {
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
func (o *Dict) UpdateHandler(c *gin.Context) {
	var p model.SysDictData
	//获取参数
	if err := c.ShouldBind(&p.SysDictDataOpt); err != nil {
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
func (o *Dict) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysDictData{}, c)
}

func (o Dict) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.GET("/type/:code", o.DictTypeHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
