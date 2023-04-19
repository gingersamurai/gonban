package main

import (
	"github.com/gingersamurai/gonban/internal/interfaces/storage"
	"github.com/gingersamurai/gonban/internal/interfaces/webserver"
	"github.com/gingersamurai/gonban/internal/usecase"
)

func main() {
	memoryTaskStorage := storage.NewMemoryTaskStorage()
	taskInteractor := usecase.NewTaskInteractor(memoryTaskStorage)
	handler := webserver.NewHandler(taskInteractor)
	server := webserver.NewServer("localhost:8080", handler)

	server.Run()
}
