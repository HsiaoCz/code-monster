package main

import (
	"github.com/HsiaoCz/code-monster/templs/handler"
	"github.com/gofiber/fiber/v2"
)

// Echo fiber
// chi gin
// so i use fiber

func main() {
	app := fiber.New()

	userHandler := handler.UserHandler{}

	app.Get("/user", userHandler.HandleUserShow)
	app.Listen(":9011")
}
