package main

type ProductStorer interface {
	Insert(*Product) error
	GetProductByID(string) (*Product, error)
}
