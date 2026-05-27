package handlers

import (
	"encoding/json"
	"errors"
	"go-backend-training/internal/stack"
	"net/http"
)

func (h *Handler) SearchStackByNameHandler(w http.ResponseWriter, r *http.Request) {
	applicationName := r.PathValue("applicationName")
	stackName := r.PathValue("stackName")

	resp, err := h.uc.GetStackByName(r.Context(), applicationName, stackName)
	if errors.Is(err, stack.ErrStackNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
