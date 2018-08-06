package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) { }

func load16BitToBC(op opcode, p *processor) {
	value := p.memory.ReadU16(p.programCounter)
	p.programCounter += 2
	p.registers.bc = value
}

func saveAtoBCAddr(op opcode, p *processor) {
	addrLE := p.registers.bc
	addr := ((0xFF00 & addrLE) >> 8) | uint16(0xFF & addrLE) << 8
 	p.memory.WriteByte(addr, p.registers.getRegister(A))
}