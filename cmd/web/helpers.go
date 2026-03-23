package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form/v4"
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

func (app *application) decodeForm(c *gin.Context, dst any) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, c.Request.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}
