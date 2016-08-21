package gongo

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/julienschmidt/httprouter"
)

// Config
const hostname = "localhost"
const port = 8000
const staticDir = "static"
const templatesDir = "templates"
const templateExtention = "html"

// Gongo holds all components necessary to make the webserver function.
type Gongo struct {
	router    Router
	server    Server
	templates map[string]*pongo2.Template
}

var g Gongo

// New returns an instance of Gongo
func New() Gongo {
	g = Gongo{
		router: Router{
			routes: httprouter.New(),
		},

		server: Server{
			Hostname: hostname,
			Port:     port,
		},

		templates: mustAllTemplates(),
	}
	g.router.routes.NotFound = http.FileServer(http.Dir(staticDir))

	return g
}

// RoutesFunc uses the function passed in to setup the Gongo router
func (gongo *Gongo) RoutesFunc(handle func(r *Router)) {
	handle(&gongo.router)
}

// Run starts the server
func (gongo *Gongo) Run() {
	gongo.server.Run(gongo.router)
}
