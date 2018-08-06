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

func TestLoad16BitImmediate(t *testing.T) {
	p := setupHandlerTest([]byte {0x01, 0x34, 0x12})
	readAndPerformNextOp(p)

	assert.Equal(t, uint16(3), p.registers.pc)
	assert.Equal(t, uint16(0), p.registers.sp)
	assert.Equal(t, uint16(0x1234), p.registers.bc)
}

func TestSaveAtoBCAddr(t *testing.T) {
	p := setupHandlerTest([]byte { 0x02, 0x46, 0x27, 0x83, 0x91, 0x27, 0x96})
	p.registers.af = 0x5657
	assert.Equal(t, uint8(0x56), p.registers.getRegister(A))
	p.registers.bc = 0x0400

	readAndPerformNextOp(p)

	assert.Equal(t, uint16(1), p.registers.pc)
	assert.Equal(t, uint16(0), p.registers.sp)
	assert.Equal(t, uint8(0x56), p.memory.ReadByte(0x0400))
	assert.Equal(t, uint8(0x56), p.registers.getRegister(A))
}

func readAndPerformNextOp(p *processor) {
	o := p.readNextInstruction()
	o.handler(o, p)
}

func TestIncBC(t *testing.T) {
	p := setupHandlerTest([]byte { 0x03, 0x03, 0x03, 0x03 })
	p.registers.bc = 0x3782


}