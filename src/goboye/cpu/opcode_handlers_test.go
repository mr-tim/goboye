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