package cpu

import (
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

type Disassembler struct {
	m   memory.MemoryMap
	pos uint16
}

func (d *Disassembler) SetPos(pos uint16) {
	d.pos = pos
}

func (d *Disassembler) GetNextInstruction() (uint16, OpcodeAndPayload) {
	addr := d.pos
	opcodeByte := d.m.ReadByte(d.pos)
	d.pos += 1
	o := LookupOpcode(opcodeByte)

	if o.Code() == OpcodeExtOps.Code() {
		// load the extended code
		opcodeByte = d.m.ReadByte(d.pos)
		d.pos += 1
		o = LookupExtOpcode(opcodeByte)
	}

	argWidth := o.PayloadLength()
	var payload = make([]byte, argWidth)
	if argWidth == 1 {
		payload = append(payload, d.m.ReadByte(d.pos))
	} else if argWidth == 2 {
		payload = append(payload, d.m.ReadByte(d.pos))
		payload = append(payload, d.m.ReadByte(d.pos+1))
	}

	op := OpcodeAndPayload{
		op:      &o,
		payload: payload,
	}

	d.pos += uint16(argWidth)

	return addr, op
}

func NewDisassembler(m memory.MemoryMap) *Disassembler {
	return &Disassembler{
		m:   m,
		pos: 0,
	}
}
