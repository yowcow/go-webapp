package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleJsonBody(c *gin.Context) {
	data := struct {
		Ver int `json:"version"`
	}{}
	if c.BindJSON(&data) == nil {
		c.JSON(http.StatusOK, gin.H{"version": data.Ver})
	}
}
