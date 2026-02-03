package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/config"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	cfg := config.Load()

	logger := log.Default()
	server := server.New(cfg, logger)

	logger.Println("Server is running on port " + cfg.Port)
	if err := server.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
