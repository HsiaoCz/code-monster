package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next Pricefetcher
}

func NewMetricService(next Pricefetcher) Pricefetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus")
	return s.next.FetchPrice(ctx, ticker)
}
