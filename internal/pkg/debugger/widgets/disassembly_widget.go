package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
	"log"
)

const DisassemblyViewName = "disassembly"

type DisassemblyWidget struct {
	emulator *goboye.Emulator
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
	da := d.emulator.GetDisassembler()
	currentPC := d.emulator.GetPC()
	h := uint16(maxY - 2)
	da.SetPos(currentPC - (currentPC % h))
	for i := 0; i < maxY-1; i++ {
		addr, o := da.GetNextInstruction()
		pointer := " "
		if addr == currentPC {
			pointer = ">"
		}
		fmt.Fprintf(v, " %s 0x%04x %s\n", pointer, addr, o.Disassembly())
	}

	return nil
}

func NewDisassemblyWidget(emulator *goboye.Emulator) *DisassemblyWidget {
	return &DisassemblyWidget{
		emulator: emulator,
	}
}
