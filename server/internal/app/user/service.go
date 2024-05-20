package user

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	Save(ctx context.Context) (User, error)
	Update(ctx context.Context, id uuid.UUID) (User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserServiceImpl struct {
	Repository UserRepository
}

func NewUserService(Repository UserRepository) UserService {
	return &UserServiceImpl{Repository: Repository}
}

func (UserServiceImpl) FindAll(ctx context.Context) ([]User, error) {
	return nil, nil
}

func (UserServiceImpl) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	return User{}, nil
}

func (UserServiceImpl) Save(ctx context.Context) (User, error) {
	return User{}, nil
}

func (UserServiceImpl) Update(ctx context.Context, id uuid.UUID) (User, error) {
	return User{}, nil
}

func (UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
