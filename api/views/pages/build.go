package pages

import (
	"fmt"
	"html/template"
	"path"
)

func buildTemplate(
	pageName,
	pathToTemplates,
	templateToExecute string,
	templateNames ...string,
) *template.Template {
	// stop program execution if it's not possible to build the template
	if len(templateNames) == 0 {
		panic("Can not build page with zerof files provided")
	}

	// prepend filepath
	withPaths := []string{}
	for _, fileName := range templateNames {
		withPaths = append(withPaths, path.Join(pathToTemplates, fileName))
	}

	// creating template
	template := template.New(templateToExecute)
	template, err := template.ParseFiles(withPaths...)

	// stop the program exeuction if it's not possible to build the template
	if err != nil {
		panic(fmt.Sprintf("Could not parse files for page %s", pageName))
	}

	return template
}
