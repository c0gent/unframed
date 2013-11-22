package unframed

import (
	"html/template"
)

type tmplDataWrapper struct {
	Data interface{}
	Net  *NetHandle
}

type TemplateStore struct {
	templates     *template.Template
	tmplFileNames []string
}

func (t *TemplateStore) LoadTemplates() {
	t.templates = template.Must(template.ParseFiles(t.tmplFileNames...))
}
func (t *TemplateStore) TemplateFiles(ts ...string) {
	t.tmplFileNames = append(t.tmplFileNames, ts...)
}

func NewTemplateStore() *TemplateStore {
	return new(TemplateStore)
}
