package server

import (
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(addr string, handler func(http.ResponseWriter, *http.Request)) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	httpServer := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	resultServer := &Server{
		server: &httpServer,
	}

	return resultServer
}

func (s *Server) Run() {
	log.Printf("starting server at %s\n", s.server.Addr)
}
