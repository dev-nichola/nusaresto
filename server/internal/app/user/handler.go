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
	var userData User
	err := user.DB.QueryRowx(QUERY_FIND_ALL).StructScan(&userData)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "error when getting user data",
		})
	}

	return c.JSON(fiber.Map{
		"data": userData,
	})
}

func (user *UserRepositoryImpl) FindById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "apa kek",
		})
	}

	rows, err := user.DB.Queryx(QUERY_FIND_BY_ID, id)
	defer rows.Close()

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "error when getting users",
		})
	}

	var users []User

	for rows.Next() {
		var userData User
		err := rows.Scan(&userData)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error when getting users data",
			})
		}

		users = append(users, userData)
	}

	return c.JSON(fiber.Map{
		"data": users,
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
