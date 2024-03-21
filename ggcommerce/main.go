package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	admin := app.Group("/admin")
	productHandler := &ProductHandler{}

	admin.Get("/product", productHandler.HandleGetProduct)
	app.Listen(":4001")
}
