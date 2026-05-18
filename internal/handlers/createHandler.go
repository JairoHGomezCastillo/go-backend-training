package handlers

import (
	"encoding/json"
	"go-backend-training/internal/stack"
	"net/http"
)

func (h *Handler) CreateStackHandler(w http.ResponseWriter, r *http.Request) {
	applicationName := r.PathValue("applicationName")
	var req CreateStackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newStack := stack.Stack{
		Name:        req.Name,
		Description: req.Description,
		Purpose: stack.Purpose{
			Id: req.PurposeId,
		},
		Segment: stack.Segment{
			Id: req.SegmentId,
		},
		ApplicationName: applicationName,
	}

	resp, err := h.uc.PostStack(r.Context(), newStack)
	if err != nil {
		http.Error(w, "failed to create stack", http.StatusInternalServerError)
		return
	}

	stackResponse := CreateStackResponse{
		Id:          resp.Id,
		Name:        resp.Name,
		Description: resp.Description,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(stackResponse); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}
