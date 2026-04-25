package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type myServer struct {
	Server *http.Server
	logger *log.Logger
}

func NewRouter(l *log.Logger) *myServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Get)
	mux.HandleFunc("/upload", handlers.Upload)

	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &myServer{Server: &s, logger: l}
}
