package main

import (
	"log/slog"
	"net/http"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type APIServer struct {
	listenAddr  string
	userHandler *UserHandler
}

func NewAPIServer(listenAddr string, userHandler *UserHandler) *APIServer {
	return &APIServer{
		listenAddr:  listenAddr,
		userHandler: userHandler,
	}
}

func (a *APIServer) Start() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /users", a.makeHTTPHandler(a.userHandler.HandleGetUsers))
	router.HandleFunc("POST /user", a.makeHTTPHandler(a.userHandler.HandleCreateUser))
	router.HandleFunc("PUT /user/{id}", a.makeHTTPHandler(a.userHandler.HandleUpdateUser))
	router.HandleFunc("DELETE /user/{id}", a.makeHTTPHandler(a.userHandler.HandleDeleteUser))

	server := http.Server{
		Handler:      router,
		Addr:         a.listenAddr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	slog.Info("the api server is running", "the address", a.listenAddr)
	return server.ListenAndServe()
}

func (a *APIServer) makeHTTPHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("the api server error", "err", err)
			if e, ok := err.(APIError); ok {
				WriteJSON(w, http.StatusInternalServerError, e)
			}
		}
	}
}
