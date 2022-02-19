package system

import (
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// File 系统管理字典类型
type File struct {
}

// ListHandler 字典类型列表
func (o *File) ListHandler(c *gin.Context) {
	var param service.FilePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysFile
	totalCount, err := orm.DbPage(&model.SysFile{}, param.Where()).Find(param.PageNum, param.PageSize, &data)
	if err == nil {
		ctx.JSONOk().Write(gin.H{"total": totalCount, "rows": data}, c)
		return
	}
	ctx.JSONWriteError(err, c)
}

// GetHandler 查询字典详细
func (o *File) GetHandler(c *gin.Context) {
	getId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var File model.SysFile
	err = orm.DbFirstById(&File, getId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(File, c)
}

// AddHandler 新增
func (o *File) AddHandler(c *gin.Context) {
	var data model.SysFile
	//获取参数
	if err := c.ShouldBind(&data.SysFileOpt); err != nil {
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
func (o *File) UpdateHandler(c *gin.Context) {
	var data model.SysFile
	//获取参数
	if err := c.ShouldBind(&data.SysFileOpt); err != nil {
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
func (o *File) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysFile{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func FileRouters(r *gin.RouterGroup) {
	sysFile := File{}
	r.GET("/file/list", sysFile.ListHandler)
	r.GET("/file/:id", sysFile.GetHandler)
	r.POST("/file", sysFile.AddHandler)
	r.PUT("/file", sysFile.UpdateHandler)
	r.DELETE("/file/:id", sysFile.DeleteHandler)
}
