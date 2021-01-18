package cpu

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupHandlerTest(bytes []byte) *processor {
	m := memory.NewControllerWithBytes(bytes)
	// disable the boot rom
	m.BootRomRegister.Write(0x01)
	rs := &Registers{}
	return &processor{
		registers: rs,
		memory:    &m,
	}
}

func TestNopHandler(t *testing.T) {
	p := setupHandlerTest([]byte{0x00})
	p.DoNextInstruction()

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0), p.registers.sp)
}

func TestLoad16BitToBC(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x01, RegisterPairBC)
}

func TestLoad16BitToDE(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x11, RegisterPairDE)
}

func TestLoad16BitToHL(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x21, RegisterPairHL)
}

func TestLoad16BitToSP(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x31, RegisterPairSP)
}

func doTestLoad16BitImmediate(t *testing.T, op byte, rp RegisterPair) {
	p := setupHandlerTest([]byte{op, 0x34, 0x12})
	p.DoNextInstruction()

	assert.Equal(t, uint16(3), p.registers.pc)
	assert.Equal(t, uint16(0x1234), p.registers.getRegisterPair(rp))
}

func TestLoad8BitToA(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x3E, RegisterA)
}

func TestLoad8BitToB(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x06, RegisterB)
}

func TestLoad8BitToC(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x0E, RegisterC)
}

func TestLoad8BitToD(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x16, RegisterD)
}

func TestLoad8BitToE(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x1E, RegisterE)
}

func TestLoad8BitToH(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x26, RegisterH)
}

func TestLoad8BitToL(t *testing.T) {
	doTestLoad8BitImmediate(t, 0x2E, RegisterL)
}

func doTestLoad8BitImmediate(t *testing.T, op uint8, reg register) {
	p := setupHandlerTest([]byte{op, 0x49})
	p.DoNextInstruction()

	assert.Equal(t, uint16(2), p.registers.pc)
	assert.Equal(t, uint8(0x49), p.registers.getRegister(reg))
}

func TestLoad8BitToHLAddr(t *testing.T) {
	p := setupHandlerTest([]byte{0x36, 0x37})
	p.registers.hl = uint16(0x8478)
	p.DoNextInstruction()

	assert.Equal(t, uint16(2), p.registers.pc)
	assert.Equal(t, uint8(0x37), p.memory.ReadByte(0x8478))
}

func TestSaveAtoBCAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x02, RegisterA, RegisterPairBC)
}

func TestSaveAtoDEAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x12, RegisterA, RegisterPairDE)
}

func TestSaveAtoHLAddrInc(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x22, RegisterA, RegisterPairHL)
	assert.Equal(t, uint16(0x9401), p.registers.getRegisterPair(RegisterPairHL))
}

func TestSaveAtoHLAddrDec(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x32, RegisterA, RegisterPairHL)
	assert.Equal(t, uint16(0x93ff), p.registers.getRegisterPair(RegisterPairHL))
}

func doTestSaveRegisterToRegPairAddr(t *testing.T, op byte, r register, rp RegisterPair) *processor {
	p := setupHandlerTest([]byte{op, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	p.registers.setRegisterPair(rp, 0x9400)

	p.DoNextInstruction()

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint8(0x56), p.memory.ReadByte(0x9400))
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	return p
}

func TestIncrementBC(t *testing.T) {
	doTestIncrementRegPair(t, 0x03, RegisterPairBC)
}

func TestDecrementBC(t *testing.T) {
	doTestDecrementRegPair(t, 0x0B, RegisterPairBC)
}

func TestIncrementDE(t *testing.T) {
	doTestIncrementRegPair(t, 0x13, RegisterPairDE)
}

func TestDecrementDE(t *testing.T) {
	doTestDecrementRegPair(t, 0x1B, RegisterPairDE)
}

func TestIncrementHL(t *testing.T) {
	doTestIncrementRegPair(t, 0x23, RegisterPairHL)
}

func TestDecrementHL(t *testing.T) {
	doTestDecrementRegPair(t, 0x2B, RegisterPairHL)
}

func TestIncrementSP(t *testing.T) {
	doTestIncrementRegPair(t, 0x33, RegisterPairSP)
}

func TestDecrementSP(t *testing.T) {
	doTestDecrementRegPair(t, 0x3B, RegisterPairSP)
}

func doTestIncrementRegPair(t *testing.T, op byte, rp RegisterPair) {
	p := setupHandlerTest([]byte{op})
	p.registers.setRegisterPair(rp, 0x13ff)

	p.DoNextInstruction()

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0x1400), p.registers.getRegisterPair(rp))
}

func doTestDecrementRegPair(t *testing.T, op byte, rp RegisterPair) {
	p := setupHandlerTest([]byte{op})
	p.registers.setRegisterPair(rp, 0x13ff)

	p.DoNextInstruction()

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0x13fe), p.registers.getRegisterPair(rp))
}

func TestIncrementA(t *testing.T) {
	doTestIncrementRegister(t, 0x3C, RegisterA)
}

func TestDecrementA(t *testing.T) {
	doTestDecrementRegister(t, 0x3D, RegisterA)
}

func TestIncrementB(t *testing.T) {
	doTestIncrementRegister(t, 0x04, RegisterB)
}

func TestDecrementB(t *testing.T) {
	doTestDecrementRegister(t, 0x05, RegisterB)
}

func TestIncrementC(t *testing.T) {
	doTestIncrementRegister(t, 0x0C, RegisterC)
}

func TestDecrementC(t *testing.T) {
	doTestDecrementRegister(t, 0x0D, RegisterC)
}

func TestIncrementD(t *testing.T) {
	doTestIncrementRegister(t, 0x14, RegisterD)
}

func TestDecrementD(t *testing.T) {
	doTestDecrementRegister(t, 0x15, RegisterD)
}

func TestIncrementE(t *testing.T) {
	doTestIncrementRegister(t, 0x1C, RegisterE)
}

func TestDecrementE(t *testing.T) {
	doTestDecrementRegister(t, 0x1D, RegisterE)
}

func TestIncrementH(t *testing.T) {
	doTestIncrementRegister(t, 0x24, RegisterH)
}

func TestDecrementH(t *testing.T) {
	doTestDecrementRegister(t, 0x25, RegisterH)
}

func TestIncrementL(t *testing.T) {
	doTestIncrementRegister(t, 0x2C, RegisterL)
}

func TestDecrementL(t *testing.T) {
	doTestDecrementRegister(t, 0x2D, RegisterL)
}

func doTestIncrementRegister(t *testing.T, op byte, reg register) {
	doTestSimpleIncrement(op, reg, t)
	doTestIncrementHalfCarry(op, reg, t)
	doTestIncrementZero(op, reg, t)
}

func doTestSimpleIncrement(op byte, reg register, t *testing.T) {
	t.Run("Simple Increment", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x47)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x48), p.registers.getRegister(reg))
		checkFlagNotSet(t, p, FlagN)
	})
}

func doTestIncrementHalfCarry(op byte, reg register, t *testing.T) {
	t.Run("Increment with half carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x0F)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x10), p.registers.getRegister(reg))
		assert.Equal(t, uint8(0x20), p.registers.getRegister(RegisterF))
		checkFlagNotSet(t, p, FlagN)
	})
}

func doTestIncrementZero(op byte, reg register, t *testing.T) {
	t.Run("Increment with zero result", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0xFF)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.registers.getRegister(reg))
		assert.Equal(t, uint8(0xa0), p.registers.getRegister(RegisterF))
		checkFlagNotSet(t, p, FlagN)
	})
}

func doTestDecrementRegister(t *testing.T, op byte, reg register) {
	doTestSimpleDecrement(op, reg, t)
	doTestDecrementHalfCarry(op, reg, t)
	doTestDecrementZero(op, reg, t)

}

func doTestSimpleDecrement(op byte, reg register, t *testing.T) {
	t.Run("Simple Decrement", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x47)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x46), p.registers.getRegister(reg))
		checkFlagSet(t, p, FlagN)
	})
}

func doTestDecrementHalfCarry(op byte, reg register, t *testing.T) {
	t.Run("Decrement with half carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x00)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0xFF), p.registers.getRegister(reg))
		checkFlagSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})
}

func doTestDecrementZero(op byte, reg register, t *testing.T) {
	t.Run("Decrement with zero result", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x01)
		p.DoNextInstruction()
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.registers.getRegister(reg))
		checkFlagSet(t, p, FlagN)
		checkFlagSet(t, p, FlagZ)
	})
}

func checkFlagSet(t *testing.T, p *processor, flag OpResultFlag) bool {
	return assert.True(t, p.registers.getFlagValue(flag))
}

func checkFlagNotSet(t *testing.T, p *processor, flag OpResultFlag) bool {
	return assert.False(t, p.registers.getFlagValue(flag))
}

func TestIncrementHLAddr(t *testing.T) {
	hlAddr = uint16(0x8234)
	t.Run("Simple increment", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0x49)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x4a), p.memory.ReadByte(hlAddr))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})

	t.Run("Increment Half Carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0x4F)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x50), p.memory.ReadByte(hlAddr))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})

	t.Run("Increment Zero", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0xFF)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.memory.ReadByte(hlAddr))
		checkFlagSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})
}

func TestDecrementHLAddr(t *testing.T) {
	hlAddr := uint16(0x8432)
	t.Run("Simple decrement", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0x49)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x48), p.memory.ReadByte(hlAddr))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})

	t.Run("Decrement Half Carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0x10)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x0F), p.memory.ReadByte(hlAddr))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})

	t.Run("Decrement Zero", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, hlAddr)
		p.memory.WriteByte(hlAddr, 0x01)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.memory.ReadByte(hlAddr))
		checkFlagSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})
}

func TestSaveSPToAddr(t *testing.T) {
	p := setupHandlerTest([]byte{0x08, 0x26, 0x89})
	p.registers.sp = 0x8542
	p.DoNextInstruction()
	assert.Equal(t, uint16(3), p.registers.pc)
	assert.Equal(t, uint8(0x42), p.memory.ReadByte(0x8926))
	assert.Equal(t, uint8(0x85), p.memory.ReadByte(0x8927))
	assert.Equal(t, uint16(0x8542), p.memory.ReadU16(0x8926))
}

func TestJR(t *testing.T) {
	p := setupHandlerTest([]byte{0x18, 0x05})
	p.DoNextInstruction()

	assert.Equal(t, uint(12), p.Cycles())
	assert.Equal(t, uint16(7), p.GetRegisterPair(RegisterPairPC))
}

func TestJRNZ(t *testing.T) {
	doTestJRFlag(t, OpcodeJrNzn.code, FlagZ, FlagNoFlags)
}

func TestJRZ(t *testing.T) {
	doTestJRFlag(t, OpcodeJrZn.code, FlagNoFlags, FlagZ)
}

func TestJRNC(t *testing.T) {
	doTestJRFlag(t, OpcodeJrNcn.code, FlagC, FlagNoFlags)
}

func TestJRC(t *testing.T) {
	doTestJRFlag(t, OpcodeJrCn.code, FlagNoFlags, FlagC)
}

func doTestJRFlag(t *testing.T, opcode byte, noActionFlag OpResultFlag, actionFlag OpResultFlag) {
	t.Run("No action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x05})
		setFlag(p, noActionFlag)
		p.DoNextInstruction()

		assert.Equal(t, uint(8), p.Cycles())
		assert.Equal(t, uint16(2), p.GetRegisterPair(RegisterPairPC))
	})
	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x05})
		setFlag(p, actionFlag)
		p.DoNextInstruction()

		assert.Equal(t, uint(12), p.Cycles())
		assert.Equal(t, uint16(7), p.GetRegisterPair(RegisterPairPC))
	})
}

func TestJ(t *testing.T) {
	p := setupHandlerTest([]byte{0xC3, 0xEF, 0xCD})
	p.DoNextInstruction()

	assert.Equal(t, uint(16), p.Cycles())
	assert.Equal(t, uint16(0xCDEF), p.GetRegisterPair(RegisterPairPC))
}

func TestJNZ(t *testing.T) {
	doTestJFlag(t, 0xC2, FlagZ, FlagNoFlags)
}

func TestJZ(t *testing.T) {
	doTestJFlag(t, 0xCA, FlagNoFlags, FlagZ)
}

func TestJNC(t *testing.T) {
	doTestJFlag(t, 0xD2, FlagC, FlagNoFlags)
}

func TestJC(t *testing.T) {
	doTestJFlag(t, 0xDA, FlagNoFlags, FlagC)
}

func doTestJFlag(t *testing.T, opcode byte, noActionFlag OpResultFlag, actionFlag OpResultFlag) {
	t.Run("No action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0xCD, 0xAB})
		setFlag(p, noActionFlag)
		p.DoNextInstruction()

		assert.Equal(t, uint(12), p.Cycles())
		assert.Equal(t, uint16(3), p.GetRegisterPair(RegisterPairPC))
	})
	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0xCD, 0xAB})
		setFlag(p, actionFlag)
		p.DoNextInstruction()

		assert.Equal(t, uint(16), p.Cycles())
		assert.Equal(t, uint16(0xABCD), p.GetRegisterPair(RegisterPairPC))
	})
}

func TestCall(t *testing.T) {
	p := setupHandlerTest([]byte{0xCD, 0x34, 0x12})
	p.registers.setRegisterPair(RegisterPairSP, 0xFFFE)
	p.DoNextInstruction()

	assert.Equal(t, uint(24), p.Cycles())
	assert.Equal(t, uint16(0x1234), p.GetRegisterPair(RegisterPairPC))
	assert.Equal(t, uint16(0xFFFC), p.GetRegisterPair(RegisterPairSP))
	assert.Equal(t, uint16(0x0003), p.memory.ReadU16(uint16(0xFFFC)))
}

func TestCallNZ(t *testing.T) {
	doTestCallFlag(t, 0xC4, FlagZ, FlagNoFlags)
}

func TestCallZ(t *testing.T) {
	doTestCallFlag(t, 0xCC, FlagNoFlags, FlagZ)
}

func TestCallNC(t *testing.T) {
	doTestCallFlag(t, 0xD4, FlagC, FlagNoFlags)
}

func TestCallC(t *testing.T) {
	doTestCallFlag(t, 0xDC, FlagNoFlags, FlagC)
}

func doTestCallFlag(t *testing.T, opcode byte, noActionFlag OpResultFlag, actionFlag OpResultFlag) {
	t.Run("No action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x34, 0x12})
		setFlag(p, noActionFlag)
		p.registers.setRegisterPair(RegisterPairSP, uint16(0xFFFE))
		p.DoNextInstruction()

		assert.Equal(t, uint(12), p.Cycles())
		assert.Equal(t, uint16(3), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(0xFFFE), p.GetRegisterPair(RegisterPairSP))
	})

	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x34, 0x12})
		setFlag(p, actionFlag)
		p.registers.setRegisterPair(RegisterPairSP, 0xFFFE)
		p.DoNextInstruction()

		assert.Equal(t, uint(24), p.Cycles())
		assert.Equal(t, uint16(0x1234), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(0xFFFC), p.GetRegisterPair(RegisterPairSP))
		assert.Equal(t, uint16(0x0003), p.memory.ReadU16(uint16(0xFFFC)))
	})
}

func TestRet(t *testing.T) {
	p := setupHandlerTest([]byte{0xC9, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x83, 0x45, 0x67})
	p.registers.setRegisterPair(RegisterPairSP, uint16(6))
	p.DoNextInstruction()

	assert.Equal(t, uint(16), p.Cycles())
	assert.Equal(t, uint16(0x8359), p.GetRegisterPair(RegisterPairPC))
	assert.Equal(t, uint16(8), p.GetRegisterPair(RegisterPairSP))
}

func TestRetI(t *testing.T) {
	p := setupHandlerTest([]byte{0xD9, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x83, 0x45, 0x67})
	p.savedRegisters = p.registers
	p.registers.setRegisterPair(RegisterPairSP, uint16(6))
	p.interruptsEnabled = false
	p.DoNextInstruction()

	assert.Equal(t, uint(16), p.Cycles())
	assert.Equal(t, uint16(0x8359), p.GetRegisterPair(RegisterPairPC))
	assert.Equal(t, uint16(8), p.GetRegisterPair(RegisterPairSP))
	assert.Equal(t, true, p.interruptsEnabled)
}

func TestRetNZ(t *testing.T) {
	doTestRetFlag(t, 0xC0, FlagZ, FlagNoFlags)
}

func TestRetZ(t *testing.T) {
	doTestRetFlag(t, 0xC8, FlagNoFlags, FlagZ)
}

func TestRetNC(t *testing.T) {
	doTestRetFlag(t, 0xD0, FlagC, FlagNoFlags)
}

func TestRetC(t *testing.T) {
	doTestRetFlag(t, 0xD8, FlagNoFlags, FlagC)
}

func doTestRetFlag(t *testing.T, opcode byte, noActionFlag OpResultFlag, actionFlag OpResultFlag) {
	t.Run("No action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x83, 0x45, 0x67})
		setFlag(p, noActionFlag)
		p.registers.setRegisterPair(RegisterPairSP, uint16(6))
		p.DoNextInstruction()

		assert.Equal(t, uint(8), p.Cycles())
		assert.Equal(t, uint16(0x0001), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(6), p.GetRegisterPair(RegisterPairSP))
	})

	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x83, 0x45, 0x67})
		setFlag(p, actionFlag)
		p.registers.setRegisterPair(RegisterPairSP, uint16(6))
		p.DoNextInstruction()

		assert.Equal(t, uint(20), p.Cycles())
		assert.Equal(t, uint16(0x8359), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(8), p.GetRegisterPair(RegisterPairSP))
	})
}

func TestDisableInterrupts(t *testing.T) {
	p := setupHandlerTest([]byte{0xF3})
	p.interruptsEnabled = true
	p.DoNextInstruction()
	assert.Equal(t, false, p.interruptsEnabled)
}

func TestEnableInterrupts(t *testing.T) {
	p := setupHandlerTest([]byte{0xFB})
	p.interruptsEnabled = false
	p.DoNextInstruction()
	assert.Equal(t, true, p.interruptsEnabled)
}

func TestXorA(t *testing.T) {
	doTestXorReg(t, RegisterA, 0xAF)
}

func TestXorB(t *testing.T) {
	doTestXorReg(t, RegisterB, 0xA8)
}

func TestXorC(t *testing.T) {
	doTestXorReg(t, RegisterC, 0xA9)
}

func TestXorD(t *testing.T) {
	doTestXorReg(t, RegisterD, 0xAA)
}

func TestXorE(t *testing.T) {
	doTestXorReg(t, RegisterE, 0xAB)
}

func TestXorH(t *testing.T) {
	doTestXorReg(t, RegisterH, 0xAC)
}

func TestXorL(t *testing.T) {
	doTestXorReg(t, RegisterL, 0xAD)
}

func doTestXorReg(t *testing.T, reg register, opcode byte) {
	p := setupHandlerTest([]byte{opcode})
	p.registers.setRegister(RegisterA, 0xF0)
	if reg != RegisterA {
		p.registers.setRegister(reg, 0x33)
	}
	p.DoNextInstruction()

	var expected byte
	if reg == RegisterA {
		expected = 0x00
	} else {
		expected = 0xC3
	}
	assert.Equal(t, expected, p.GetRegister(RegisterA))
}

func TestLoadRegToReg(t *testing.T) {
	var bases = map[register]uint8{
		RegisterA: 0x78,
	}

	regs := []register{RegisterB, RegisterC, RegisterD, RegisterE, RegisterH, RegisterL, RegisterF, RegisterA}

	for to, baseAddr := range bases {
		for i, from := range regs {
			if from != RegisterF {
				opcode := baseAddr + uint8(i)
				t.Run(fmt.Sprintf("%x02: %s <- %s", opcode, to, from), func(t *testing.T) {
					doRegToRegTest(t, opcode, to, from)
				})
			}
		}
	}
}

func doRegToRegTest(t *testing.T, opcode uint8, to register, from register) {
	p := setupHandlerTest([]byte{opcode})
	p.registers.setRegister(from, 0x53)
	p.DoNextInstruction()

	assert.Equal(t, uint8(0x53), p.GetRegister(to))
}

func TestLoadHLAddrToReg(t *testing.T) {
	base := uint8(0x46)
	registers := []register{RegisterB, RegisterC, RegisterD, RegisterE, RegisterH, RegisterL, RegisterF, RegisterA}

	for i, to := range registers {
		opcode := base + uint8(i*8)
		if to == RegisterF {
			continue
		}
		t.Run(fmt.Sprintf("%02x: %s < $HL", opcode, to), func(t *testing.T) {
			doHLAddrToRegTest(t, opcode, to)
		})
	}
}

func doHLAddrToRegTest(t *testing.T, opcode uint8, to register) {
	addr := uint16(0x89ab)
	value := uint8(0x79)

	p := setupHandlerTest([]byte{opcode})
	p.memory.WriteByte(addr, value)
	assert.Equal(t, value, p.memory.ReadByte(addr))
	p.registers.hl = addr
	p.DoNextInstruction()
	assert.Equal(t, value, p.GetRegister(to))
}

func TestLoadRegToHLAddr(t *testing.T) {
	base := uint8(0x70)
	registers := []register{RegisterB, RegisterC, RegisterD, RegisterE, RegisterH, RegisterL, RegisterF, RegisterA}

	for i, from := range registers {
		opcode := base + uint8(i)
		if from == RegisterF {
			continue
		}
		t.Run(fmt.Sprintf("%02x: %s> $HL", opcode, from), func(t *testing.T) {
			doRegToHLAddrTest(t, opcode, from)
		})
	}
}

func doRegToHLAddrTest(t *testing.T, opcode uint8, from register) {
	addr := uint16(0x89ac)
	value := uint8(0x96)

	p := setupHandlerTest([]byte{opcode})
	p.registers.setRegister(from, value)
	p.registers.hl = addr
	p.DoNextInstruction()
	actual := p.memory.ReadByte(addr)
	if from == RegisterH {
		assert.Equal(t, uint8(0x89), actual)
	} else if from == RegisterL {
		assert.Equal(t, uint8(0xac), actual)
	} else {
		assert.Equal(t, value, actual)
	}
}

func TestAdd8BitSignedToSPSaveInHL(t *testing.T) {
	doTestAdd8BitSignedToSpSaveInHL(t, 0x0001, 0x01, 0x0002, false, false)
	doTestAdd8BitSignedToSpSaveInHL(t, 0x0001, 0xFF, 0x0000, false, false)
	doTestAdd8BitSignedToSpSaveInHL(t, 0x0FFF, 0x01, 0x1000, false, true)
	doTestAdd8BitSignedToSpSaveInHL(t, 0xFFFF, 0x01, 0x0000, true, true)
}

func doTestAdd8BitSignedToSpSaveInHL(t *testing.T, spValue uint16, operand uint8, expected uint16,
	carry bool, halfCarry bool) {
	p := setupHandlerTest([]byte{0xF8, operand})
	p.registers.sp = spValue
	p.DoNextInstruction()
	assert.Equal(t, expected, p.registers.hl)
	assert.Equal(t, false, p.GetFlagValue(FlagZ))
	assert.Equal(t, false, p.GetFlagValue(FlagN))
	assert.Equal(t, halfCarry, p.GetFlagValue(FlagH))
	assert.Equal(t, carry, p.GetFlagValue(FlagC))
}

func TestPop(t *testing.T) {
	p := setupHandlerTest([]byte{0xC1})
	p.registers.sp = 0xFFFC
	p.memory.WriteByte(0xFFFC, 0x5F)
	p.memory.WriteByte(0xFFFD, 0x3C)
	p.DoNextInstruction()
	assert.Equal(t, uint8(0x3C), p.registers.getRegister(RegisterB))
	assert.Equal(t, uint8(0x5F), p.registers.getRegister(RegisterC))
	assert.Equal(t, uint16(0xFFFE), p.registers.getRegisterPair(RegisterPairSP))
}

func TestPopAF(t *testing.T) {
	p := setupHandlerTest([]byte{0xF1})
	p.registers.sp = 0xFFFC
	p.memory.WriteByte(0xFFFC, 0xFF)
	p.memory.WriteByte(0xFFFD, 0xFF)
	p.DoNextInstruction()

	assert.Equal(t, uint8(0xFF), p.registers.getRegister(RegisterA))
	assert.Equal(t, uint8(0xF0), p.registers.getRegister(RegisterF))
	assert.Equal(t, uint16(0xFFFE), p.registers.getRegisterPair(RegisterPairSP))
}

func TestCompareReg(t *testing.T) {
	p := setupHandlerTest([]byte{0xB8})
	p.registers.setRegister(RegisterA, 0x3C)
	p.registers.setRegister(RegisterB, 0x2F)
	p.DoNextInstruction()
	assert.False(t, p.GetFlagValue(FlagZ))
	assert.True(t, p.GetFlagValue(FlagH))
	assert.True(t, p.GetFlagValue(FlagN))
	assert.False(t, p.GetFlagValue(FlagC))
}

func TestCompareImmediate(t *testing.T) {
	p := setupHandlerTest([]byte{0xFE, 0x3C})
	p.registers.setRegister(RegisterA, 0x3C)
	p.DoNextInstruction()
	assert.True(t, p.GetFlagValue(FlagZ))
	assert.False(t, p.GetFlagValue(FlagH))
	assert.True(t, p.GetFlagValue(FlagN))
	assert.False(t, p.GetFlagValue(FlagC))
}

func TestCompareHL(t *testing.T) {
	p := setupHandlerTest([]byte{0xBE})
	p.registers.setRegister(RegisterA, 0x3C)
	p.registers.setRegisterPair(RegisterPairHL, 0x8004)
	p.memory.WriteByte(0x8004, 0x40)
	p.DoNextInstruction()
	assert.False(t, p.GetFlagValue(FlagZ))
	assert.False(t, p.GetFlagValue(FlagH))
	assert.True(t, p.GetFlagValue(FlagN))
	assert.True(t, p.GetFlagValue(FlagC))
}

func TestAdjustAForBCDAddition(t *testing.T) {
	t.Run("Add register to A then DAA", func(t *testing.T) {
		doDAAAdditionTests(t, func(a, b uint8) *processor {
			p := setupHandlerTest([]byte{0x80, 0x27})
			p.registers.setRegister(RegisterA, a)
			p.registers.setRegister(RegisterB, b)
			return p
		})
	})

	t.Run("Add immediate to A then DAA", func(t *testing.T) {
		doDAAAdditionTests(t, func(a, b uint8) *processor {
			p := setupHandlerTest([]byte{0xC6, b, 0x27})
			p.registers.setRegister(RegisterA, a)
			return p
		})
	})

	t.Run("Subtract register from A then DAA", func(t *testing.T) {
		doDAASubtractionTests(t, func(a, b uint8) *processor {
			p := setupHandlerTest([]byte{0x90, 0x27})
			p.registers.setRegister(RegisterA, a)
			p.registers.setRegister(RegisterB, b)
			return p
		})
	})
}

func doDAAAdditionTests(t *testing.T, init func(a, b uint8) *processor) {
	doDAATest(t, init, 0x66, 0x11, 0x77, false)
	doDAATest(t, init, 0x74, 0x17, 0x91, false)
	doDAATest(t, init, 0x99, 0x01, 0x00, true)
}

func doDAASubtractionTests(t *testing.T, init func(a, b uint8) *processor) {
	doDAATest(t, init, 0x66, 0x11, 0x55, false)
	doDAATest(t, init, 0x83, 0x38, 0x45, false)
	doDAATest(t, init, 0x30, 0x01, 0x29, false)
	doDAATest(t, init, 0x03, 0x04, 0x99, true)
}

func doDAATest(t *testing.T, init func(a, b uint8) *processor, a uint8, b uint8,
	expectedValue uint8, expectCarry bool) {
	p := init(a, b)
	p.DoNextInstruction()
	p.DoNextInstruction()
	assert.Equal(t, expectedValue, p.GetRegister(RegisterA))
	assert.Equal(t, expectedValue == 0x00, p.GetFlagValue(FlagZ))
	assert.Equal(t, expectCarry, p.GetFlagValue(FlagC))
}

func TestDAAExhaustive(t *testing.T) {
	var cases = []struct {
		nBefore bool
		cBefore bool
		hBefore bool
		before  uint8
		after   uint8
		cAfter  bool
	}{
		// post additions
		{false, false, false, 0x99, 0x99, false},
		{false, false, false, 0x7B, 0x81, false},
		{false, false, true, 0x63, 0x69, false},
		{false, false, false, 0xB8, 0x18, true},
		{false, false, false, 0xCC, 0x32, true},
		{false, false, true, 0xA3, 0x09, true},
		{false, true, false, 0x28, 0x88, true},
		{false, true, false, 0x1C, 0x82, true},
		{false, true, true, 0x32, 0x98, true},

		// post subtractions
		{true, false, false, 0x11, 0x11, false},
		{true, false, true, 0x76, 0x70, false},
		{true, true, false, 0xD3, 0x73, true},
		{true, true, true, 0xBB, 0x55, true},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%#v,c=%#v,h=%#v,value=0x%02X", c.nBefore, c.cBefore, c.hBefore, c.before),
			func(t *testing.T) {
				p := setupHandlerTest([]byte{0x27})
				f := FlagNoFlags
				if c.nBefore {
					f |= FlagN
				}
				if c.cBefore {
					f |= FlagC
				}
				if c.hBefore {
					f |= FlagH
				}
				p.registers.setRegister(RegisterA, c.before)
				p.registers.setFlags(f)
				p.DoNextInstruction()

				assert.Equal(t, c.after, p.registers.getRegister(RegisterA))
				assert.Equal(t, c.cAfter, p.registers.getFlagValue(FlagC))
				assert.Equal(t, false, p.registers.getFlagValue(FlagH))
				assert.Equal(t, c.nBefore, p.registers.getFlagValue(FlagN))
				assert.Equal(t, c.after == 0x00, p.registers.getFlagValue(FlagZ))
			})
	}
}

func TestCpnExhaustive(t *testing.T) {
	for a := 0; a < 256; a += 1 {
		for x := 0; x < 256; x += 1 {
			testName := fmt.Sprintf("CP A=0x%02X, 0x%02X: %%s", uint8(a), uint8(x))
			p := setupHandlerTest([]byte{0xFE, uint8(x)})
			p.registers.setRegister(RegisterA, uint8(a))
			p.DoNextInstruction()

			assert.Equal(t, a == x, p.registers.getFlagValue(FlagZ), testName, "Incorrect Z flag")
			assert.True(t, p.registers.getFlagValue(FlagN), testName, "Incorrect N flag")
			h := (a^(a-x)^x)&0x10 != 0
			assert.Equal(t, h, p.registers.getFlagValue(FlagH), testName, "Incorrect H flag")
			assert.Equal(t, a < x, p.registers.getFlagValue(FlagC), testName, "Incorrect C flags")
		}
	}
}

func TestADCnExhaustive(t *testing.T) {
	for a := 0; a < 256; a += 1 {
		for x := 0; x < 256; x += 1 {
			for c := 0; c < 2; c += 1 {
				testName := fmt.Sprintf("ADC A=0x%02X, 0x%02X (C=%t): %%s", uint8(a), uint8(x), c == 1)
				p := setupHandlerTest([]byte{0xCE, uint8(x)})
				p.registers.setRegister(RegisterA, uint8(a))
				if c == 1 {
					p.registers.setFlags(FlagC)
				}
				p.DoNextInstruction()

				expectedValue := uint8(a) + uint8(x) + uint8(c)
				h := uint8(a&0x0f)+uint8(x&0x0f)+uint8(c) > 0x0f
				c := expectedValue < uint8(a) || x != 0 && expectedValue == uint8(a)
				assert.Equal(t, expectedValue, p.registers.getRegister(RegisterA), testName, "Incorrect result")
				assert.Equal(t, expectedValue == 0, p.registers.getFlagValue(FlagZ), testName, "Incorrect Z flag")
				assert.False(t, p.registers.getFlagValue(FlagN), testName, "Incorrect N flag")
				assert.Equal(t, h, p.registers.getFlagValue(FlagH), testName, "Incorrect H flag")
				assert.Equal(t, c, p.registers.getFlagValue(FlagC), testName, "Incorrect C flag")
			}
		}
	}
}
