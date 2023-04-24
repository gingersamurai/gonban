package main

import (
	"gonban/internal/config"
	"gonban/internal/interfaces/storage"
	"gonban/internal/interfaces/webserver"
	"gonban/internal/usecase"
	"log"
)

func main() {
	cfg, err := config.NewConfig(config.ConfigFilePath, config.ConfigFileName)
	if err != nil {
		log.Fatal("config parsing:", err)
	}
	taskStorage, err := storage.NewPostgresTaskStorage(cfg.Postgres)
	if err != nil {
		log.Fatal("connecting to db:", err)
	}
	taskInteractor := usecase.NewTaskInteractor(taskStorage)
	handler := webserver.NewHandler(taskInteractor)
	server := webserver.NewServer(cfg.Server, handler)

	server.Run()
}
