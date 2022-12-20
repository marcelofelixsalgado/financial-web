package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("web/views/*.html"))
	templates = template.Must(templates.ParseGlob("web/views/templates/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		panic(err)
	}
}