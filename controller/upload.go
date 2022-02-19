package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/ctx"
)

func FileUploadHandler(c *gin.Context) {
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
	ctx.JSONOk().WriteTo(c)
}
