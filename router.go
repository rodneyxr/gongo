package gongo

import "github.com/julienschmidt/httprouter"

// Router wraps julienschmidt/httprouter.Router
type Router struct {
	routes *httprouter.Router
}

// Handle is a function that can be registered to a route to handle HTTP requests.
type Handle func(request *Request)

// PUT adds a put route to router
func (r *Router) PUT(path string, handle Handle) {
	r.routes.PUT(path, pack(handle))
}

// POST adds a post route to router
func (r *Router) POST(path string, handle Handle) {
	r.routes.POST(path, pack(handle))
}

// GET adds a get route to router
func (r *Router) GET(path string, handle Handle) {
	r.routes.GET(path, pack(handle))
}

// PATCH adds a patch route to router
func (r *Router) PATCH(path string, handle Handle) {
	r.routes.PATCH(path, pack(handle))
}

// DELETE adds a delete route to router
func (r *Router) DELETE(path string, handle Handle) {
	r.routes.DELETE(path, pack(handle))
}
