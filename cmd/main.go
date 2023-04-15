package main

import (
	"gonban/internal/interfaces/storage"
	"gonban/internal/interfaces/webserver"
	"gonban/internal/usecase"
)

func main() {
	memoryTaskStorage := storage.NewMemoryTaskStorage()
	taskInteractor := usecase.NewTaskInteractor(memoryTaskStorage)
	handler := webserver.NewHandler(taskInteractor)
	server := webserver.NewServer("localhost:8080", handler)

	server.Run()
}
