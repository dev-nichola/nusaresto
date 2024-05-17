package menu

import (
	"context"
	"fmt"
	"time"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MenuRepository interface {
	Save(ctx context.Context, menu Menu) (Menu, error)
	FindAll(ctx context.Context) ([]Menu, error)
	FindById(ctx context.Context, id uuid.UUID) (Menu, error)
	Update(ctx context.Context, menu Menu) (Menu, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type MenuRepositoryImpl struct {
	DB *sqlx.DB
}

const (
	QUERY_SAVE_MENU   = "INSERT INTO menu (name, description) VALUES ($1, $2) RETURNING id, created_at, update_at"
	QUERY_FIND_ALL    = "SELECT id, name, description, created_at, update_at FROM menu"
	QUERY_FIND_BY_ID  = "SELECT id, name, description, created_at, update_at FROM menu WHERE id = $1"
	QUERY_UPDATE_MENU = "UPDATE menu SET name = $1, description = $2, update_at = now() WHERE id = $3 RETURNING id, name, description, created_at, update_at"
	QUERY_DELETE_MENU = "DELETE FROM menu WHERE id = $1"
)

func NewMenuRepository(DB *sqlx.DB) MenuRepository {
	return &MenuRepositoryImpl{
		DB: DB,
	}
}

func (repository *MenuRepositoryImpl) Save(ctx context.Context, menu Menu) (Menu, error) {
	var id uuid.UUID
	var createdAt, updatedAt time.Time

	err := repository.DB.QueryRowContext(ctx, QUERY_SAVE_MENU, menu.Name, menu.Description).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		helper.PanicIfError(err)
		return menu, fmt.Errorf("cant save data menu: %w", err)
	}

	menu.ID = id
	menu.Created_at = createdAt
	menu.Update_at = updatedAt

	return menu, nil
}

func (repository *MenuRepositoryImpl) FindAll(ctx context.Context) ([]Menu, error) {
	var menus []Menu
	err := repository.DB.SelectContext(ctx, &menus, QUERY_FIND_ALL)
	if err != nil {
		helper.PanicIfError(err)
		return nil, fmt.Errorf("cant fetch all data: %w", err)
	}
	return menus, nil
}

func (repository *MenuRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (Menu, error) {
	var menu Menu
	err := repository.DB.GetContext(ctx, &menu, QUERY_FIND_BY_ID, id)
	if err != nil {
		helper.PanicIfError(err)
		return menu, fmt.Errorf("cant fetch data in menu table: %w", err)
	}
	return menu, nil
}

func (repository *MenuRepositoryImpl) Update(ctx context.Context, menu Menu) (Menu, error) {
	var updatedMenu Menu
	err := repository.DB.QueryRowContext(ctx, QUERY_UPDATE_MENU, menu.Name, menu.Description, menu.ID).Scan(&updatedMenu.ID, &updatedMenu.Name, &updatedMenu.Description, &updatedMenu.Created_at, &updatedMenu.Update_at)
	if err != nil {
		helper.PanicIfError(err)
		return updatedMenu, fmt.Errorf("cant update manu: %w", err)
	}
	return updatedMenu, nil
}

func (repository *MenuRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := repository.DB.ExecContext(ctx, QUERY_DELETE_MENU, id)
	if err != nil {
		helper.PanicIfError(err)
		return fmt.Errorf("cant delete data on menu: %w", err)
	}
	return nil
}
