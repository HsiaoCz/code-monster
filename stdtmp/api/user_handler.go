package api

import (
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/code-monster/stdtmp/store"
	"github.com/HsiaoCz/code-monster/stdtmp/views/usersv"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(store *store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := u.store.User.GetUser(r.Context(), "gg@gg.com")
	if err != nil {
		slog.Error("get user from the store error", "err", err)
		return
	}
	usersv.Show(*user).Render(r.Context(), w)
}
