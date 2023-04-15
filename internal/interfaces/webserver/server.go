package webserver

import (
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string, handler *Handler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handle)

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
	log.Printf("started running server")
	log.Fatal(s.server.ListenAndServe())
}
