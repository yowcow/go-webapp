package main

import (
	"github.com/yowcow/go-webapp/action"
	"gopkg.in/gin-gonic/gin.v1"
)

func Build(app *gin.Engine) {
	app.GET("/", action.HandleRoot)
	app.POST("/json", action.HandleJsonBody)
	app.POST("/form", action.HandleFormBody)
}
