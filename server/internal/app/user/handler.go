package user

import (
	"database/sql"
	"log"

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

func (userRepo UserRepositoryImpl) FindAll(c *fiber.Ctx) error {
	var users []User

	rows, err := userRepo.DB.Queryx(QUERY_FIND_ALL)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when getting users",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.StructScan(&user)

		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"message": "error when getting users",
			})
		}

		users = append(users, user)
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (userRepo *UserRepositoryImpl) FindById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when getting user",
		})
	}

	rows := userRepo.DB.QueryRowx(QUERY_FIND_BY_ID, id)

	var user User
	err = rows.StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when getting user data",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (userRepo *UserRepositoryImpl) Save(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})
}

func (userRepo *UserRepositoryImpl) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})

}

func (userRepo *UserRepositoryImpl) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": 1,
	})

}
