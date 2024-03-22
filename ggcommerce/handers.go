package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	store ProductStorer
}

func NewProductHandler(store ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: store,
	}
}

func (p *ProductHandler) HandleGetProduct(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(&Product{SKU: "SHOE-1111"})
}

func (p *ProductHandler) HandlePostProduct(c *fiber.Ctx) error {
	productReq := &CreateProductRequest{}
	if err := c.BodyParser(productReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if err := ValidateCreateProductRequest(productReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	product := &Product{
		SKU:  productReq.SKU,
		Name: productReq.Name,
		Slug: "SHOE",
	}

	if err := p.store.Insert(product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "post product successed!",
	})
}
