package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongourl = ""
const dbname = ""
const collname = ""

// 6 interfaces
// Storage
// Storer
// io.Reader
// io.Writer
type Storer interface {
	GetUser(context.Context) ([]*User, error)
	CreateUser(context.Context, *CreateParams) (*User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, bson.M, *UpdateUserParams) (*User, error)
}

type MongoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongourl))
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		client: client,
		coll:   client.Database(dbname).Collection(collname),
	}, nil
}

func (m *MongoStore) GetUser(ctx context.Context) ([]*User, error) {
	resp, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*User
	if err := resp.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *MongoStore) CreateUser(ctx context.Context, param CreateParams) (*User, error) {
	user := User{
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
	}
	result, err := m.coll.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}
	user.UserID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (m *MongoStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := m.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("not delete any user")
	}
	return nil
}

func (m *MongoStore) UpdateUser(ctx context.Context, filter bson.M, params *UpdateUserParams) (*User, error) {
	update := bson.D{
		{
			Key: "$set", Value: params.ToBSON(),
		},
	}
	_, err := m.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	var user User
	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
