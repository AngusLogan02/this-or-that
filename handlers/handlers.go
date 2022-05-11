package handlers

import (
	"net/http"
	"this-or-that/middlewares"
	"this-or-that/utility"

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
	session := sessions.Default(c)
	key := session.Get("this")
	utility.Increment(key.(string))

	this, that := middlewares.GetOptions("Video%20game")
	session.Set("this", this)
	session.Set("that", that)
	session.Save()

	c.HTML(http.StatusOK, "index.html", gin.H{"this": this, "that": that})
}

func ThatHandler(c *gin.Context) {
	session := sessions.Default(c)
	key := session.Get("that")
	utility.Increment(key.(string))

	this, that := middlewares.GetOptions("Video%20game")
	session.Set("this", this)
	session.Set("that", that)
	session.Save()

	c.HTML(http.StatusOK, "index.html", gin.H{"this": this, "that": that})
}

func StatsHandler(c *gin.Context) {
	category := "Video_game"
	utility.GenGraph(category)
	c.HTML(http.StatusOK, "stats.html", gin.H{"category": category})
}
