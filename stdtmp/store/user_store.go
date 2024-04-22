package store

import (
	"context"

	"github.com/HsiaoCz/code-monster/stdtmp/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	GetUser(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(DBCOLL),
	}
}

func (m *MongoUserStore) GetUser(ctx context.Context, email string) (*types.User, error) {
	return &types.User{
		Username: "gg",
		Email:    email,
	}, nil
}
