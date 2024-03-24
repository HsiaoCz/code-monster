package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-monster/pfetcher/client"
	"github.com/HsiaoCz/code-monster/pfetcher/protopkg"
)

func main() {
	var (
		json_address = flag.String("listenAddr", ":9021", "set listen address")
		grpc_address = flag.String("grpc_address", ":9022", "set the grpc address")
		svc          = NewLoggingService(&pricefetcher{})
		ctx          = context.Background()
	)
	flag.Parse()

	grpc_client, err := client.NewGRPCClient(":9022")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		resp, err := grpc_client.FetchPrice(ctx, &protopkg.FetchPriceRequest{Ticker: "GG"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go func() {
		if err := MakeGRPCServerAndRun(*grpc_address, svc); err != nil {
			log.Fatal(err)
		}
	}()

	jsonApi := NewJsonApiServer(*json_address, svc)
	jsonApi.Run()

	select {}
}
