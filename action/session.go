package action

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
