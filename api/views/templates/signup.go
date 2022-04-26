package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type signUpPage struct {
	template *template.Template
}

func NewSignUpPage(pathToTemplates string) *signUpPage {
	templateFileNames := []string{"layout.html", "sign-up-form.html", "not-signed-in-navbar.html", "footer.html"}
	for i, fileName := range templateFileNames {
		templateFileNames[i] = path.Join(pathToTemplates, fileName)
	}

	template := template.New("layout")

	template, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		fmt.Println(err)
		panic("Could not build the OrderPage template")
	}

	return &signUpPage{
		template: template,
	}
}

func (page *signUpPage) Execute(w http.ResponseWriter) {
	page.template.Execute(w, nil)
}
