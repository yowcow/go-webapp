package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	ua := c.Request.Header.Get("user-agent")
	c.Header("X-Powered-By", "Gin")
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
		"ua":    ua,
	})
}
