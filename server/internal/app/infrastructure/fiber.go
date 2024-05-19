package infrastructure

import (
	"time"

	"github.com/dev-nichola/nusaresto/internal/app/menu"
	"github.com/dev-nichola/nusaresto/internal/app/user"
	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	db, err := NewDB()
	helper.PanicIfError(err)

	menuRepository := menu.NewMenuRepository(db)
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

	userRepo := user.NewUserRepository(db)

	v1.Get("/users", userRepo.FindAll)
	v1.Get("/users/:id", userRepo.FindById)
	v1.Post("/users/", userRepo.Save)
	v1.Patch("/users/:id", userRepo.Update)
	v1.Delete("/users/:id", userRepo.Delete)

	app.Listen("localhost:8080")
}
