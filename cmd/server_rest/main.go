package main

import (
	"github.com/gingersamurai/gonban/internal/config"
	"github.com/gingersamurai/gonban/internal/interfaces/storage"
	"github.com/gingersamurai/gonban/internal/interfaces/webserver"
	"github.com/gingersamurai/gonban/internal/usecase"
	"log"
)

const (
	configFilePath = "./internal/config/"
	configFileName = "config"
)

func main() {
	cfg, err := config.NewConfig(configFilePath, configFileName)
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
