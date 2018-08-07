package cpu

import "goboye/memory"

type processor struct {
	registers *registers
	memory    memory.MemoryMap
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.memory.ReadByte(p.registers.pc)
	p.registers.pc++
	return lookupOpcode(opCodeByte)
}
