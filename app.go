package main

import (
	"flag"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yowcow/go-webapp/action"
)

var (
	port string
)

func Build(router *gin.Engine) {
	router.Static("/static", "./static")

	store := sessions.NewCookieStore([]byte("hogefuga"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", action.HandleRoot)
	router.POST("/session", action.HandleSetSession)
	router.GET("/session", action.HandleGetSession)
	router.POST("/json", action.HandleJsonBody)
	router.POST("/form", action.HandleFormBody)
	router.POST("/form-multipart", action.HandleMultipartFormBody)
	router.POST("/login", action.HandleLogin)
}

func main() {
	flag.StringVar(&port, "port", "8888", "Port to listen to")
	flag.Parse()

	router := gin.Default()
	Build(router)

	router.Run(":" + port)
}
