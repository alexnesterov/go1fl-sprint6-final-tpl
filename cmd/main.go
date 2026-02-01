package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()
	server := server.New(logger)

	logger.Println("Server is running")
	if err := server.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
