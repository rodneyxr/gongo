package gongo

import (
	"log"
	"net/http"
	"strconv"
)

// Server holds information about the server
type Server struct {
	Hostname string
	Port     int
}

// Run starts ListenAndServe on the specified host:port
func (server *Server) Run(router Router) {
	host := server.Hostname
	port := strconv.Itoa(server.Port)
	log.Fatal(http.ListenAndServe(host+":"+port, router.routes))
}
