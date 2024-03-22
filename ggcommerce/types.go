package main

import "fmt"

const minProductNameLen = 3

type CreateProductRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

type Product struct {
	SKU  string `bson:"sku"  json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}

func ValidateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product is too short")
	}
	return nil
}
