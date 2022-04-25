package views

import (
	"net/http"
	"yaroslavl-parkings/web/templates"
)

type viewsResource struct {
	templates *templates.Templates
}

func newViewsResource(pathToTemplates string) *viewsResource {
	return &viewsResource{
		templates: templates.NewTemplates(pathToTemplates),
	}
}

func (viewsResource *viewsResource) SignInView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.SignInPage.Execute(w)
}

func (viewsResource *viewsResource) SignUpView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.SignUpPage.Execute(w)
}

func (viewsResource *viewsResource) makeOrderView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.MakeOrderPage.Execute(w)
}
