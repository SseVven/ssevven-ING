package model

import "html/template"

var Templates *template.Template

func LoadTemplates() {
	Templates = template.New("templates")
	template.Must(Templates.ParseGlob("client/templates/*.html"))
}
