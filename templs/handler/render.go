package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(c *fiber.Ctx, comp templ.Component, options ...func(*templ.ComponentHandler)) error {
	compHandler := templ.Handler(comp)
	for _, o := range options {
		o(compHandler)
	}
	return adaptor.HTTPHandler(compHandler)(c)
}
