package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const ID = 12345
const PASSWD = "mypassword"

func HandleLogin(c *gin.Context) {
	data := struct {
		Id     int    `json:"id"`
		Passwd string `json:"password"`
	}{}
	if c.BindJSON(&data) == nil {
		if data.Id == ID && data.Passwd == PASSWD {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
			})
		}
	}
}
