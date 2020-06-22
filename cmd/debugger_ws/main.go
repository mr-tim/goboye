package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return strings.HasPrefix(r.RemoteAddr, "127.0.0.1:") || strings.HasPrefix(r.RemoteAddr, "localhost:")
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	client := &Client{conn: ws}
	go client.writeMessages()
	go client.readMessages()
}

type Client struct {
	conn *websocket.Conn
}

type Message struct {
	Update UpdateMessage `json:"update"`
}

type UpdateMessage struct {
	Instructions []Instruction `json:"instructions"`
	Registers    map[string]int `json:"registers"`
}

type Instruction struct {
	Address     int `json:"address"`
	Disassembly string `json:"disassembly"`
}

func (c *Client) readMessages() {

}

func (c *Client) writeMessages() {
	msg := Message{
		Update: UpdateMessage{
			Instructions: []Instruction{
				{Address: 0, Disassembly: "LD SP,0xFFFE"},
				{Address: 3, Disassembly: "XOR A"},
				{Address: 4, Disassembly: "LD HL,0"},
			},
			Registers: map[string]int{
				"AF": 0,
				"BC": 0,
				"DE": 0,
				"HL": 0,
				"SP": 0,
				"PC": 0,
			},
		},
	}
	err := c.conn.WriteJSON(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	defer c.conn.Close()
}

func main() {
	flag.Parse()

	http.HandleFunc("/ws", serveWs)
	fmt.Printf("Listening on %s...\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
