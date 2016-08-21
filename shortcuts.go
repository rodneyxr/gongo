package gongo

import (
	"net/http"

	"github.com/flosch/pongo2"
)

// Render renders a template along with a base template if one exists.
func Render(request *Request, templatePath string, context pongo2.Context) {
	templatePath = templatesDir + "/" + templatePath
	template, err := pongo2.FromFile(templatePath)
	if err != nil {
		panic(err)
	}
	err = template.ExecuteWriter(context, request.Writer)
	if err != nil {
		http.Error(request.Writer, err.Error(), http.StatusInternalServerError)
	}
}
