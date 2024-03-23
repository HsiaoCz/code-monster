package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40001/ws"

type Login struct {
	ClientID int    `json:"clientID"`
	Username string `json:"username"`
}

type GameClient struct {
	conn     *websocket.Conn
	clientID int
	username string
}

func (g *GameClient) login() error {
	return g.conn.WriteJSON(&Login{
		ClientID: g.clientID,
		Username: g.username,
	})
}

func NewGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientID: rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func main() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := NewGameClient(conn, "james")
	if err := c.login(); err != nil {
		log.Fatal(err)
	}
}

func UserLogin(conn *websocket.Conn, data Login) error {
	return conn.WriteJSON(data)
}
