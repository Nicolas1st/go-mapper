package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type addParkingPage struct {
	template *template.Template
}

func NewAddParkingPage(pathToTemplates string) *addParkingPage {
	templateFileNames := []string{"layout.html", "admin/parkings/add-parking-place.html", "admin-navbar.html", "footer.html"}
	for i, fileName := range templateFileNames {
		templateFileNames[i] = path.Join(pathToTemplates, fileName)
	}

	template := template.New("layout")

	template, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		fmt.Println(err)
		panic("Could not build the RemoveParking Page template")
	}

	return &addParkingPage{
		template: template,
	}
}

func (page *addParkingPage) Execute(w http.ResponseWriter) {
	page.template.Execute(w, nil)
}
