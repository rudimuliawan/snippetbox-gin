package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/rudimuliawan/snippetbox-gin/internal/models"
	"github.com/rudimuliawan/snippetbox-gin/ui"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
	Form     SnippetCreateForm
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) render(c *gin.Context, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template %s doesn't exist", page)
		app.serverError(c, err)
		return
	}

	fmt.Println(len(data.Snippets))

	buff := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buff, "base", data)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.Writer.WriteHeader(status)
	_, _ = buff.WriteTo(c.Writer)
}
