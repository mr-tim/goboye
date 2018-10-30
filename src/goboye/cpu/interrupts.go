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

func (r interruptRegister) GetIsrAddress() interruptAddress {
	var addr interruptAddress
	if r.VBlank() {
		addr = vBlankInterrupt
	} else if r.LcdStatus() {
		addr = lcdStatusInterrupt
	} else if r.TimerOverflow() {
		addr = timerOverflowInterrupt
	} else if r.SerialLink() {
		addr = serialLinkInterrupt
	} else if r.JoypadPress() {
		addr = joypadPressInterrupt
	}
	return addr
}
