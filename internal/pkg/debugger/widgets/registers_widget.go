package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
)

type RegistersWidget struct {
	emulator *goboye.Emulator
}

const tick = "✓"
const cross = "×"

func (w *RegistersWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("registers", maxX-30, 0, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Registers"
	v.Clear()

	w.outputRegister(cpu.RegisterPairAF, v)
	w.outputRegister(cpu.RegisterPairBC, v)
	w.outputRegister(cpu.RegisterPairDE, v)
	w.outputRegister(cpu.RegisterPairHL, v)
	w.outputRegister(cpu.RegisterPairSP, v)
	w.outputRegister(cpu.RegisterPairPC, v)
	fmt.Fprintf(v, "\n")
	fmt.Fprintf(v, " z = %s\n", w.getOpFlag(cpu.FlagZ))
	fmt.Fprintf(v, " n = %s\n", w.getOpFlag(cpu.FlagN))
	fmt.Fprintf(v, " h = %s\n", w.getOpFlag(cpu.FlagH))
	fmt.Fprintf(v, " c = %s\n", w.getOpFlag(cpu.FlagC))
	return nil
}

func (w *RegistersWidget) outputRegister(rp cpu.RegisterPair, v *gocui.View) {
	pairName := rp.String()[len(rp.String())-2:]
	fmt.Fprintf(v, " %s = %04x\n", pairName, w.emulator.GetRegisterPair(rp))
}

func (w *RegistersWidget) getOpFlag(flagName cpu.OpResultFlag) string {
	if w.emulator.GetFlagValue(flagName) {
		return tick
	} else {
		return cross
	}
}

func NewRegistersWidget(emulator *goboye.Emulator) *RegistersWidget {
	return &RegistersWidget{
		emulator: emulator,
	}
}
