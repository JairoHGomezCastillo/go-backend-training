package handlers

import (
	"go-backend-training/internal/stack"
)

type Handler struct {
	uc stack.UseCaseStack
}

func NewStackHandler(uc stack.UseCaseStack) *Handler {
	return &Handler{
		uc: uc,
	}
}
