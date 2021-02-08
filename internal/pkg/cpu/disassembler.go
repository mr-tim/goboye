package cpu

import (
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

type Disassembler struct {
	m   *memory.Controller
	pos uint16
}

func (d *Disassembler) SetPos(pos uint16) {
	d.pos = pos
}

func (d *Disassembler) GetNextInstruction() (uint16, OpcodeAndPayload) {
	addr, op, bytesRead := Disassemble(d.m, d.pos)
	d.pos += bytesRead
	return addr, op
}

func Disassemble(m *memory.Controller, pos uint16) (uint16, OpcodeAndPayload, uint16) {
	addr := pos
	opcodeByte := m.ReadAddr(pos)
	pos += 1
	o := LookupOpcode(opcodeByte)

	if o.Code() == OpcodeExtOps.Code() {
		// load the extended code
		opcodeByte = m.ReadAddr(pos)
		pos += 1
		o = LookupExtOpcode(opcodeByte)
	}

	argWidth := o.PayloadLength()
	var payload = make([]byte, argWidth)
	if argWidth == 1 {
		payload[0] = m.ReadAddr(pos)
	} else if argWidth == 2 {
		payload[0] = m.ReadAddr(pos)
		payload[1] = m.ReadAddr(pos + 1)
	}

	op := OpcodeAndPayload{
		op:      &o,
		payload: payload,
	}

	pos += uint16(argWidth)

	bytesRead := pos - addr

	return addr, op, bytesRead
}

func NewDisassembler(m *memory.Controller) Disassembler {
	return Disassembler{
		m:   m,
		pos: 0,
	}
}
