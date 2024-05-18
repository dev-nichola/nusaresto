package infrastructure

import (
	"time"

	"github.com/dev-nichola/nusaresto/internal/app/menu"
	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	helper.PanicIfError(err)

	menuRepository := menu.NewMenuRepository(DB)
	menuService := menu.NewMenuService(menuRepository)
	menuHandler := menu.NewMenuHandler(menuService)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	v1 := app.Group("v1")

	v1.Post("/menu", menuHandler.Create)
	v1.Get("/menus", menuHandler.FindAll)
	v1.Get("/menu/:id", menuHandler.FindById)
	v1.Put("/menu/:id", menuHandler.Update)
	v1.Delete("/menu/:id", menuHandler.Delete)

	app.Listen("localhost:8080")
}
