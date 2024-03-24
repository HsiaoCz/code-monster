package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLogginService(svc)

	fact, err := svc.GetCatFact(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("fact=%v\n", fact)
}
