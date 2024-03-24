package main

import (
	"context"
	"net"

	"github.com/HsiaoCz/code-monster/pfetcher/protopkg"
	"google.golang.org/grpc"
)

type GRPCPriceFetcher struct {
	svc Pricefetcher
	protopkg.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcher(svc Pricefetcher) *GRPCPriceFetcher {
	return &GRPCPriceFetcher{
		svc: svc,
	}
}

func (pf *GRPCPriceFetcher) FetchPrice(ctx context.Context, in *protopkg.FetchPriceRequest) (*protopkg.FetchPriceResponse, error) {
	price, err := pf.svc.FetchPrice(ctx, in.GetTicker())
	if err != nil {
		return nil, err
	}
	return &protopkg.FetchPriceResponse{Ticker: in.GetTicker(), Price: price}, err
}

func MakeGRPCServerAndRun(listenAddr string, svc Pricefetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	protopkg.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	return server.Serve(ln)
}
