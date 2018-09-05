package cpu

import "goboye/memory"

type processor struct {
	registers *registers
	memory    memory.MemoryMap
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.Read8BitImmediate()
	return lookupOpcode(opCodeByte)
}

func (p *processor) Read8BitImmediate() byte {
	value := p.memory.ReadByte(p.registers.pc)
	p.registers.pc++
	return value
}

func (p *processor) Read16BitImmediate() uint16 {
	value := p.memory.ReadU16(p.registers.pc)
	p.registers.pc += 2
	return value
}