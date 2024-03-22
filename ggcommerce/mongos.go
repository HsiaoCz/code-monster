package main

import "go.mongodb.org/mongo-driver/mongo"

type MongoProductStore struct {
	client *mongo.Client
}

func NewMongoProductStore(c *mongo.Client) *MongoProductStore {
	return &MongoProductStore{
		client: c,
	}
}

func (s *MongoProductStore) Insert(p *Product) error {
	return nil
}

func (s *MongoProductStore) GetProductByID(id string) (*Product, error) {
	return nil, nil
}
