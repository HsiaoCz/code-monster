package main

import (
	"context"
	"flag"
	"log/slog"
	"net"

	"github.com/HsiaoCz/code-monster/lost/internal/user/pb"
	"github.com/HsiaoCz/code-monster/lost/internal/user/service"
	"github.com/HsiaoCz/code-monster/lost/internal/user/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(store.DBURL))
	if err != nil {
		slog.Error("connect to the mongodb error", "err", err)
		return
	}

	listenAddr := flag.String("listenAddr", ":9001", "set the listen address of the grpc user serveice")
	flag.Parse()

	var (
		userStore   = store.NewMongoUserStore(client)
		store       = &store.Store{UserStore: userStore}
		userService = service.NewUserService(store)
		server      = grpc.NewServer()
	)

	listen, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		slog.Error("net listen error", "err", err)
		return
	}
	pb.RegisterLostServer(server, userService)
	if err := server.Serve(listen); err != nil {
		slog.Error("server serve error", "err", err)
		return
	}
}
