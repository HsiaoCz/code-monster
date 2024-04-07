package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-monster/napis/store"
)

type UserHandler struct {
	store store.Store
}

func NewUserHandler(store store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	
	return nil
}
