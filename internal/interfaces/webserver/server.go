package webserver

import (
	"fmt"
	"gonban/internal/config"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg config.ServerConfig, handler *Handler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handle)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
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
