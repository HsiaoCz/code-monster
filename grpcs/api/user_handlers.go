package api

import "github.com/gofiber/fiber/v2"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return nil
}
