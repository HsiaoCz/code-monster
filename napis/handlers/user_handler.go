package handlers

import "github.com/HsiaoCz/code-monster/napis/store"

type UserHandler struct {
	store store.Store
}

func NewUserHandler(store store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}
