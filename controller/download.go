package controller

import (
	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	filename := c.GetString("file")
	c.File(filename)
}
