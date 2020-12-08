package cpu

import (
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupHandlerTest(bytes []byte) *processor {
	m := memory.NewMemoryMapWithBytes(bytes)
	m.GetBootRomRegister().SetBootRomPageDisabled(true)
	rs := &Registers{}
	return &processor{
		registers: rs,
		memory:    m,
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
	p.registers.hl = uint16(0x1478)
	p.DoNextInstruction()

	assert.Equal(t, uint16(2), p.registers.pc)
	assert.Equal(t, uint8(0x37), p.memory.ReadByte(0x1478))
}

func TestSaveAtoBCAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x02, RegisterA, RegisterPairBC)
}

func TestSaveAtoDEAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x12, RegisterA, RegisterPairDE)
}

func TestSaveAtoHLAddrInc(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x22, RegisterA, RegisterPairHL)
	assert.Equal(t, uint16(0x0401), p.registers.getRegisterPair(RegisterPairHL))
}

func TestSaveAtoHLAddrDec(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x32, RegisterA, RegisterPairHL)
	assert.Equal(t, uint16(0x03ff), p.registers.getRegisterPair(RegisterPairHL))
}

func doTestSaveRegisterToRegPairAddr(t *testing.T, op byte, r register, rp RegisterPair) *processor {
	p := setupHandlerTest([]byte{op, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	p.registers.setRegisterPair(rp, 0x0400)

	p.DoNextInstruction()

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint8(0x56), p.memory.ReadByte(0x0400))
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
	t.Run("Simple increment", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x49)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x4a), p.memory.ReadByte(0x1234))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})

	t.Run("Increment Half Carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x4F)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x50), p.memory.ReadByte(0x1234))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})

	t.Run("Increment Zero", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0xFF)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.memory.ReadByte(0x1234))
		checkFlagSet(t, p, FlagZ)
		checkFlagNotSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})
}

func TestDecrementHLAddr(t *testing.T) {
	t.Run("Simple decrement", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x49)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x48), p.memory.ReadByte(0x1234))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})

	t.Run("Decrement Half Carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x10)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x0F), p.memory.ReadByte(0x1234))
		checkFlagNotSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagSet(t, p, FlagH)
	})

	t.Run("Decrement Zero", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x35})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x01)

		p.DoNextInstruction()

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.memory.ReadByte(0x1234))
		checkFlagSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})
}

func TestSaveSPToAddr(t *testing.T) {
	p := setupHandlerTest([]byte{0x08, 0x26, 0x39})
	p.registers.sp = 0x8542
	p.DoNextInstruction()
	assert.Equal(t, uint16(3), p.registers.pc)
	assert.Equal(t, uint8(0x42), p.memory.ReadByte(0x3926))
	assert.Equal(t, uint8(0x85), p.memory.ReadByte(0x3927))
	assert.Equal(t, uint16(0x8542), p.memory.ReadU16(0x3926))
}

func TestJR(t *testing.T) {
	p := setupHandlerTest([]byte{0x18, 0x05})
	p.DoNextInstruction()

	assert.Equal(t, uint(12), p.Cycles())
	assert.Equal(t, uint16(7), p.GetRegisterPair(RegisterPairPC))
}

func TestJRNZ(t *testing.T) {
	doTestJRFlag(t, 0x20, FlagZ, FlagNoFlags)
}

func TestJRZ(t *testing.T) {
	doTestJRFlag(t, 0x28, FlagNoFlags, FlagZ)
}

func TestJRNC(t *testing.T) {
	doTestJRFlag(t, 0x30, FlagC, FlagNoFlags)
}

func TestJRC(t *testing.T) {
	doTestJRFlag(t, 0x38, FlagNoFlags, FlagC)
}

func doTestJRFlag(t *testing.T, opcode byte, noActionFlag OpResultFlag, actionFlag OpResultFlag) {
	t.Run("No action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x05})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(noActionFlag))
		p.DoNextInstruction()

		assert.Equal(t, uint(8), p.Cycles())
		assert.Equal(t, uint16(2), p.GetRegisterPair(RegisterPairPC))
	})
	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x05})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(actionFlag))
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
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(noActionFlag))
		p.DoNextInstruction()

		assert.Equal(t, uint(12), p.Cycles())
		assert.Equal(t, uint16(3), p.GetRegisterPair(RegisterPairPC))
	})
	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0xCD, 0xAB})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(actionFlag))
		p.DoNextInstruction()

		assert.Equal(t, uint(16), p.Cycles())
		assert.Equal(t, uint16(0xABCD), p.GetRegisterPair(RegisterPairPC))
	})
}

func TestCall(t *testing.T) {
	p := setupHandlerTest([]byte{0xCD, 0x34, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	p.registers.setRegisterPair(RegisterPairSP, 10)
	p.DoNextInstruction()

	assert.Equal(t, uint(24), p.Cycles())
	assert.Equal(t, uint16(0x1234), p.GetRegisterPair(RegisterPairPC))
	assert.Equal(t, uint16(8), p.GetRegisterPair(RegisterPairSP))
	assert.Equal(t, uint16(0x0003), p.memory.ReadU16(uint16(8)))
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
		p := setupHandlerTest([]byte{opcode, 0x34, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(noActionFlag))
		p.registers.setRegisterPair(RegisterPairSP, uint16(10))
		p.DoNextInstruction()

		assert.Equal(t, uint(12), p.Cycles())
		assert.Equal(t, uint16(3), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(10), p.GetRegisterPair(RegisterPairSP))
	})

	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x34, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(actionFlag))
		p.registers.setRegisterPair(RegisterPairSP, 10)
		p.DoNextInstruction()

		assert.Equal(t, uint(24), p.Cycles())
		assert.Equal(t, uint16(0x1234), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(8), p.GetRegisterPair(RegisterPairSP))
		assert.Equal(t, uint16(0x0003), p.memory.ReadU16(uint16(8)))
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
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(noActionFlag))
		p.registers.setRegisterPair(RegisterPairSP, uint16(6))
		p.DoNextInstruction()

		assert.Equal(t, uint(8), p.Cycles())
		assert.Equal(t, uint16(0x0001), p.GetRegisterPair(RegisterPairPC))
		assert.Equal(t, uint16(6), p.GetRegisterPair(RegisterPairSP))
	})

	t.Run("Action taken", func(t *testing.T) {
		p := setupHandlerTest([]byte{opcode, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x83, 0x45, 0x67})
		p.registers.setRegister(RegisterF, p.registers.getRegister(RegisterF)|uint8(actionFlag))
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
