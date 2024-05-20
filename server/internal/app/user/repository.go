package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	Save(ctx context.Context, user User) error
	Update(ctx context.Context, id uuid.UUID, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

const (
	QUERY_FIND_ALL     = "SELECT * FROM users"
	QUERY_FIND_BY_ID   = "SELECT * FROM users WHERE id = $1"
	QUERY_SAVE         = "INSERT INTO users(name, email, password) values($1, $2, $3)" // TODO
	QUERY_UPDATE       = "UPDATE users SET name=$2,email=$3 WHERE id = $1"             // TODO
	QUERY_DELETE_BY_ID = "DELETE FROM users WHERE id = $1"
)

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]User, error) {
	var users []User
	rows, err := repo.DB.QueryxContext(ctx, QUERY_FIND_ALL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("error when querying user: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, fmt.Errorf("error when returning user: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	var user User
	row := repo.DB.QueryRowxContext(ctx, QUERY_FIND_BY_ID, id)
	err := row.StructScan(&user)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error when getting user: %w", err)
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Save(ctx context.Context, user User) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when starting transaction: %w", err)
	}

	_, err = tx.Exec(QUERY_SAVE, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return fmt.Errorf("error when deleting user: %w", err)
	}

	tx.Commit()
	return nil
}

func (userRepo *UserRepositoryImpl) Update(ctx context.Context, id uuid.UUID, user User) error {
	tx, err := userRepo.DB.Begin()
	if err != nil {
		return fmt.Errorf("error when trying to update user: %w", err)
	}

	_, err = tx.ExecContext(ctx, QUERY_UPDATE, id, &user.Name, &user.Email)
	if err != nil {
		return fmt.Errorf("error when updating user data: %w", err)
	}

	tx.Commit()

	return nil
}

func (userRepo *UserRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := userRepo.DB.Begin()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when preparing to delete user")
	}

	_, err = tx.Exec(QUERY_DELETE_BY_ID, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when deleting user")
	}

	return err
}
