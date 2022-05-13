package pages

import (
	"html/template"
	"net/http"
	"yaroslavl-parkings/api"
)

type Page struct {
	template *template.Template
}

// BuildPage - builds page
// wrapper around build template
func BuildPage(
	pageName,
	pathToTemplates,
	templateToExecute string,
	templateNames ...string,
) *Page {
	return &Page{
		template: buildTemplate(pageName, pathToTemplates, templateToExecute, templateNames...),
	}
}

// Execute - executes the page template without any data provided to it
func (p *Page) Execute(w http.ResponseWriter) error {
	return p.template.Execute(w, struct {
		Endpoints api.Endpoints
	}{
		Endpoints: api.DefaultEndpoints,
	})
}
