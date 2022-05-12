package main

import (
	"this-or-that/handlers"
	"this-or-that/utility"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	utility.Init("my.db")

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("public/html/*")
	r.Static("/static/css", "./public/css")

	r.GET("/", handlers.IndexHandler)
	r.GET("/this", handlers.ThisHandler)
	r.GET("/that", handlers.ThatHandler)
	r.GET("/stats", handlers.StatsHandler)

	r.Run()
}
