package cpu

import "goboye/utils"

type interruptAddress uint16

const (
	vBlankInterrupt        interruptAddress = 0x0040
	lcdStatusInterrupt     interruptAddress = 0x0048
	timerOverflowInterrupt interruptAddress = 0x0050
	serialLinkInterrupt    interruptAddress = 0x0058
	joypadPressInterrupt   interruptAddress = 0x0060
)

type interruptRegister byte

const (
	interruptsEnabledAddress uint16 = 0xFFFF
	interruptFlagsAddress    uint16 = 0xFF0F
)

func (r interruptRegister) VBlank() bool {
	return utils.IsBitSet(byte(r), 0)
}

func (r interruptRegister) LcdStatus() bool {
	return utils.IsBitSet(byte(r), 1)
}

func (r interruptRegister) TimerOverflow() bool {
	return utils.IsBitSet(byte(r), 2)
}

func (r interruptRegister) SerialLink() bool {
	return utils.IsBitSet(byte(r), 3)
}

func (r interruptRegister) JoypadPress() bool {
	return utils.IsBitSet(byte(r), 4)
}

func (r interruptRegister) GetIsrAddress() (interruptAddress, byte) {
	if r.VBlank() {
		return vBlankInterrupt, 0
	} else if r.LcdStatus() {
		return lcdStatusInterrupt, 1
	} else if r.TimerOverflow() {
		return timerOverflowInterrupt, 2
	} else if r.SerialLink() {
		return serialLinkInterrupt, 3
	} else if r.JoypadPress() {
		return joypadPressInterrupt, 4
	}
	return 0x0000, 255
}
