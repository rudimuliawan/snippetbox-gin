package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) Setup() *gin.Engine {
	r := gin.New()

	r.Static("/static", "./ui/static")

	r.GET("", app.home)
	r.GET("/snippet/view/:id", app.snippetView)
	r.GET("/snippet/create", app.snippetCreate)
	r.POST("/snippet/create", app.snippetCreatePost)

	return r
}
