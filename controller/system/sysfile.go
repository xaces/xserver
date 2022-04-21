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
	var p orm.DbPage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysFile
	total, _ := orm.DbByWhere(&model.SysFile{}, p.DbWhere()).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询详细
func (o *File) GetHandler(c *gin.Context) {
	service.QueryById(&model.SysFile{}, c)
}

// AddHandler 新增
func (o *File) AddHandler(c *gin.Context) {
	var p model.SysFile
	//获取参数
	if err := c.ShouldBind(&p.SysFileOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *File) UpdateHandler(c *gin.Context) {
	var p model.SysFile
	//获取参数
	if err := c.ShouldBind(&p.SysFileOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
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
	data := &model.SysFile{
		SysFileOpt: model.SysFileOpt{
			FileName: fileHead.Filename,
			FilePath: filename,
			FileSize: fileHead.Size,
			FileDesc: fileHead.Filename,
		},
		CreatedBy: tok.UserName,
	}
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
