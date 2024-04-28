package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":9001"
	}

	srv := http.Server{
		Handler:      router,
		Addr:         port,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}
	slog.Info("the server is running", "port", port)
	log.Fatal(srv.ListenAndServe())
}
