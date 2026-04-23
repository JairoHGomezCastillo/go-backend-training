package server

import (
	"net/http"

	"go-backend-training/cmd/api/handlers"
)

// Server representa el servidor HTTP.
type Server struct {
	router *http.ServeMux
}

// New crea una nueva instancia del servidor.
func New() *Server {
	router := http.NewServeMux()
	router.HandleFunc("/health", handlers.HealthHandler)

	return &Server{router: router}
}

// Start inicia el servidor en la dirección especificada.
func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
