package infrastructure

import (
	"time"

	"github.com/dev-nichola/nusaresto/internal/app/user"
	"github.com/dev-nichola/nusaresto/internal/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

func Run() {

	_, err := NewDB()
	helper.PanicIfError(err)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	v1 := app.Group("v1")

	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HALO HALO HALO")
	})

	userRepo := user.NewUserRepository()

	v1.Get("/users", userRepo.FindAll)
	v1.Get("/users/:id", userRepo.FindById)
	v1.Post("/users/", userRepo.Save)
	v1.Patch("/users/:id", userRepo.Update)
	v1.Delete("/users/:id", userRepo.Delete)

	app.Listen("localhost:8080")

}
