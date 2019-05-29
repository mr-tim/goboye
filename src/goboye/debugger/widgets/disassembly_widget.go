package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"goboye/cpu"
)

type DisassemblyWidget struct {
	da        *cpu.Disassembler
	currentPc uint16
}

func (d *DisassemblyWidget) Layout(g *gocui.Gui) error {
	_, maxY := g.Size()
	v, err := g.SetView("disassembly", 0, 0, 30, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Disassembly"
	v.Clear()
	// TODO: Handle scrolling in the disassembly view
	d.da.SetPos(0)
	for i := 0; i<maxY-1; i++ {
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

func NewDisassemblyWidget(disassembler *cpu.Disassembler, currentPc uint16) *DisassemblyWidget {
	return &DisassemblyWidget{
		da: disassembler,
		currentPc: currentPc,
	}
}