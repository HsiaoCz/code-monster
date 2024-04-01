package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id" json:"userID"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
}

type CreateParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserParams struct {
	Username string `json:"username"`
}

func (u *UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(u.Username) > 0 {
		m["username"] = u.Username
	}
	return m
}

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a APIError) Error() string {
	return a.Message
}
