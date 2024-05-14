package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
