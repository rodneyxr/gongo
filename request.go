package gongo

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Request holds information about an HTTP request
type Request struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  httprouter.Params
}

// pack accepts a gongo.Handle and converts the function signature to httprouter.Handle.
// Essentially, all this function does is packs three parameters into a single context struct.
func pack(handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(&Request{w, r, p})
	}
}
