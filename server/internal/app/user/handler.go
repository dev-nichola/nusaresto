package user

import (
	"database/sql"
	"fmt"
	"log"

	"dario.cat/mergo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	QUERY_FIND_ALL     = "SELECT * FROM users"
	QUERY_FIND_BY_ID   = "SELECT * FROM users WHERE id = $1"
	QUERY_SAVE         = "INSERT INTO users(name, email, password) values($1, $2, $3)" // TODO
	QUERY_UPDATE       = "UPDATE users SET name=$2,email=$3 WHERE id = $1"             // TODO
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
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when saving user",
		})
	}

	tx, err := userRepo.DB.Begin()
	if err != nil {
		tx.Rollback()
		log.Println("error when starting transaction")
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when saving user",
		})
	}

	result, err := tx.Exec(QUERY_SAVE, &user.Name, &user.Email, &user.Password)

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when saving user",
		})
	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"message": "sucessfully saved new user",
		"data":    result,
	})
}

func (userRepo *UserRepositoryImpl) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	var user User

	result := userRepo.DB.QueryRowx(QUERY_FIND_BY_ID, id)
	err = result.StructScan(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when checking existing user",
		})
	}

	var incomingUser User
	err = c.BodyParser(&incomingUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing payload",
		})
	}

	if err := mergo.Merge(&incomingUser, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when updating user",
		})
	}

	tx, err := userRepo.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when trying to update user",
		})
	}

	_, err = tx.Exec(QUERY_UPDATE, id, &incomingUser.Name, &incomingUser.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when updating user data",
		})
	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("successfully updated user of id %s", id),
	})
}

func (userRepo *UserRepositoryImpl) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	tx, err := userRepo.DB.Begin()
	if err != nil {
		tx.Rollback()

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when preparing to delete user",
		})
	}

	_, err = tx.Exec(QUERY_DELETE_BY_ID, id)
	if err != nil {
		tx.Rollback()

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when deleting user",
		})
	}

	tx.Commit()
	return c.JSON(fiber.Map{
		"message": "successfully deleted user data",
	})
}
