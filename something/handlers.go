package main

import "net/http"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandleUserShow(w http.ResponseWriter, r *http.Request) error {
	return APIError{
		Status: http.StatusInternalServerError,
		Msg:    "request refuesed!",
	}
}
