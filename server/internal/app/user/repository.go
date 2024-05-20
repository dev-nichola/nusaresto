package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	Save(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, id uuid.UUID, user User) (User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (*UserRepositoryImpl) FindAll(ctx context.Context) ([]User, error) {
	return nil, nil
const (
	QUERY_FIND_ALL     = "SELECT * FROM users"
	QUERY_FIND_BY_ID   = "SELECT * FROM users WHERE id = $1"
	QUERY_SAVE         = "INSERT INTO users(name, email, password) values($1, $2, $3)" // TODO
	QUERY_UPDATE       = "UPDATE users SET name=$2,email=$3 WHERE id = $1"             // TODO
	QUERY_DELETE_BY_ID = "DELETE FROM users WHERE id = $1"
)

}

func (*UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	return User{}, nil
}

func (*UserRepositoryImpl) Save(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (*UserRepositoryImpl) Update(ctx context.Context, id uuid.UUID, user User) (User, error) {
	return User{}, nil
}

func (*UserRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
