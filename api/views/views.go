package views

import (
	"fmt"
	"net/http"
	"yaroslavl-parkings/api"
	"yaroslavl-parkings/api/views/pages"
	"yaroslavl-parkings/persistence/model"
)

// dependencies
type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*model.Session, bool)
}

type viewsDependencies struct {
	pages    *pages.Pages
	sessions SessionsInterface
}

type Views struct {
	SignIn             http.HandlerFunc
	SignUp             http.HandlerFunc
	MakeOrder          http.HandlerFunc
	AddParkingPlace    http.HandlerFunc
	RemoveParkingPlace http.HandlerFunc
}

func NewViews(
	pathToTemplates string,
	sessions SessionsInterface,
) *Views {
	dependencies := &viewsDependencies{
		pages:    pages.NewPages(pathToTemplates),
		sessions: sessions,
	}

	return &Views{
		SignIn:             dependencies.SignIn,
		SignUp:             dependencies.SignUp,
		MakeOrder:          dependencies.MakeOrder,
		AddParkingPlace:    dependencies.AddParkingPlace,
		RemoveParkingPlace: dependencies.RemoveParkingPlace,
	}
}

// SignIn - serves SignIn page
func (d *viewsDependencies) SignIn(w http.ResponseWriter, r *http.Request) {
	if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		fmt.Println(d.pages.Public.SignIn.Execute(w))
	}
}

// SignUp - serves SignUp page
func (d *viewsDependencies) SignUp(w http.ResponseWriter, r *http.Request) {
	if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		fmt.Println(d.pages.Public.SignUp.Execute(w))
	}
}

// MakeOrder - serves MakeOrder page
func (d *viewsDependencies) MakeOrder(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.MakeOrder.Execute(w))
	} else if api.IsAuth(d.sessions, r) {
		fmt.Println(d.pages.Private.MakeOrder.Execute(w))
	} else {
		// redirection to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// AddParkingPlace - serves AddParkingPlace page
func (d *viewsDependencies) AddParkingPlace(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.AddParking.Execute(w))
	} else if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		// redirect to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}

// RemoveParkingPlace - serves RemoveParkingPlace page
func (d *viewsDependencies) RemoveParkingPlace(w http.ResponseWriter, r *http.Request) {
	if api.IsAuthAndAdmin(d.sessions, r) {
		fmt.Println(d.pages.Admin.RemoveParking.Execute(w))
	} else if api.IsAuth(d.sessions, r) {
		// redirect to the page for authenticated users
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
	} else {
		// redirect to the page for unauthenticated users
		http.Redirect(w, r, api.DefaultEndpoints.LoginPage, http.StatusSeeOther)
	}
}
