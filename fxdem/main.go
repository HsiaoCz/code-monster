package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewWebServer,
			NewServeMux,
			NewEchoHandler,
			NewDB,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewWebServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{
		Addr:    ":3001",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			slog.Info("the http server is running", "port", srv.Addr)
			go func() {
				err = srv.Serve(ln)
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	db *DB
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(db *DB) *EchoHandler {
	return &EchoHandler{
		db: db,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (e *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(e.db.GetUserByID(10)))
}

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}

type DB struct{}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) GetUserByID(id int) string {
	return fmt.Sprintf("anth-%d", id)
}
