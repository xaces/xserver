package system

import (
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// File
type File struct {
}

// ListHandler 列表
func (o *File) ListHandler(c *gin.Context) {
	var param service.BasePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysFile
	total, _ := orm.DbPage(&model.SysFile{}, param.Where()).Find(param.PageNum, param.PageSize, &data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询详细
func (o *File) GetHandler(c *gin.Context) {
	var data model.SysFile
	service.QueryById(&data, c)
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
	service.Deletes(&model.SysFile{}, c)
}

func (o *File) UploadHandler(c *gin.Context) {
	fileHead, err := c.FormFile("file")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	filename := "" + fileHead.Filename
	// TODO save db
	if err := c.SaveUploadedFile(fileHead, filename); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	data := &model.SysFile{}
	data.FileName = fileHead.Filename
	data.FilePath = filename
	data.FileSize = fileHead.Size
	data.FileDesc = fileHead.Filename
	data.CreatedBy = tok.UserName
	if err := orm.DbCreate(data); err != nil {
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
	r.POST("/file/upload", sysFile.UploadHandler)
}
