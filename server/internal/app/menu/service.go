package menu

import (
	"context"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/google/uuid"
)

type MenuService interface {
	Create(ctx context.Context, request Menu) (Menu, error)
	FindAll(ctx context.Context) ([]Menu, error)
	FindById(ctx context.Context, id uuid.UUID) (Menu, error)
	Update(ctx context.Context, request Menu) (Menu, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type MenuServiceImpl struct {
	Repository MenuRepository
}

func NewMenuService(repository MenuRepository) MenuService {
	return &MenuServiceImpl{
		Repository: repository,
	}
}

func (service *MenuServiceImpl) Create(ctx context.Context, request Menu) (Menu, error) {
	menu, err := service.Repository.Save(ctx, request)

	helper.PanicIfError(err)

	return menu, nil
}

func (service *MenuServiceImpl) FindAll(ctx context.Context) ([]Menu, error) {
	menus, err := service.Repository.FindAll(ctx)

	helper.PanicIfError(err)

	return menus, nil
}

func (service *MenuServiceImpl) FindById(ctx context.Context, id uuid.UUID) (Menu, error) {
	menu, err := service.Repository.FindById(ctx, id)

	helper.PanicIfError(err)

	return menu, nil
}

func (service *MenuServiceImpl) Update(ctx context.Context, request Menu) (Menu, error) {
	menu, err := service.Repository.Update(ctx, request)

	helper.PanicIfError(err)

	return menu, nil
}

func (service *MenuServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := service.Repository.Delete(ctx, id)

	helper.PanicIfError(err)

	return nil
}
