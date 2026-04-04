package data_repos

import (
	"bytes"
	"html/template"
)

func getStringFromTemplate(data any, tmpl *template.Template) string {
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	str := buf.String()
	return str
}

func ParseFiles(data any, filenames ...string) string {
	tmpl := template.Must(template.ParseFiles(filenames...))
	str := getStringFromTemplate(data, tmpl)
	return str
}

func Parse(data any, text string) string {
	tmpl := template.Must(template.ParseGlob(text))
	str := getStringFromTemplate(data, tmpl)
	return str
}
