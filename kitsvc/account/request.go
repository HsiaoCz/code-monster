package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Messsage string `json:"message"`
}

type GetUserResponse struct {
	Email string `json:"email"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(&response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (any, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (any, error) {
	idvars := mux.Vars(r)

	id, ok := idvars["id"]
	if !ok {
		return nil, errors.New("get id failed")
	}
	return GetUserRequest{ID: id}, nil
}
