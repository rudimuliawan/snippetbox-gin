package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) serverError(c *gin.Context, err error) {
	var (
		uri    = c.Request.RequestURI
		method = c.Request.Method
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	c.Status(http.StatusInternalServerError)
}

func (app *application) clientError(c *gin.Context, status int) {
	http.Error(c.Writer, http.StatusText(status), status)
}
