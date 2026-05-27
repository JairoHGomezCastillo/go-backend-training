package stack

import "context"

type UseCaseStack interface {
	CreateStack(ctx context.Context, stack Stack) (Stack, error)
	GetStackByName(ctx context.Context, applicationName string, stackName string) (Stack, error)
	PutStack(ctx context.Context, stack Stack) (Stack, error)
	DeleteStackByName(ctx context.Context, applicationName string, stackName string) error
}

type useCaseStackImpl struct {
	repository RepositoryStack
}

func NewStackUseCase(repository RepositoryStack) UseCaseStack {
	return &useCaseStackImpl{
		repository: repository,
	}
}
