package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	admin := app.Group("/admin")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongo"))
	if err != nil {
		log.Fatal(err)
	}
	productStore := NewMongoProductStore(client)
	productHandler := NewProductHandler(productStore)

	admin.Get("/product", productHandler.HandleGetProduct)
	admin.Post("/product", productHandler.HandlePostProduct)
	app.Listen(":4001")
}
