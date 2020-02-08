package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jroimartin/gocui"

	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/debugger/widgets"
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

func main() {
	f, err := os.OpenFile("debugger.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Printf(strings.Repeat("=", 80))
	log.Printf("Starting debugger...")

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	buf := make([]byte, memory.MEM_SIZE)
	mm := memory.NewMemoryMapWithBytes(buf)

	if len(os.Args) != 2 {
		fmt.Printf("Please specify a rom to run.\n")
		os.Exit(1)
	}

	// set up the emulator
	log.Printf("Loading rom: %s", os.Args[1])
	e := mm.LoadRomImage(os.Args[1])
	if e != nil {
		panic(e)
	}

	p := cpu.NewProcessor(mm)
	pc := p.GetRegisterPair(cpu.RegisterPairPC)

	// set up ui
	log.Printf("Creating ui...")
	da := cpu.NewDisassembler(mm)
	dw := widgets.NewDisassemblyWidget(da, pc)
	rw := widgets.NewRegistersWidget()
	g.SetManager(dw, rw)

	err = g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, dw.Increment)
	if err != nil {
		log.Fatalf("Failed to set key binding for increment: %s", err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Printf("Received ctrl-c")
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
