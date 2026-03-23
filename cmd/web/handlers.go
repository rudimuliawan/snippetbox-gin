package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudimuliawan/snippetbox-gin/internal/validator"
)

type SnippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) home(c *gin.Context) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(c, err)
		return
	}

	data := templateData{
		Snippets: snippets,
	}
	app.render(c, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.serverError(c, err)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	data := templateData{Snippet: snippet}
	app.render(c, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreate(c *gin.Context) {
	form := SnippetCreateForm{Expires: 365}
	data := templateData{Form: form}
	app.render(c, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(c *gin.Context) {
	var form SnippetCreateForm

	err := app.decodeForm(c, &form)
	if err != nil {
		app.clientError(c, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "content", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := templateData{Form: form}
		app.render(c, http.StatusBadRequest, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/snippet/view/%d", id))
}
