package main

import (
	"database/sql"
	"flag"
	"log/slog"

	"github.com/HsiaoCz/code-monster/cpba/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9001", "set the listen address")
	flag.Parse()
	server := api.NewAPIServer(*listenAddr, &sql.DB{})
	if err := server.Run(); err != nil {
		slog.Error("server run err:", "err", err)
		return
	}
}
