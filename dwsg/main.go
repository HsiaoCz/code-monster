package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type VisitRequest struct {
	links []string
}

type Manager struct{}

func (m *Manager) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case VisitRequest:
		m.handleVisitRequest(msg)
	case actor.Started:
		slog.Info("manager started")
	case actor.Stopped:
	}
}

func (m *Manager) handleVisitRequest(msg VisitRequest) error {
	for _, link := range msg.links {
		slog.Info("visiting url", "url", link)
	}
	return nil
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
	time.Sleep(time.Second * 10)
}
