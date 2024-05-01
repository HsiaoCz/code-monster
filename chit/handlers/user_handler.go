package handlers

import (
	"net/http"

	"github.com/HsiaoCz/code-monster/chit/store"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(store *store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	user, err := u.store.U.GetUserByID(r.Context(), id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, user)
}
