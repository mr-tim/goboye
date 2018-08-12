package cpu

import (
	"github.com/stretchr/testify/assert"
	"goboye/memory"
	"testing"
)

func setupHandlerTest(bytes []byte) *processor {
	m := memory.NewMemoryMapWithBytes(bytes)
	rs := &registers{}
	return &processor{
		registers: rs,
		memory:    m,
	}
}

func TestNopHandler(t *testing.T) {
	p := setupHandlerTest([]byte{0x00})
	readAndPerformNextOp(p)

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

func doTestLoad16BitImmediate(t *testing.T, op byte, rp registerPair) {
	p := setupHandlerTest([]byte{op, 0x34, 0x12})
	readAndPerformNextOp(p)

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
	readAndPerformNextOp(p)

	assert.Equal(t, uint16(2), p.registers.pc)
	assert.Equal(t, uint8(0x49), p.registers.getRegister(reg))
}

func TestLoad8BitToHLAddr(t *testing.T) {
	p := setupHandlerTest([]byte{0x36, 0x37})
	p.registers.hl = uint16(0x1478)
	readAndPerformNextOp(p)

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

func doTestSaveRegisterToRegPairAddr(t *testing.T, op byte, r register, rp registerPair) *processor {
	p := setupHandlerTest([]byte{op, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	p.registers.setRegisterPair(rp, 0x0400)

	readAndPerformNextOp(p)

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

func doTestIncrementRegPair(t *testing.T, op byte, rp registerPair) {
	p := setupHandlerTest([]byte{op})
	p.registers.setRegisterPair(rp, 0x13ff)

	readAndPerformNextOp(p)

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0x1400), p.registers.getRegisterPair(rp))
}

func doTestDecrementRegPair(t *testing.T, op byte, rp registerPair) {
	p := setupHandlerTest([]byte{op})
	p.registers.setRegisterPair(rp, 0x13ff)

	readAndPerformNextOp(p)

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
		readAndPerformNextOp(p)
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x48), p.registers.getRegister(reg))
		checkFlagNotSet(t, p, FlagN)
	})
}

func doTestIncrementHalfCarry(op byte, reg register, t *testing.T) {
	t.Run("Increment with half carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x0F)
		readAndPerformNextOp(p)
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
		readAndPerformNextOp(p)
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
		readAndPerformNextOp(p)
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x46), p.registers.getRegister(reg))
		checkFlagSet(t, p, FlagN)
	})
}

func doTestDecrementHalfCarry(op byte, reg register, t *testing.T) {
	t.Run("Decrement with half carry", func(t *testing.T) {
		p := setupHandlerTest([]byte{op})
		p.registers.setRegister(reg, 0x00)
		readAndPerformNextOp(p)
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
		readAndPerformNextOp(p)
		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.registers.getRegister(reg))
		checkFlagSet(t, p, FlagN)
		checkFlagSet(t, p, FlagZ)
	})
}

func checkFlagSet(t *testing.T, p *processor, flag opResultFlag) bool {
	return assert.True(t, p.registers.getFlagValue(flag))
}

func checkFlagNotSet(t *testing.T, p *processor, flag opResultFlag) bool {
	return assert.False(t, p.registers.getFlagValue(flag))
}

func TestIncrementHLAddr(t *testing.T) {
	t.Run("Simple increment", func(t *testing.T) {
		p := setupHandlerTest([]byte{0x34})
		p.registers.setRegisterPair(RegisterPairHL, 0x1234)
		p.memory.WriteByte(0x1234, 0x49)

		readAndPerformNextOp(p)

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

		readAndPerformNextOp(p)

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

		readAndPerformNextOp(p)

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

		readAndPerformNextOp(p)

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

		readAndPerformNextOp(p)

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

		readAndPerformNextOp(p)

		assert.Equal(t, uint16(1), p.registers.pc)
		assert.Equal(t, uint8(0x00), p.memory.ReadByte(0x1234))
		checkFlagSet(t, p, FlagZ)
		checkFlagSet(t, p, FlagN)
		checkFlagNotSet(t, p, FlagH)
	})
}

func readAndPerformNextOp(p *processor) {
	o := p.readNextInstruction()
	o.handler(o, p)
}
