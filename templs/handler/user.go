package handler

import (
	"github.com/HsiaoCz/code-monster/templs/models"
	"github.com/HsiaoCz/code-monster/templs/views/uservs"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (us UserHandler) HandleUserShow(c *fiber.Ctx) error {
	user := models.User{
		Username: "bob",
		Email:    "bob@bob.com",
	}
	c.Set("Content-Type", "text/html")
	return uservs.Show(user).Render(c.Context(), c.Response().BodyWriter())
}
