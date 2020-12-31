package memory

import "github.com/mr-tim/goboye/internal/pkg/utils"

type InterruptAddress uint16

const (
	vBlankInterrupt        InterruptAddress = 0x0040
	lcdStatusInterrupt     InterruptAddress = 0x0048
	timerOverflowInterrupt InterruptAddress = 0x0050
	serialLinkInterrupt    InterruptAddress = 0x0058
	joypadPressInterrupt   InterruptAddress = 0x0060
)

const (
	InterruptsEnabledAddress uint16 = 0xFFFF
	InterruptFlagsAddress    uint16 = 0xFF0F
)

const (
	vBlankIndex        byte = 0
	lcdStatusIndex     byte = 1
	timerOverflowIndex byte = 2
	serialLinkIndex    byte = 3
	joypadPressIndex   byte = 4
)

type InterruptFlagsRegister struct {
	value byte
}

func (r *InterruptFlagsRegister) VBlank() bool {
	return utils.IsBitSet(r.value, vBlankIndex)
}

func (r *InterruptFlagsRegister) VBlankInterrupt() {
	r.value = utils.SetBit(r.value, vBlankIndex)
}

func (r *InterruptFlagsRegister) LcdStatus() bool {
	return utils.IsBitSet(r.value, lcdStatusIndex)
}

func (r *InterruptFlagsRegister) LcdStatusInterrupt() {
	r.value = utils.SetBit(r.value, lcdStatusIndex)
}

func (r *InterruptFlagsRegister) TimerOverflow() bool {
	return utils.IsBitSet(r.value, timerOverflowIndex)
}

func (r *InterruptFlagsRegister) TimerOverflowInterrupt() {
	r.value = utils.SetBit(r.value, timerOverflowIndex)
}

func (r *InterruptFlagsRegister) SerialLink() bool {
	return utils.IsBitSet(r.value, serialLinkIndex)
}

func (r *InterruptFlagsRegister) SerialLinkInterrupt() {
	r.value = utils.SetBit(r.value, serialLinkIndex)
}

func (r *InterruptFlagsRegister) JoypadPress() bool {
	return utils.IsBitSet(r.value, joypadPressIndex)
}

func (r *InterruptFlagsRegister) JoypadPressInterrupt() {
	r.value = utils.SetBit(r.value, joypadPressIndex)
}

func GetIsrAddress(eif byte) (InterruptAddress, byte) {
	r := InterruptFlagsRegister{value: eif}
	if r.VBlank() {
		return vBlankInterrupt, vBlankIndex
	} else if r.LcdStatus() {
		return lcdStatusInterrupt, lcdStatusIndex
	} else if r.TimerOverflow() {
		return timerOverflowInterrupt, timerOverflowIndex
	} else if r.SerialLink() {
		return serialLinkInterrupt, serialLinkIndex
	} else if r.JoypadPress() {
		return joypadPressInterrupt, joypadPressIndex
	}
	return 0x0000, 255
}

func (r *InterruptFlagsRegister) Read() byte {
	return r.value
}

func (r *InterruptFlagsRegister) Write(value byte) {
	r.value = value
}

type InterruptEnabledRegister struct {
	value byte
}

func (r *InterruptEnabledRegister) Read() byte {
	return r.value
}

func (r *InterruptEnabledRegister) Write(value byte) {
	r.value = value
}
