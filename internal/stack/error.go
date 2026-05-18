package stack

import "errors"

var (
	ErrCreateStack                  = errors.New("error creating stack")
	ErrStackAlreadyExist            = errors.New("error stack already exists")
	ErrStackNotFound                = errors.New("stack not found")
	ErrSearchStackByName            = errors.New("error searching stack by name")
	ErrSearchAllStacksByApplication = errors.New("error searching all stacks by app")
	ErrUpdateStack                  = errors.New("error updating stack")
	ErrDeleteStack                  = errors.New("error deleting stack")
	ErrStackNotEmpty                = errors.New("error stack not empty")
)
