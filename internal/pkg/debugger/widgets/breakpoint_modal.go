package widgets

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
)

const BreakpointModalViewName = "breakpoint_modal"

type BreakpointModal struct {
	emulator *goboye.Emulator
	visible  bool
}

var _ gocui.Manager = &BreakpointModal{}

func (bm *BreakpointModal) Layout(g *gocui.Gui) error {
	if bm.visible {
		maxX, maxY := g.Size()
		v, err := g.SetView(BreakpointModalViewName, 0, 0, maxX-1, maxY-1)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Add Breakpoint"
		v.Clear()
		fmt.Fprintf(v, "Enter a breakpoint")
		g.SetViewOnTop(BreakpointModalViewName)
	} else {
		g.SetViewOnBottom(BreakpointModalViewName)
	}

	return nil
}

func (bm *BreakpointModal) Toggle() {
	bm.visible = !bm.visible
}

func NewBreakpointModal(emulator *goboye.Emulator) *BreakpointModal {
	return &BreakpointModal{
		emulator: emulator,
		visible:  false,
	}
}
