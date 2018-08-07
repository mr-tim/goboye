package cpu

import (
	"testing"
	"goboye/memory"
	"github.com/stretchr/testify/assert"
)

func setupHandlerTest(bytes []byte) *processor {
	m := memory.NewMemoryMapWithBytes(bytes)
	rs := &registers {}
	return &processor{
		registers: rs,
		memory: m,
	}
}

func TestNopHandler(t *testing.T) {
	p := setupHandlerTest([]byte { 0x00 })
	readAndPerformNextOp(p)

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0), p.registers.sp)
}

func TestLoad16BitToBC(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x01, BC)
}

func TestLoad16BitToDE(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x11, DE)
}

func TestLoad16BitToHL(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x21, HL)
}

func TestLoad16BitToSP(t *testing.T) {
	doTestLoad16BitImmediate(t, 0x31, SP)
}

func doTestLoad16BitImmediate(t *testing.T, op byte, rp registerPair) {
	p := setupHandlerTest([]byte {op, 0x34, 0x12})
	readAndPerformNextOp(p)

	assert.Equal(t, uint16(3), p.registers.pc)
	assert.Equal(t, uint16(0x1234), p.registers.getRegisterPair(rp))
}

func TestSaveAtoBCAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x02, A, BC)
}

func TestSaveAtoDEAddr(t *testing.T) {
	doTestSaveRegisterToRegPairAddr(t, 0x12, A, DE)
}

func TestSaveAtoHLAddrInc(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x22, A, HL)
	assert.Equal(t, uint16(0x0401), p.registers.getRegisterPair(HL))
}

func TestSaveAtoHLAddrDec(t *testing.T) {
	p := doTestSaveRegisterToRegPairAddr(t, 0x32, A, HL)
	assert.Equal(t, uint16(0x03ff), p.registers.getRegisterPair(HL))
}

func doTestSaveRegisterToRegPairAddr(t *testing.T, op byte, r register, rp registerPair) *processor {
	p := setupHandlerTest([]byte { op, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	p.registers.setRegisterPair(rp, 0x0400)

	readAndPerformNextOp(p)

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint8(0x56), p.memory.ReadByte(0x0400))
	assert.Equal(t, uint8(0x56), p.registers.getRegister(r))
	return p
}

func readAndPerformNextOp(p *processor) {
	o := p.readNextInstruction()
	o.handler(o, p)
}