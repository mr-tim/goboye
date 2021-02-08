package register

import (
	"github.com/mr-tim/goboye/internal/pkg/utils"
)

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
	value byte
}

func (f *StatFlags) Read() byte {
	return f.value
}

func (f *StatFlags) Write(value byte) {
	mask := uint8(0x78)
	maskedValue := mask & value
	currentFlags := ^mask & f.value
	f.value = currentFlags | maskedValue
}

func (f *StatFlags) GetMode() LcdcMode {
	return LcdcMode(f.value & 0x03)
}

func (f *StatFlags) SetMode(mode LcdcMode) {
	updated := (f.value & 0xFC) | byte(mode)
	f.value = updated
}

func (f *StatFlags) IsInterruptEnabled(selector LcdInterruptSelector) bool {
	return utils.IsBitSet(f.value, byte(selector))
}
