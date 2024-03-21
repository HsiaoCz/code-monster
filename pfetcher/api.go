package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type JSONAPIServer struct {
	svc Pricefetcher
}

func NewJsonApiServer(svc Pricefetcher) *JSONAPIServer {
	return &JSONAPIServer{
		svc: svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/price", transHandler(s.handleFetchPrice))
	http.ListenAndServe(":9021", nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, &PriceResponse{
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
		if err := h(ctx, w, r); err != nil {
			slog.Error("error", "err", err)
			WriteJson(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}
