package action

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleRoot(c *gin.Context) {
	ua := c.Request.Header.Get("user-agent")
	c.Header("X-Powered-By", "Gin")
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
		"ua":    ua,
	})
}

func HandleJsonBody(c *gin.Context) {
	data := struct {
		Ver int `json:"version"`
	}{}
	if c.BindJSON(&data) == nil {
		c.JSON(http.StatusOK, gin.H{"version": data.Ver})
	}
}

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

func HandleSetSession(c *gin.Context) {
	store := sessions.Default(c)
	val := c.PostForm("val")
	store.Set("val", val)
	store.Save()
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func HandleGetSession(c *gin.Context) {
	store := sessions.Default(c)
	val := store.Get("val")
	if val == nil {
		c.String(http.StatusNoContent, "")
	} else {
		c.String(http.StatusOK, val.(string))
	}
}
