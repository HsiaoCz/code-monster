package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/HsiaoCz/code-monster/pfetcher/types"
)

type Key string

type JSONAPIServer struct {
	listenAddr string
	svc        Pricefetcher
}

func NewJsonApiServer(listenAddr string, svc Pricefetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/price", transHandler(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, &types.PriceResponse{
		Ticker: ticker,
		Price:  price,
	})
}

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

func transHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, Key("requestID"), rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt))
		if err := h(ctx, w, r); err != nil {
			slog.Error("error", "err", err)
			WriteJson(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}
