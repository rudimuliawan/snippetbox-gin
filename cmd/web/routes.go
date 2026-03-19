package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) Setup() *gin.Engine {
	r := gin.New()

	r.GET("/home", app.home)
	r.GET("/snippet", app.snippetView)

	return r
}
