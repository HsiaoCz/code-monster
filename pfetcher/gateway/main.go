package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/get-price/{ticker}", HandleGetPrice)
	app.Listen(":9091")
}
