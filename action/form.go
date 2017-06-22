package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleFormBody(c *gin.Context) {
	ver, _ := strconv.Atoi(c.PostForm("version"))
	c.JSON(http.StatusOK, gin.H{"version": ver})
}

func HandleMultipartFormBody(c *gin.Context) {
	_, header, _ := c.Request.FormFile("myupload")
	// Do `io.Copy(os.Stdout, file)` to copy content to somewhere
	hello := c.PostForm("hello")
	c.JSON(http.StatusOK, gin.H{
		"filename": header.Filename,
		"hello":    hello,
	})
}
