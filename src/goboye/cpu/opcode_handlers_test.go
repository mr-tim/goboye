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
		programCounter: 0,
		stackPointer: 0,
		registers: rs,
		memory: m,
	}
}

func TestNopHandler(t *testing.T) {
	p := setupHandlerTest([]byte { 0x00 })
	o := p.readNextInstruction()
	o.handler(o, p)

	assert.Equal(t, 1, p.programCounter)
	assert.Equal(t, 0, p.stackPointer)
}

func TestLoad16BitImmediate(t *testing.T) {
	p := setupHandlerTest([]byte {0x01, 0x34, 0x12})
	o := p.readNextInstruction()
	o.handler(o, p)

	assert.Equal(t, 3, p.programCounter)
	assert.Equal(t, 0, p.stackPointer)
	assert.Equal(t, uint16(0x1234), p.registers.bc)
}

func TestSaveAtoBCAddr(t *testing.T) {
	p := setupHandlerTest([]byte { 0x02, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	o := p.readNextInstruction()
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(A))
	p.registers.bc = 0x0004
	o.handler(o, p)

	assert.Equal(t, 1, p.programCounter)
	assert.Equal(t, 0, p.stackPointer)
	assert.Equal(t, uint8(0x56), p.memory.ReadByte(0x0400))
	assert.Equal(t, uint8(0x56), p.registers.getRegister(A))
}