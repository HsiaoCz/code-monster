package api

import (
	"database/sql"

	"github.com/HsiaoCz/code-monster/cpba/service/user"
	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	listenAddr string
	db         *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()

	v1 := app.Group("/api/v1")

	userHandler:=user.NewHandler()
	userHandler.RegisterRouters(v1)

	return app.Listen(s.listenAddr)
}
