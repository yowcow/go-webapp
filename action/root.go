package action

import (
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
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
	jsonb, _ := ioutil.ReadAll(c.Request.Body)
	data := gin.H{}
	json.Unmarshal(jsonb, &data)
	c.JSON(http.StatusOK, gin.H{"version": data["version"]})
}

func HandleFormBody(c *gin.Context) {
	ver, _ := strconv.Atoi(c.PostForm("version"))
	c.JSON(http.StatusOK, gin.H{"version": ver})
}
