package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/code-monster/pfetcher/client"
)

func main() {
	client := client.NewClient("http://localhost:9001/price")
	responsePrice, err := client.FetchPrice(context.Background(), "GG")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response %v\n", responsePrice)
}
