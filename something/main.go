package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var (
		logger      = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{}))
		port        = os.Getenv("PORT")
		userHandler = NewUserHandler()
		router      = http.NewServeMux()
		srv         = http.Server{
			Handler:      router,
			Addr:         port,
			ReadTimeout:  time.Millisecond * 1500,
			WriteTimeout: time.Millisecond * 1500,
		}
	)
	slog.SetDefault(logger)

	router.HandleFunc("GET /user/show", TransferHandlerfunc(userHandler.HandleUserShow))

	slog.Info("the http server is running", "listen address", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
