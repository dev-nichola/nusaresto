package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	QUERY_FIND_ALL     = "SELECT * FROM users"
	QUERY_FIND_BY_ID   = "SELECT * FROM users WHERE id = $1"
	QUERY_SAVE         = "INSERT INTO users ..."              // TODO
	QUERY_UPDATE       = "UPDATE users SET ... WHERE id = $1" // TODO
	QUERY_DELETE_BY_ID = "DELETE FROM users WHERE id = $1"
)

func (user *UserRepositoryImpl) FindAll(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})
}

func (user *UserRepositoryImpl) FindById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
	return c.JSON(fiber.Map{
		"data": 1,
	})
}

func (user *UserRepositoryImpl) Save(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})
}

func (user *UserRepositoryImpl) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})

}

func (user *UserRepositoryImpl) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})

}
