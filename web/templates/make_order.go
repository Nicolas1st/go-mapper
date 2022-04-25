package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type makeOrderPage struct {
	template *template.Template
}

func NewMakeOrderPage(pathToTemplates string) *makeOrderPage {
	templateFileNames := []string{"layout.html", "make-order-form.html", "signed-in-navbar.html", "footer.html"}
	for i, fileName := range templateFileNames {
		templateFileNames[i] = path.Join(pathToTemplates, fileName)
	}

	template := template.New("layout")

	template, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		fmt.Println(err)
		panic("Could not build the OrderPage template")
	}

	return &makeOrderPage{
		template: template,
	}
}

func (page *makeOrderPage) Execute(w http.ResponseWriter) {
	page.template.Execute(w, nil)
}
