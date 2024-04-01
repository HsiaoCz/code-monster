package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

type Server struct {
	peers map[net.Conn]bool
	ln    net.Listener
}

func NewServer() *Server {
	return &Server{
		peers: make(map[net.Conn]bool),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":30111")
	if err != nil {
		return err
	}
	s.ln = ln

	slog.Info("tcp server started", "port", ":3000")

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("err==>", err.Error())
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	s.peers[conn] = true
	fmt.Println("handle the connention", conn.RemoteAddr())
}

func main() {
	server := NewServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
