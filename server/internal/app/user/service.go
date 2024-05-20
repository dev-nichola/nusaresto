package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	Save(ctx context.Context, user User) error
	Update(ctx context.Context, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserServiceImpl struct {
	Repository UserRepository
}

func NewUserService(Repository UserRepository) UserService {
	return &UserServiceImpl{Repository: Repository}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]User, error) {
	menus, err := service.Repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error when getting users: %w", err)
	}

	return menus, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	user, err := service.Repository.FindById(ctx, id)
	if err != nil {
		return user, fmt.Errorf("error when getting user: %w", err)
	}

	return user, nil
}

func (service *UserServiceImpl) Save(ctx context.Context, user User) error {
	err := service.Repository.Save(ctx, user)
	if err != nil {
		return fmt.Errorf("error when saving user: %w", err)
	}

	return err
}

func (UserServiceImpl) Update(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
