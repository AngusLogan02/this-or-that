package main

import (
	"this-or-that/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/html/*")
	r.Static("/static", "./public/css")
	r.GET("/", handlers.IndexHandler)
	r.Run()
}
