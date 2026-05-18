package stack

import "context"

func (uc *useCaseStackImpl) PostStack(ctx context.Context, stack Stack) (Stack, error) {
	stack, err := uc.repository.CreateStack(ctx, stack)
	if err != nil {
		return Stack{}, err
	}
	return stack, nil
}

func (uc *useCaseStackImpl) GetStackByName(ctx context.Context, applicationName string, stackName string) (Stack, error) {
	stack, err := uc.repository.SelectStackByName(ctx, applicationName, stackName)
	if err != nil {
		return Stack{}, err
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
