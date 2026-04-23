package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := HealthResponse{Status: "ok"}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error escribiendo respuesta: %v", err)
	}
}
