package service

import (
	"bytes"
	"html/template"
	"path/filepath"
)

type theme struct {
	isReady bool
	tmpl    *template.Template
}

// Render a template with the given contextual data
func (t *theme) Render(templateName string, data interface{}) string {
	t.ready()
	buf := new(bytes.Buffer)
	err := t.tmpl.ExecuteTemplate(buf, templateName, data)
	Logger.Err(err)
	return buf.String()
}

func (t *theme) ready() {
	if t.isReady {
		return
	}

	filenames, _ := filepath.Glob("templates/*.tmpl")
	t.tmpl = template.Must(template.ParseFiles(filenames...))

	t.isReady = true
}
