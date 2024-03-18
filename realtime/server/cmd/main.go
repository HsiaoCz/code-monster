package main

import (
	"log"

	"github.com/HsiaoCz/code-monster/realtime/server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
