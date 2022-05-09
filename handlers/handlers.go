package handlers

import (
	"net/http"
	"this-or-that/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	this, that := middlewares.GetOptions("Video%20game")
	session := sessions.Default(c)
	session.Set("this", this)
	session.Set("that", that)
	session.Save()
	c.HTML(http.StatusOK, "index.html", gin.H{"this": this, "that": that})
}

func ThisHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ThatHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}
