package menu

import (
	"context"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MenuRepository interface {
	FindAll()
	FindById()
	Save(ctx context.Context, tx *sqlx.Tx, menu Menu) Menu
	Update()
	Delete()
}

type MenuRepositoryImpl struct {
}

const (
	QUERY_SAVE_MENU = "INSERT INTO MENU (name) VALUES (?)"
)

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) FindAll() {

}

func (repository *MenuRepositoryImpl) FindById() {

}

func (repository *MenuRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, menu Menu) Menu {
	result, err := tx.ExecContext(ctx, QUERY_SAVE_MENU, menu.Name)
	helper.PanicIfError(err)

	newUUID := uuid.New(uint(id))

	return menu
}

func (repository *MenuRepositoryImpl) Update() {

}

func (repository *MenuRepositoryImpl) Delete() {

}
