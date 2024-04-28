package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME = "gotth"
const COLL = "users"

type UserStorer interface {
	CreateUser(context.Context, *User) (*User, error)
	GetUser(context.Context, string) (*User, error)
}

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

func (m *MongoUserStoer) CreateUser(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (m *MongoUserStoer) GetUser(ctx context.Context, email string) (*User, error) {
	return nil, nil
}
