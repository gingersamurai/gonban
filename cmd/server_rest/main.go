package main

import (
	"github.com/gingersamurai/gonban/internal/interfaces/storage"
	"github.com/gingersamurai/gonban/internal/interfaces/webserver"
	"github.com/gingersamurai/gonban/internal/usecase"
	"log"
)

func main() {
	taskStorage, err := storage.NewPostgresTaskStorage("host=localhost user=postgres password=15092003 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	taskInteractor := usecase.NewTaskInteractor(taskStorage)
	handler := webserver.NewHandler(taskInteractor)
	server := webserver.NewServer("localhost:8080", handler)

	server.Run()
}
