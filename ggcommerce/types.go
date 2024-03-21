package main

type CreateProductRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

type Product struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
