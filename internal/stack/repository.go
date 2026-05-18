package stack

import (
	"context"
	"database/sql"
)

type stackRepository struct {
	db *sql.DB
}

type RepositoryStack interface {
	CreateStack(ctx context.Context, stack Stack) (Stack, error)
	SelectStackByName(ctx context.Context, applicationName string, stackName string) (Stack, error)
	SearchAllStacksByApplicationName(ctx context.Context, applicationName string) ([]Stack, error)
	UpdateStack(ctx context.Context, stack Stack) error
	DeleteStackByName(ctx context.Context, applicationName string, stackName string) error
}

func NewStackRepository(db *sql.DB) RepositoryStack {
	return &stackRepository{
		db: db,
	}
}
