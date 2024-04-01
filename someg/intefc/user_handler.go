package main

import "net/http"

type UserHandler struct {
	store Storer
}

func NewUserHandler(store Storer) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleGetUsers(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
