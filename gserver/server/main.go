package main

import (
	"log"

	"github.com/anthdm/hollywood/actor"
)

type GameServer struct{}

func NewGameServer() actor.Receiver {
	return &GameServer{}
}

func (s *GameServer) Receive(c *actor.Context) {}

func main() {
	e, err := actor.NewEngine(actor.EngineConfig{})
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(NewGameServer, "server")
}
