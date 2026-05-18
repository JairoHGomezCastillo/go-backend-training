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
	if err != nil {
		if errors.Is(err, stack.ErrStackNotFound) {
			http.Error(w, stack.ErrStackNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, stack.ErrSearchStackByName.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
