package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/code-monster/traceman/dao"
	"github.com/HsiaoCz/code-monster/traceman/db"
	"github.com/HsiaoCz/code-monster/traceman/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	var (
		port        = os.Getenv("PORT")
		userCase    = dao.UserCaseInit(db.Get())
		userHandler = handlers.UserHandlersInit(userCase)
		router      = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandler.HandleUserCreate))
		router.HandleFunc("GET /user", handlers.TransferHandlerfunc(userHandler.GetUserByID))
	}
	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("http server is running")
	http.ListenAndServe(port, router)
}
