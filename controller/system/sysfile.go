package system

import (
	"os"
	"path"
	"path/filepath"
	"xserver/configs"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

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
	total, _ := p.DbWhere().Model(&model.SysFile{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// GetHandler 查询详细
func (o *File) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysFile{}, c)
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
	ctx.JSONOk(c)
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
	ctx.JSONOk(c)
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
	filename := configs.Public("files", fileHead.Filename)
	os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err := c.SaveUploadedFile(fileHead, filename); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	data := &model.SysFile{
		SysFileOpt: model.SysFileOpt{
			Name: fileHead.Filename,
			Path: filename,
			Size: fileHead.Size,
			Desc: fileHead.Filename,
			Type: path.Ext(fileHead.Filename),
		},
		CreatedBy: tok.UserName,
	}
	if err := orm.DbCreate(data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

func (o File) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
	r.POST("/upload", o.UploadHandler)
}
