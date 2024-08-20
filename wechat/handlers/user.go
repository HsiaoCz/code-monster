package handlers

import "net/http"

type UserHandler struct{}

func (u *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
