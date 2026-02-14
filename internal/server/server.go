package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/config"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func New(cfg *config.Config, logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.Root)
	router.HandleFunc("POST /upload", handlers.Upload)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}
