package user

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouters(r fiber.Router) {
	userGroup := r.Group("/user")
	{
		userGroup.Post("/register", h.handleRegister)
		userGroup.Post("/login", h.handleLogin)
	}
}

func (h *Handler) handleRegister(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) handleLogin(c *fiber.Ctx) error {
	return nil
}
