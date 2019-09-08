package cpu

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

type Disassembler struct {
	m   memory.MemoryMap
	pos uint16
}

func (d *Disassembler) SetPos(pos uint16) {
	d.pos = pos
}

func (d *Disassembler) GetNextInstruction() (uint16, Opcode, string) {
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

	payloadStr := ""
	argWidth := o.PayloadLength()
	if argWidth == 1 {
		payloadStr = fmt.Sprintf("0x%02x", d.m.ReadByte(d.pos))
	} else if argWidth == 2 {
		payloadStr = fmt.Sprintf("0x%04x", d.m.ReadU16(d.pos))
	}
	d.pos += uint16(argWidth)

	return addr, &o, payloadStr
}

func NewDisassembler(m memory.MemoryMap) *Disassembler {
	return &Disassembler{
		m:   m,
		pos: 0,
	}
}
