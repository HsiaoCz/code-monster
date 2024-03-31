package main

import (
	"log"

	"github.com/anthdm/hollywood/actor"
)

type Manager struct{}

func (m *Manager) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
	case actor.Stopped:
	}
}

func NewManager() actor.Producer {
	return func() actor.Receiver {
		return &Manager{}
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(NewManager(), "manager")
}
