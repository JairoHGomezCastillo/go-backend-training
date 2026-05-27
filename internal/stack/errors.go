package stack

import "errors"

var (
	ErrorCreatingStack = errors.New("error creating stack")

	ErrSearchStackByName  = errors.New("error searching stack by name")
	ErrStackNotFound      = errors.New("stack not found")
	ErrStackAlreadyExists = errors.New("stack already exists")
)
