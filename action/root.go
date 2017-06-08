package action

import (
	"gopkg.in/gin-gonic/gin.v1"
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
