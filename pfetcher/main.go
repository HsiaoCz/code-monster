package main

import "flag"

func main() {
	listenAddr := flag.String("listenAddr", ":9001", "set listen address")
	flag.Parse()
	svc := NewLoggingService(&pricefetcher{})

	jsonApi := NewJsonApiServer(*listenAddr, svc)
	jsonApi.Run()
}
