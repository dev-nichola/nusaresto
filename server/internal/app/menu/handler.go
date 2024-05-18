package menu

import (
	"net/http"

	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuHandler interface {
	Create(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type MenuHandlerImpl struct {
	Service MenuService
}

func NewMenuHandler(service MenuService) MenuHandler {
	return &MenuHandlerImpl{
		Service: service,
	}
}

func (handler *MenuHandlerImpl) Create(ctx *fiber.Ctx) error {
	var request Menu

	if err := ctx.BodyParser(&request); err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	menu, err := handler.Service.Create(ctx.Context(), request)
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"message": "Menu created successfully",
		"data":    menu,
	})
}

func (handler *MenuHandlerImpl) FindAll(ctx *fiber.Ctx) error {
	menus, err := handler.Service.FindAll(ctx.Context())
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Menus fetched successfully",
		"data":    menus,
	})
}

func (handler *MenuHandlerImpl) FindById(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
			"error":   err.Error(),
		})
	}

	menu, err := handler.Service.FindById(ctx.Context(), id)
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Menu fetched successfully",
		"data":    menu,
	})
}

func (handler *MenuHandlerImpl) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
			"error":   err.Error(),
		})
	}

	var request Menu
	if err := ctx.BodyParser(&request); err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	request.ID = id
	menu, err := handler.Service.Update(ctx.Context(), request)
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Menu updated successfully",
		"data":    menu,
	})
}

func (handler *MenuHandlerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
			"error":   err.Error(),
		})
	}

	err = handler.Service.Delete(ctx.Context(), id)
	if err != nil {
		helper.PanicIfError(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Menu deleted successfully",
	})
}
