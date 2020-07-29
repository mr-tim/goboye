package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
	"image/png"
	"log"
	"net/http"
	"strings"
	"sync"
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
	log.Printf("Starting client...\n")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error from upgrade: %s", err)
		return
	}

	client := &Client{
		conn:      ws,
		inbox:     make(chan InboundMessage),
		outbox:    make(chan OutboundMessage),
		emulator:  goboye.NewEmulator(),
		closeOnce: &(sync.Once{}),
	}
	go client.writeMessages()
	go client.readMessages()
	go client.handleMessages()

	client.emulator.LoadRomImage("/Users/tim/Desktop/goboye_research/Tetris (World).gb")
	client.refreshState()
}

type Client struct {
	conn      *websocket.Conn
	inbox     chan InboundMessage
	outbox    chan OutboundMessage
	emulator  *goboye.Emulator
	closeOnce *sync.Once
}

type OutboundMessage struct {
	Update UpdateMessage `json:"update"`
}

type UpdateMessage struct {
	Instructions  []Instruction  `json:"instructions"`
	Registers     map[string]int `json:"registers"`
	MemoryUpdates []MemoryUpdate `json:"memory_updates"`
	Breakpoints   []uint16       `json:"breakpoints"`
	DebugImage    string         `json:"debug_image"`
}

type MemoryUpdate struct {
	Start        uint16 `json:"start"`
	Length       uint16 `json:"length"`
	MemoryBase64 string `json:"memory_base64"`
}

type InboundMessage struct {
	Command CommandMessage `json:"command"`
}

type CommandMessage struct {
	Step       *StepCommand       `json:"step"`
	Breakpoint *BreakpointCommand `json:"breakpoint"`
	Continue   *ContinueCommand   `json:"continue"`
}

type StepCommand struct {
}

type BreakpointCommand struct {
	Address uint16 `json:"address"`
	Break   bool   `json:"break"`
}

type ContinueCommand struct {
}

type Instruction struct {
	Address     int    `json:"address"`
	Disassembly string `json:"disassembly"`
}

func (c *Client) readMessages() {
	for {
		var msg InboundMessage
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(err, 1001) {
				log.Printf("Client socket closed: %s", err)
			} else {
				log.Printf("Error reading json from ws: %s", err)
			}
			c.close()
			return
		}
		log.Printf("Read message from client: %#v", msg)
		c.inbox <- msg
	}
}

func (c *Client) writeMessages() {
	for {
		select {
		case msg, ok := <-c.outbox:
			if !ok {
				//outbox closed - should hangup on client
				log.Print("Outbox closed.\n")
				return
			}
			s := fmt.Sprintf("Sending message to client: %#v", msg)
			log.Printf(s[:200])
			err := c.conn.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing json to ws: %s\n", err)
				c.close()
				return
			}
		}
	}
}

func (c *Client) handleMessages() {
	for {
		select {
		case msg, ok := <-c.inbox:
			if !ok {
				log.Printf("Inbox closed.\n")
				return
			}

			cmd := msg.Command
			if cmd.Step != nil {
				log.Print("Received step command")
				c.emulator.Step()
				c.refreshState()
			} else if cmd.Breakpoint != nil {
				log.Print("Received breakpoint command")
				if cmd.Breakpoint.Break {
					c.emulator.AddBreakpoint(cmd.Breakpoint.Address)
				} else {
					c.emulator.RemoveBreakpoint(cmd.Breakpoint.Address)
				}
				c.refreshState()
			} else if cmd.Continue != nil {
				log.Print("Received continue command")
				c.emulator.ContinueDebugging()
				c.refreshState()
			}
		}
	}
	fmt.Printf("Finished handling messages\n")
}

func (c *Client) close() {
	c.closeOnce.Do(func() {
		close(c.outbox)
		close(c.inbox)
		c.conn.Close()
		fmt.Printf("Closed outbox and inbox\n")
	})
}

func (c *Client) refreshState() {
	log.Printf("Refreshing state...")
	disassembly := c.emulator.GetDisassembler()
	disassembly.SetPos(c.emulator.GetPC())

	instructions := make([]Instruction, 0)
	for i := 0; i < 100; i += 1 {
		addr, o, payload := disassembly.GetNextInstruction()
		instructions = append(instructions, Instruction{
			Address:     int(addr),
			Disassembly: o.DisassemblyWithArg(payload),
		})
	}

	debugRenderImage := c.emulator.DebugRender()
	b := new(bytes.Buffer)
	err := png.Encode(b, debugRenderImage)
	if err != nil {
		panic(err)
	}
	base64debugImage := base64.StdEncoding.EncodeToString(b.Bytes())

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
			MemoryUpdates: []MemoryUpdate{
				{
					MemoryBase64: c.emulator.MemoryBase64(),
					Start:        0,
					Length:       0xFFFF,
				},
			},
			Breakpoints: c.emulator.GetBreakpoints(),
			DebugImage:  base64debugImage,
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
