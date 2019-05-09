package cpu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type byteGetter = func(p *processor) uint8
type byteSetter = func(p *processor, value uint8)

var hlAddr = uint16(0x6532)

func getRegisterValue(reg register) byteGetter {
	return func(p *processor) uint8 {
		return p.registers.getRegister(reg)
	}
}

func setRegisterValue(reg register) byteSetter {
	return func(p *processor, value uint8) {
		p.registers.setRegister(reg, value)
	}
}

func getFromHLAddr(p *processor) uint8 {
	return p.memory.ReadByte(hlAddr)
}

func setAtHLAddr(p *processor, value uint8) {
	p.registers.hl = hlAddr
	p.memory.WriteByte(hlAddr, value)
}

func TestRotateLeftWithCarry(t *testing.T) {
	doTestRotateRegLeftWithCarry(t, 0x07, RegisterA)
	doTestRotateRegLeftWithCarry(t, 0x00, RegisterB)
	doTestRotateRegLeftWithCarry(t, 0x01, RegisterC)
	doTestRotateRegLeftWithCarry(t, 0x02, RegisterD)
	doTestRotateRegLeftWithCarry(t, 0x03, RegisterE)
	doTestRotateRegLeftWithCarry(t, 0x04, RegisterH)
	doTestRotateRegLeftWithCarry(t, 0x05, RegisterL)
}

func doTestRotateRegLeftWithCarry(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestRotateLeftWithCarry(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestRotateHLLeftWithCarry(t *testing.T) {
	doTestRotateLeftWithCarry(t, uint8(0x06), getFromHLAddr, setAtHLAddr)
}

func doTestRotateLeftWithCarry(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Rotate 0x85 left with carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x85)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x0B), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Rotate 0x00 left with carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x00)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestRotateARightWithCarry(t *testing.T) {
	doTestRotateRegRightWithCarry(t, 0x0F, RegisterA)
	doTestRotateRegRightWithCarry(t, 0x08, RegisterB)
	doTestRotateRegRightWithCarry(t, 0x09, RegisterC)
	doTestRotateRegRightWithCarry(t, 0x0A, RegisterD)
	doTestRotateRegRightWithCarry(t, 0x0B, RegisterE)
	doTestRotateRegRightWithCarry(t, 0x0C, RegisterH)
	doTestRotateRegRightWithCarry(t, 0x0D, RegisterL)
}

func doTestRotateRegRightWithCarry(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestRotateRightWithCarry(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestRotateHLRightWithCarry(t *testing.T) {
	doTestRotateRightWithCarry(t, uint8(0x0E), getFromHLAddr, setAtHLAddr)
}

func doTestRotateRightWithCarry(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Rotate 0x01 Right with carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x01)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x80), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Rotate 0x00 Right with carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x00)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestRotateLeft(t *testing.T) {
	doTestRotateRegLeft(t, 0x17, RegisterA)
	doTestRotateRegLeft(t, 0x10, RegisterB)
	doTestRotateRegLeft(t, 0x11, RegisterC)
	doTestRotateRegLeft(t, 0x12, RegisterD)
	doTestRotateRegLeft(t, 0x13, RegisterE)
	doTestRotateRegLeft(t, 0x14, RegisterH)
	doTestRotateRegLeft(t, 0x15, RegisterL)
}

func doTestRotateRegLeft(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestRotateLeft(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestRotateHLLeft(t *testing.T) {
	opcode := uint8(0x16)
	doTestRotateLeft(t, opcode, getFromHLAddr, setAtHLAddr)
}

func doTestRotateLeft(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Rotate 0x80 left", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x80)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Rotate 0x11 left", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x11)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x22), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestRotateRight(t *testing.T) {
	doTestRotateRegRight(t, 0x1F, RegisterA)
	doTestRotateRegRight(t, 0x18, RegisterB)
	doTestRotateRegRight(t, 0x19, RegisterC)
	doTestRotateRegRight(t, 0x1A, RegisterD)
	doTestRotateRegRight(t, 0x1B, RegisterE)
	doTestRotateRegRight(t, 0x1C, RegisterH)
	doTestRotateRegRight(t, 0x1D, RegisterL)
}

func doTestRotateRegRight(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestRotateRight(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestRotateHLRight(t *testing.T) {
	doTestRotateRight(t, uint8(0x1E), getFromHLAddr, setAtHLAddr)

}

func doTestRotateRight(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Rotate 0x01 Right", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x01)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Rotate 0x8A Right", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x8A)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x45), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestShiftLeftPreservingSign(t *testing.T) {
	doTestShiftRegLeftPreservingSign(t, 0x27, RegisterA)
	doTestShiftRegLeftPreservingSign(t, 0x20, RegisterB)
	doTestShiftRegLeftPreservingSign(t, 0x21, RegisterC)
	doTestShiftRegLeftPreservingSign(t, 0x22, RegisterD)
	doTestShiftRegLeftPreservingSign(t, 0x23, RegisterE)
	doTestShiftRegLeftPreservingSign(t, 0x24, RegisterH)
	doTestShiftRegLeftPreservingSign(t, 0x25, RegisterL)
}

func doTestShiftRegLeftPreservingSign(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestShiftLeftPreservingSign(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestShiftHLAddrLeftPreservingSign(t *testing.T) {
	doTestShiftLeftPreservingSign(t, uint8(0x26), getFromHLAddr, setAtHLAddr)
}

func doTestShiftLeftPreservingSign(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Shift 0x80 left preserving sign", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x80)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Shift 0xFF left preserving sign", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0xFF)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0xFE), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})
}

func TestShiftRightPreservingSign(t *testing.T) {
	doTestShiftRegRightPreservingSign(t, 0x2F, RegisterA)
	doTestShiftRegRightPreservingSign(t, 0x28, RegisterB)
	doTestShiftRegRightPreservingSign(t, 0x29, RegisterC)
	doTestShiftRegRightPreservingSign(t, 0x2A, RegisterD)
	doTestShiftRegRightPreservingSign(t, 0x2B, RegisterE)
	doTestShiftRegRightPreservingSign(t, 0x2C, RegisterH)
	doTestShiftRegRightPreservingSign(t, 0x2D, RegisterL)
}

func doTestShiftRegRightPreservingSign(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestShiftRightPreservingSign(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestShiftHLAddrRightPreservingSign(t *testing.T) {
	doTestShiftRightPreservingSign(t, uint8(0x2E), getFromHLAddr, setAtHLAddr)
}

func doTestShiftRightPreservingSign(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Shift 0x01 Right preserving sign", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x01)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Shift 0x8A Right preserving sign", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x8A)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0xC5), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestShiftRight(t *testing.T) {
	doTestShiftRegRight(t, 0x3F, RegisterA)
	doTestShiftRegRight(t, 0x38, RegisterB)
	doTestShiftRegRight(t, 0x39, RegisterC)
	doTestShiftRegRight(t, 0x3A, RegisterD)
	doTestShiftRegRight(t, 0x3B, RegisterE)
	doTestShiftRegRight(t, 0x3C, RegisterH)
	doTestShiftRegRight(t, 0x3D, RegisterL)
}

func doTestShiftRegRight(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestShiftRight(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestShiftHLAddrRight(t *testing.T) {
	doTestShiftRight(t, uint8(0x3E), getFromHLAddr, setAtHLAddr)
}

func doTestShiftRight(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Shift 0x01 Right", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x01)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, true, p.registers.getFlagValue(FlagC))
	})

	t.Run("Shift 0x8A Right", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x8A)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x45), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestSwapNybbles(t *testing.T) {
	doTestSwapRegNybbles(t, 0x37, RegisterA)
	doTestSwapRegNybbles(t, 0x30, RegisterB)
	doTestSwapRegNybbles(t, 0x31, RegisterC)
	doTestSwapRegNybbles(t, 0x32, RegisterD)
	doTestSwapRegNybbles(t, 0x33, RegisterE)
	doTestSwapRegNybbles(t, 0x34, RegisterH)
	doTestSwapRegNybbles(t, 0x35, RegisterL)
}

func doTestSwapRegNybbles(t *testing.T, opcode uint8, reg register) {
	t.Run(reg.String(), func(t *testing.T) {
		doTestSwapNybbles(t, opcode, getRegisterValue(reg), setRegisterValue(reg))
	})
}

func TestSwapHLAddrNybbles(t *testing.T) {
	opcode := uint8(0x36)
	doTestSwapNybbles(t, opcode, getFromHLAddr, setAtHLAddr)
}

func doTestSwapNybbles(t *testing.T, opcode uint8, get byteGetter, set byteSetter) {
	t.Run("Swap 0x00", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0x00)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x00), get(p))
		assert.Equal(t, true, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})

	t.Run("Swap 0xF0", func(t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, opcode})
		set(p, 0xF0)
		p.DoNextInstruction()
		assert.Equal(t, uint16(2), p.registers.pc)
		assert.Equal(t, uint8(0x0F), get(p))
		assert.Equal(t, false, p.registers.getFlagValue(FlagZ))
		assert.Equal(t, false, p.registers.getFlagValue(FlagN))
		assert.Equal(t, false, p.registers.getFlagValue(FlagH))
		assert.Equal(t, false, p.registers.getFlagValue(FlagC))
	})
}

func TestCoverageCollection(t *testing.T) {
	doNothing()
}

func TestTestBitsOfA(t *testing.T) {
	doTestTestBit(t, 0, RegisterA, OpcodeExtBit0a)
	doTestTestBit(t, 1, RegisterA, OpcodeExtBit1a)
	doTestTestBit(t, 2, RegisterA, OpcodeExtBit2a)
	doTestTestBit(t, 3, RegisterA, OpcodeExtBit3a)
	doTestTestBit(t, 4, RegisterA, OpcodeExtBit4a)
	doTestTestBit(t, 5, RegisterA, OpcodeExtBit5a)
	doTestTestBit(t, 6, RegisterA, OpcodeExtBit6a)
	doTestTestBit(t, 7, RegisterA, OpcodeExtBit7a)
}

func doTestTestBit(t *testing.T, bit uint8, reg register, op opcode) {
	t.Run(fmt.Sprintf("Bit %d of %s set to 0", bit, reg), func (t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, op.code})
		setRegisterValue(reg)(p, 0)
		p.DoNextInstruction()
		assert.Equal(t, p.registers.getFlagValue(FlagZ), true)
	})
	t.Run(fmt.Sprintf("Bit %d of %s set to 1", bit, reg), func (t *testing.T) {
		p := setupHandlerTest([]byte{0xCB, op.code})
		setRegisterValue(reg)(p, 1 << bit)
		p.DoNextInstruction()
		assert.Equal(t, p.registers.getFlagValue(FlagZ), false)
	})
}