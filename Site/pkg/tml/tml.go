package tml

import "html/template"

//GetTemplates ...
func GetTemplates() *template.Template {
	return template.Must(template.ParseGlob("web/*.gohtml"))
}
