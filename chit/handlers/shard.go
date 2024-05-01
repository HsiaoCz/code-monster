package handlers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

func TransferHTTPHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("server error", "err", err, "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr)
			if e, ok := err.(APIError); ok {
				if err := WriteJSON(w, e.Status, &e); err != nil {
					log.Fatalf("json encoding to the response error %v\n", err)
				}
			} else {
				apiErr := APIError{
					Status: http.StatusInternalServerError,
					Msg:    err.Error(),
				}
				if err := WriteJSON(w, apiErr.Status, &apiErr); err != nil {
					log.Fatalf("json encoding to the response error %v\n", err)
				}
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
