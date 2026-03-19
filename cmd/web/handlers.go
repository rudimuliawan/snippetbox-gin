package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) home(c *gin.Context) {
	c.Writer.Write([]byte("Hello"))
}

func (app *application) snippetView(c *gin.Context) {
	snippets, err := app.snippet.Latest()
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, snippets)
}
