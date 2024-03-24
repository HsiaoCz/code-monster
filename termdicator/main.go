package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/VictorLowther/btree"
	"github.com/gorilla/websocket"
)

const (
	wsendpoint = "wss://fstream.binance.com/stream?streams=btcusdt@depth"
)

func byBestBid(a, b *OrderbookEnty) bool {
	return a.Price > b.Price
}

func byBestAsk(a, b *OrderbookEnty) bool {
	return a.Price < b.Price
}

type OrderbookEnty struct {
	Price  float64
	Volume float64
}

type Orderbook struct {
	Asks *btree.Tree[*OrderbookEnty]
	Bids *btree.Tree[*OrderbookEnty]
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks: btree.New(byBestAsk),
		Bids: btree.New(byBestBid),
	}
}

func (o *Orderbook) handleDepthResponse(data BinanceOrderbookResult) {
	for _, ask := range data.Asks {
		price, _ := strconv.ParseFloat(ask[0], 64)
		volume, _ := strconv.ParseFloat(ask[1], 64)
		entry := &OrderbookEnty{
			Price:  price,
			Volume: volume,
		}
		fmt.Printf("%+v\n", entry)
	}
}

type BinanceDepthResponse struct {
	Stream string                 `json:"stream"`
	Data   BinanceOrderbookResult `json:"data"`
}

type BinanceOrderbookResult struct {
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsendpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	ob := NewOrderbook()

	var result BinanceDepthResponse
	for {
		if err := conn.ReadJSON(&result); err != nil {
			log.Fatal(err)
		}
		ob.handleDepthResponse(result.Data)
	}
}
