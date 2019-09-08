package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"

	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/debugger/widgets"
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

func main() {
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

	e := mm.LoadRomImage(os.Args[1])
	if e != nil {
		panic(e)
	}

	p := cpu.NewProcessor(mm)
	pc := p.GetRegisterPair(cpu.RegisterPairPC)

	da := cpu.NewDisassembler(mm)
	dw := widgets.NewDisassemblyWidget(da, pc)
	rw := widgets.NewRegistersWidget()
	g.SetManager(dw, rw)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
