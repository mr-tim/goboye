package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
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
		log.Printf("upgrade: %s", err)
		return
	}

	client := &Client{
		conn: ws,
		inbox: make(chan InboundMessage),
		outbox: make(chan OutboundMessage),
		emulator: &goboye.Emulator{},
	}
	go client.writeMessages()
	go client.readMessages()
	go client.handleMessages()

	client.emulator.LoadRomImage("/Users/tim/Desktop/goboye_research/Tetris (World).gb")
	client.refreshState()
}

type Client struct {
	conn     *websocket.Conn
	inbox    chan InboundMessage
	outbox   chan OutboundMessage
	emulator *goboye.Emulator
}

type OutboundMessage struct {
	Update UpdateMessage `json:"update"`
}

type UpdateMessage struct {
	Instructions []Instruction `json:"instructions"`
	Registers    map[string]int `json:"registers"`
}

type InboundMessage struct {
	Command CommandMessage `json:"command"`
}

type CommandMessage struct {
	Step StepCommand `json:"step"`
}

type StepCommand struct {

}

type Instruction struct {
	Address     int `json:"address"`
	Disassembly string `json:"disassembly"`
}

func (c *Client) readMessages() {
	for {
		var msg InboundMessage
		err := c.conn.ReadJSON(&msg)
		log.Printf("Read message from client: %#v", msg)
		if err != nil {
			log.Printf("Error reading json from ws: %s", err)
			continue
		}
		c.inbox <- msg
	}
}

func (c *Client) writeMessages() {
	for {
		select {
		case msg, ok := <- c.outbox:
			if !ok {
				//outbox closed - should hangup on client
				return
			}
			log.Printf("Sending message to client: %#v", msg)
			err := c.conn.WriteJSON(msg)
			if err != nil{
				log.Printf("err: %s\n", err)
			}
		}
	}
	defer c.conn.Close()
}

func (c *Client) handleMessages() {

}

func (c *Client) close() {
	close(c.outbox)
}

func (c *Client) refreshState() {
	log.Printf("Refreshing state...")
	disassembly := c.emulator.GetDisassembler()
	disassembly.SetPos(c.emulator.GetPC())

	instructions := make([]Instruction, 0)
	for i:=0; i<100; i+=1 {
		addr, o, payload := disassembly.GetNextInstruction()
		instructions = append(instructions, Instruction{
			Address: int(addr),
			Disassembly: o.DisassemblyWithArg(payload),
		})
	}

	msg := OutboundMessage{
		Update: UpdateMessage{
			Instructions: instructions,
			Registers: map[string]int{
				"AF": int(c.emulator.GetRegisterPair(cpu.RegisterPairAF)),
				"BC": int(c.emulator.GetRegisterPair(cpu.RegisterPairBC)),
				"DE": int(c.emulator.GetRegisterPair(cpu.RegisterPairDE)),
				"HL": int(c.emulator.GetRegisterPair(cpu.RegisterPairHL)),
				"SP": int(c.emulator.GetRegisterPair(cpu.RegisterPairSP)),
				"PC": int(c.emulator.GetRegisterPair(cpu.RegisterPairPC)),
			},
		},
	}

	c.outbox <- msg
}

func main() {
	flag.Parse()

	http.HandleFunc("/ws", serveWs)
	log.Printf("Listening on %s...", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
