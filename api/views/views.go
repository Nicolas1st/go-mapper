package views

import (
	"net/http"
	"yaroslavl-parkings/api/views/templates"
)

type viewsResource struct {
	templates *templates.Templates
}

func NewViews(templates *templates.Templates) *viewsResource {
	return &viewsResource{
		templates: templates,
	}
}

func (viewsResource *viewsResource) SignInView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.SignInPage.Execute(w)
}

func (viewsResource *viewsResource) SignUpView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.SignUpPage.Execute(w)
}

func (viewsResource *viewsResource) MakeOrderView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.MakeOrderPage.Execute(w)
}

func (viewsResource *viewsResource) AddParkingPlaceView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.AddParkingPage.Execute(w)
}

func (viewsResource *viewsResource) RemoveParkingPlaceView(w http.ResponseWriter, r *http.Request) {
	viewsResource.templates.RemoveParkingPage.Execute(w)
}
