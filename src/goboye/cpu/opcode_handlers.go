package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) { }

func load16BitToBC(op opcode, p *processor) {
	load16BitToRegPair(p, BC)
}

func load16BitToRegPair(p *processor, pair registerPair) {
	value := p.memory.ReadU16(p.registers.pc)
	p.registers.pc += 2
	p.registers.setRegisterPair(pair, value)
}

func saveAtoBCAddr(op opcode, p *processor) {
	addr := p.registers.bc
	p.memory.WriteByte(addr, p.registers.getRegister(A))
}