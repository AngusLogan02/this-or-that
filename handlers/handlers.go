package handlers

import (
	"fmt"
	"net/http"
	"this-or-that/middlewares"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	this, that := middlewares.GetOptions("Video%20game")
	fmt.Println("this", this)
	fmt.Println("that", that)
	c.HTML(http.StatusOK, "index.html", gin.H{"this": this, "that": that})
}
