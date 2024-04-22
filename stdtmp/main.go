package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-monster/stdtmp/api"
	"github.com/HsiaoCz/code-monster/stdtmp/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9001", "set the server address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(store.DBURI))
	if err != nil {
		slog.Error("connect the mongo db error", "err", err)
		return
	}

	var (
		userStore = store.NewMongoUserStore(client)
		store     = &store.Store{
			User: userStore,
		}
		userHander = api.NewUserHandler(store)

		router = http.NewServeMux()
	)

	router.HandleFunc("GET /user", userHander.HandleCreateUser)
	srv := http.Server{
		Handler:      router,
		Addr:         *listenAddr,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}
	slog.Info("the server is running", "port", *listenAddr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("the server error", "err", err)
		return
	}
}
