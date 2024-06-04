package user

import (
	"fmt"

	"dario.cat/mergo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler interface {
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type UserHandlerImpl struct {
	Service UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &UserHandlerImpl{
		Service: userService,
	}
}

func (handler *UserHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	users, err := handler.Service.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when getting users",
		})
	}

	return ctx.JSON(fiber.Map{
		"data": users,
	})
}

func (handler *UserHandlerImpl) FindById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when getting user",
		})
	}

	user, err := handler.Service.FindById(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Errorf("%w", err),
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func (handler *UserHandlerImpl) Save(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when saving user",
		})
	}

	err := handler.Service.Save(c.Context(), user)

	return c.JSON(fiber.Map{
		"message": "sucessfully saved new user",
		"data":    fmt.Errorf("%w", err),
	})
}

func (handler *UserHandlerImpl) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	var incomingUser User
	err = c.BodyParser(&incomingUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing payload",
		})
	}

	//TODO: handle error properly
	user, err := handler.Service.FindById(c.Context(), id)

	if err := mergo.Merge(&incomingUser, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when updating user",
		})
	}

	err = handler.Service.Update(c.Context(), id, incomingUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when updating user",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("successfully updated user of id %s", id),
	})
}

func (handler *UserHandlerImpl) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user id",
		})
	}

	err = handler.Service.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when preparing to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "successfully deleted user data",
	})
}
