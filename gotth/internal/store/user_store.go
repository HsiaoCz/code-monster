package store

import "go.mongodb.org/mongo-driver/mongo"

const DBNAME = ""
const COLL = ""

type UserStorer interface{}

type MongoUserStoer struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStoer {
	return &MongoUserStoer{
		client: client,
		coll:   client.Database(DBNAME).Collection(COLL),
	}
}
