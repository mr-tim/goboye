package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"log"
)

const DisassemblyViewName = "disassembly"

type DisassemblyWidget struct {
	da        *cpu.Disassembler
	currentPc uint16
}

func (d *DisassemblyWidget) Layout(g *gocui.Gui) error {
	log.Printf("Running DisassemblyWidget.Layout")
	_, maxY := g.Size()
	v, err := g.SetView(DisassemblyViewName, 0, 0, 30, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Disassembly"
	v.Clear()
	// TODO: Handle scrolling in the disassembly view
	d.da.SetPos(d.currentPc)
	for i := 0; i < maxY-1; i++ {
		addr, o, payload := d.da.GetNextInstruction()
		pointer := " "
		if addr == d.currentPc {
			pointer = ">"
		}
		fmt.Fprintf(v, " %s 0x%04x %s\n", pointer, addr, o.DisassemblyWithArg(payload))
	}

	return nil
}

func (d *DisassemblyWidget) SetPc(newPc uint16) {
	d.currentPc = newPc
}

func (d *DisassemblyWidget) Increment(g *gocui.Gui, v *gocui.View) error {
	log.Printf("Incrementing...")
	pos, _, _ := d.da.GetNextInstruction()
	d.SetPc(pos)
	log.Printf("Pc = %#v", pos)
	return nil
}

func NewDisassemblyWidget(disassembler *cpu.Disassembler, currentPc uint16) *DisassemblyWidget {
	return &DisassemblyWidget{
		da:        disassembler,
		currentPc: currentPc,
	}
}
