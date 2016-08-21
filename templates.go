package gongo

import (
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
)

// mustAllTemplates compiles all templates under the base template directory and returns them in a map.
// The map should be saved so that this function is only called once at startup.
func mustAllTemplates() map[string]*pongo2.Template {
	templates := make(map[string]*pongo2.Template)
	filepath.Walk(templatesDir, storeTemplate(templates))
	return templates
}

// storeTemplate compiles and stores the template that was passed by Walk.
func storeTemplate(templates map[string]*pongo2.Template) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			templates[path] = pongo2.Must(pongo2.FromFile(path))
		}
		return err
	}
}
