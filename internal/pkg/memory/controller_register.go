package memory

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/goboye/button"
	"github.com/mr-tim/goboye/internal/pkg/utils"
)

type controllerRegister struct {
	selectRow1 bool
	selectRow2 bool
	buttonDown map[button.Button]bool
}

var row1 = []button.Button{button.Right, button.Left, button.Up, button.Down}
var row2 = []button.Button{button.A, button.B, button.Select, button.Start}

func NewControllerRegister() controllerRegister {
	return controllerRegister{
		buttonDown: make(map[button.Button]bool),
	}
}

func (r *controllerRegister) Read() byte {
	// output is pinned high by default
	selection := uint8(0x00)
	if r.selectRow1 {
		selection |= 0x01 << 4 | r.row(row1)
	}
	if r.selectRow2 {
		selection |= 0x01 << 5 | r.row(row2)
	}
	return selection
}

func (r *controllerRegister) row(bs []button.Button) uint8 {
	result := uint8(0)
	for idx, b := range bs {
		if v, exists := r.buttonDown[b]; !exists || !v {
			result |= 0x01 << idx
		}
	}
	return result
}

func (r *controllerRegister) Write(value byte) {
	r.selectRow1 = utils.IsBitSet(value, 5)
	r.selectRow2 = utils.IsBitSet(value, 4)
}

func (r *controllerRegister) SetButtonState(button button.Button, isDown bool) {
	fmt.Printf("Set button state: %s is down: %s => %02X\n", button, isDown, r.Read())
	r.buttonDown[button] = isDown
}
