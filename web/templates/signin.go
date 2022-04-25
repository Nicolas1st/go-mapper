package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type signInPage struct {
	template *template.Template
}

func NewSignInPage(pathToTemplates string) *signInPage {
	templateFileNames := []string{"layout.html", "sign-in-form.html", "not-signed-in-navbar.html", "footer.html"}
	for i, fileName := range templateFileNames {
		templateFileNames[i] = path.Join(pathToTemplates, fileName)
	}

	template := template.New("layout")

	template, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		fmt.Println(err)
		panic("Could not build the OrderPage template")
	}

	return &signInPage{
		template: template,
	}
}

func (page *signInPage) Execute(w http.ResponseWriter) {
	page.template.Execute(w, nil)
}
