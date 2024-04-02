package store

import "go.mongodb.org/mongo-driver/mongo"

type UserStorer interface{}

type MongoUserStoer struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore() {}
