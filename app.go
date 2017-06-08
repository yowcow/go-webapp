package main

import (
	"flag"
	"github.com/yowcow/go-webapp/action"
	"gopkg.in/gin-gonic/gin.v1"
)

func Build(app *gin.Engine) {
	app.GET("/", action.HandleRoot)
	app.POST("/json", action.HandleJsonBody)
	app.POST("/form", action.HandleFormBody)
}

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", "8888", "Port to listen to")
	flag.Parse()

	router := gin.Default()
	Build(router)

	router.Run(":" + port)
}
