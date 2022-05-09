package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type removeParkingPage struct {
	template *template.Template
}

func NewRemoveParkingPage(pathToTemplates string) *removeParkingPage {
	templateFileNames := []string{"layout.html", "admin/parkings/remove-parking-place.html", "admin-navbar.html", "footer.html"}
	for i, fileName := range templateFileNames {
		templateFileNames[i] = path.Join(pathToTemplates, fileName)
	}

	template := template.New("layout")

	template, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		fmt.Println(err)
		panic("Could not build the RemoveParking template")
	}

	return &removeParkingPage{
		template: template,
	}
}

func (page *removeParkingPage) Execute(w http.ResponseWriter) {
	page.template.Execute(w, nil)
}
