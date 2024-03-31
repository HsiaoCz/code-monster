package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"userID"`
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", HandleGetUserByID)
	http.ListenAndServe(":3001", router)
}

func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, &User{UserID: id})

}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
