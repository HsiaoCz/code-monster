package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var cur CreateUserReq
	if err := c.BodyParser(&cur); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
	}
	res, err := h.Service.CreateUser(c.Context(), &cur)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  http.StatusInternalServerError,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code": http.StatusOK,
		"data": res,
	})
}

// update user

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	return nil
}
