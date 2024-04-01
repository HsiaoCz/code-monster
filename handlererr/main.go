package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"userID"`
}

type apiError struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func (e apiError) Error() string {
	return e.Msg
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", makeHandler(handleGetUserByID))
	router.HandleFunc("GET /user", makeHandler(handleGetUserList))
	http.ListenAndServe(":3001", router)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return apiError{
			Status: http.StatusBadRequest,
			Msg:    "check out the id",
		}
	}
	return writeJSON(w, http.StatusOK, &User{UserID: id})
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func handleGetUserList(w http.ResponseWriter, r *http.Request) error {
	users, err := getUsers()
	if err != nil {
		slog.Error("the get users error", "err", err)
		return apiError{
			Status: http.StatusInternalServerError,
			Msg:    "the server error",
		}
	}
	return writeJSON(w, http.StatusOK, users)
}

func getUsers() ([]User, error) {
	return []User{}, nil
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func makeHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				slog.Error("api error", "err", e.Error(), "status", e.Status)
				writeJSON(w, e.Status, e.Msg)
			}
		}
	}
}
