package handlers

import (
	"net/http"
)

// HealthHandler es un manejador para la ruta de salud del servidor.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
