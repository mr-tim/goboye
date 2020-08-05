package display

import "github.com/mr-tim/goboye/internal/pkg/memory"

type LcdcMode byte

const (
	EnableCPUAccessToDisplayRAM LcdcMode = 0
	VerticalBlank               LcdcMode = 1
	SearchingOAMRAM             LcdcMode = 2
	TransferringDataToLCDDriver LcdcMode = 3
)

type LcdInterruptSelector byte

const (
	Mode00   LcdInterruptSelector = 3
	Mode01   LcdInterruptSelector = 4
	Mode10   LcdInterruptSelector = 5
	LycMatch LcdInterruptSelector = 6
)

type StatFlags struct {
	r memory.RwRegister
}

func (f *StatFlags) GetMode() LcdcMode {
	return LcdcMode(f.r.GetByte() & 0x03)
}

func (f *StatFlags) SetMode(mode LcdcMode) {
	updated := (f.r.GetByte() & 0b11111100) | byte(mode)
	f.r.SetValue(updated)
}

func (f *StatFlags) IsInterruptEnabled(selector LcdInterruptSelector) bool {
	return f.r.IsBitSet(byte(selector))
}
