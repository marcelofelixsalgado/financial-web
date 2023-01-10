package utils

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

// var templates *template.Template

func LoadTemplates() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
		// templates: template.Must(template.ParseGlob("web/views/templates/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
