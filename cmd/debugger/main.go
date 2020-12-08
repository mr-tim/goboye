package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jroimartin/gocui"

	"github.com/mr-tim/goboye/internal/pkg/debugger/widgets"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
)

func main() {
	f, err := os.OpenFile("debugger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	if len(os.Args) != 2 {
		fmt.Printf("Please specify a rom to run.\n")
		os.Exit(1)
	}

	// set up the emulator
	emulator := &goboye.Emulator{}
	emulator.LoadRomImage(os.Args[1])

	// set up ui
	log.Printf("Creating ui...")
	dw := widgets.NewDisassemblyWidget(emulator)
	rw := widgets.NewRegistersWidget(emulator)
	bm := widgets.NewBreakpointModal(emulator)

	g.SetManager(dw, rw)

	err = g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		emulator.Step()
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to set key binding for increment: %s", err)
	}

	err = g.SetKeybinding("", 'b', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		bm.Toggle()
		return nil
	})

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
