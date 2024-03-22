package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	netx Pricefetcher
}

func NewLoggingService(next Pricefetcher) Pricefetcher {
	return &loggingService{
		netx: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value(Key("requestID")),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())
	return s.netx.FetchPrice(ctx, ticker)
}
