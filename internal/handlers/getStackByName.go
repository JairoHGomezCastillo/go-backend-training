package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (h *Handler) SearchStackByNameHandler(w http.ResponseWriter, r *http.Request) {
	applicationName := r.PathValue("applicationName")
	stackName := r.PathValue("stackName")

	resp, err := h.uc.GetStackByName(r.Context(), applicationName, stackName)
	if err != nil {
		if errors.Is(err, err) {
			http.Error(w, errors2.ErrStackNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, errors2.ErrSearchStackByName.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
