package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func TransferHTTPHandler(fn Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			slog.Error("handler error", "err", err)
			if e, ok := err.(APIError); ok {
				WriteJSON(w, e.Status, e)
			} else {
				WriteJSON(w, http.StatusInternalServerError, &APIError{
					Type:    "error",
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
