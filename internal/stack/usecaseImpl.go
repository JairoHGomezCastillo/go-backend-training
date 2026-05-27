package stack

import (
	"context"
	"errors"
	"fmt"
	"go-backend-training/internal/mserror"
)

func (uc *useCaseStackImpl) CreateStack(ctx context.Context, stack Stack) (Stack, error) {
	_, err := uc.repository.SelectStackByName(ctx, stack.ApplicationName, stack.Name)
	if errors.Is(err, ErrStackNotFound) {
		stack, err = uc.repository.InsertStack(ctx, stack)
		if err != nil {
			return Stack{}, mserror.Wrap(ErrorCreatingStack, fmt.Sprintf("failed to create stack: %s for application: %s", stack.Name, stack.ApplicationName))
		}
		return stack, nil
	}

	if err != nil {
		return Stack{}, err
	}
	return Stack{}, mserror.Wrap(ErrStackAlreadyExists, fmt.Sprintf("stack: %s already exists for application: %s", stack.Name, stack.ApplicationName))
}

func (uc *useCaseStackImpl) GetStackByName(ctx context.Context, applicationName string, stackName string) (Stack, error) {
	stack, err := uc.repository.SelectStackByName(ctx, applicationName, stackName)
	if errors.Is(err, ErrStackNotFound) {
		return Stack{}, mserror.Wrap(ErrStackNotFound, fmt.Sprintf("stack: %s not found for application: %s", stackName, applicationName))
	}
	if err != nil {
		return Stack{}, mserror.Wrap(ErrSearchStackByName, fmt.Sprintf("failed to search stack: %s for application: %s", stackName, applicationName))
	}
	return stack, nil
}

func (uc *useCaseStackImpl) PutStack(ctx context.Context, stack Stack) (Stack, error) {
	err := uc.repository.UpdateStack(ctx, stack)
	if err != nil {
		return Stack{}, err
	}
	return stack, nil
}

func (uc *useCaseStackImpl) DeleteStackByName(ctx context.Context, applicationName string, stackName string) error {
	err := uc.repository.DeleteStackByName(ctx, applicationName, stackName)
	if err != nil {
		return err
	}
	return nil
}
