package system

import (
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Dict 系统管理字典
type Dict struct {
}

// ListHandler 字典列表
func (o *Dict) ListHandler(c *gin.Context) {
	var param service.DictDataPage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var dicts []model.SysDictData
	totalCount, err := orm.DbPage(&model.SysDictData{}, param.Where()).Find(param.PageNum, param.PageSize, &dicts)
	if err == nil {
		ctx.JSONOk().Write(gin.H{"total": totalCount, "rows": dicts}, c)
		return
	}
	ctx.JSONWriteError(err, c)
}

// ListExcludeHandler 字典列表（排除节点）
func (o *Dict) ListExcludeHandler(c *gin.Context) {
	// id, err := ctxQueryInt(c, "id")
	// if err != nil {
	// 	JSONP(StatusError).WriteTo(c)
	// }
	// where := fmt.Sprintf("id != %d", id)
	// var depts []model.Dept
	// orm.DbFindAll(where, depts, "order_num asc")
	ctx.JSONOk().WriteTo(c)
}

// GetHandler 查询字典详细
func (o *Dict) GetHandler(c *gin.Context) {
	getId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data model.SysDictData
	err = orm.DbFirstById(&data, getId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// DictTypeHandler 根据字典类型查询字典数据信息
func (o *Dict) DictTypeHandler(c *gin.Context) {
	dtype := ctx.ParamString(c, "id")
	var data []model.SysDictData
	_, err := orm.DbFindBy(&data, "dict_type like ?", dtype)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(data, c)
}

// RoleDeptTreeselectHandler 根据角色ID查询字典树结构
func (o *Dict) RoleDeptTreeselectHandler(c *gin.Context) {
	ctx.JSONOk().WriteTo(c)
}

// AddHandler 新增字典
func (o *Dict) AddHandler(c *gin.Context) {
	var data model.SysDictData
	//获取参数
	if err := c.ShouldBind(&data.SysDictDataOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改字典
func (o *Dict) UpdateHandler(c *gin.Context) {
	var data model.SysDictData
	//获取参数
	if err := c.ShouldBind(&data.SysDictDataOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除字典
func (o *Dict) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysDictData{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func DictDataRouters(r *gin.RouterGroup) {
	sysDict := Dict{}
	r.GET("/dict/data/list", sysDict.ListHandler)
	r.GET("/dict/data/:id", sysDict.GetHandler)
	r.GET("/dict/data/type/:id", sysDict.DictTypeHandler)
	r.POST("/dict/data", sysDict.AddHandler)
	r.PUT("/dict/data", sysDict.UpdateHandler)
	r.DELETE("/dict/data/:id", sysDict.DeleteHandler)
}
