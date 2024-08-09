package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-monster/traceman/db"
	"github.com/HsiaoCz/code-monster/traceman/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	var (
		port        = os.Getenv("PORT")
		userHandler = &handlers.UserHandlers{}
		router      = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandler.HandleUserCreate))
	}

	http.ListenAndServe(port, router)
}
