package html_templating

import (
	"bytes"
	"html/template"
)

func getBufferFromTemplate(data any, tmpl *template.Template) bytes.Buffer {
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return buf
}

func ParseFiles(data any, filenames ...string) bytes.Buffer {
	tmpl := template.Must(template.ParseFiles(filenames...))
	buf := getBufferFromTemplate(data, tmpl)
	return buf
}

func Parse(data any, text string) bytes.Buffer {
	tmpl := template.Must(template.ParseGlob(text))
	buf := getBufferFromTemplate(data, tmpl)
	return buf
}
