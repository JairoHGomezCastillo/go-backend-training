package router

import (
	"net/http"

	"go-backend-training/internal/handlers"
)

type Server struct {
	router *http.ServeMux
}

func New() *Server {
	router := http.NewServeMux()
	router.HandleFunc("/health", handlers.HealthHandler)
	return &Server{router: router}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
