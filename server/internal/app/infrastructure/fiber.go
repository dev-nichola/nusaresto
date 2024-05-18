package infrastructure

import (
	"time"

	"github.com/dev-nichola/nusaresto/internal/app/menu"
	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	DB, err := NewDB()
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

	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HALO HALO HALO")
	})

	app.Listen("localhost:8080")
}
