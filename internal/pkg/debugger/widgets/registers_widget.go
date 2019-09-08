package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type RegistersWidget struct {
}

func (w *RegistersWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("registers", maxX-30, 0, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Registers"
	v.Clear()

	fmt.Fprintf(v, "AB = 0x0000")

	return nil
}

func NewRegistersWidget() *RegistersWidget {
	return &RegistersWidget{}
}
