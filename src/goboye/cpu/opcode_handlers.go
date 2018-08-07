package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) {}

func load16BitToBC(op opcode, p *processor) {
	load16BitToRegPair(p, BC)
}

func load16BitToDE(op opcode, p *processor) {
	load16BitToRegPair(p, DE)
}

func load16BitToHL(op opcode, p *processor) {
	load16BitToRegPair(p, HL)
}

func load16BitToSP(op opcode, p *processor) {
	load16BitToRegPair(p, SP)
}

func load16BitToRegPair(p *processor, pair registerPair) {
	value := p.memory.ReadU16(p.registers.pc)
	p.registers.pc += 2
	p.registers.setRegisterPair(pair, value)
}

func saveAToBCAddr(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.bc, p.registers.getRegister(A))
}

func saveAToDEAddr(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.de, p.registers.getRegister(A))
}

func saveAToHLAddrInc(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, p.registers.getRegister(A))
	p.registers.hl++
}

func saveAToHLAddrDec(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, p.registers.getRegister(A))
	p.registers.hl--
}

func incrementBC(op opcode, p *processor) {
	incrementRegPair(p, BC)
}

func incrementDE(op opcode, p *processor) {
	incrementRegPair(p, DE)
}

func incrementHL(op opcode, p *processor) {
	incrementRegPair(p, HL)
}

func incrementSP(op opcode, p *processor) {
	incrementRegPair(p, SP)
}

func incrementRegPair(p *processor, rp registerPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)+1)
}
