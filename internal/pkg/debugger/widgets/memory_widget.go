package widgets

import (
	"github.com/jroimartin/gocui"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
)

type MemoryWidget struct {
	emulator *goboye.Emulator
}

func (w *MemoryWidget) Layout(g *gocui.Gui) error {

	return nil
}
