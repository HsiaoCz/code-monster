package store

import (
	"context"

	"github.com/HsiaoCz/code-monster/chit/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type UserMongoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserMongoStore(client *mongo.Client, dbname string, coll string) *UserMongoStore {
	return &UserMongoStore{
		client: client,
		coll:   client.Database(dbname).Collection(coll),
	}
}

func (u *UserMongoStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	resp := u.coll.FindOne(ctx, bson.M{"_id": id})
	if err := resp.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
