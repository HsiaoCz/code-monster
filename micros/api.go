package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type APIServer struct {
	listenAddr string
	svc        Service
}

func NewAPIServer(svc Service, listenAddr string) *APIServer {
	return &APIServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleGetFact(w http.ResponseWriter, r *http.Request) error {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		return WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{
			"error": err.Error(),
		})
	}
	return WriteJSON(w, http.StatusOK, fact)
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func TransferHTTPHandlefunc(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}

func (s *APIServer) Start() error {
	http.HandleFunc("/catfact", TransferHTTPHandlefunc(s.handleGetFact))
	return http.ListenAndServe(s.listenAddr, nil)
}
