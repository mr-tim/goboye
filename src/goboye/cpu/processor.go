package cpu

import "goboye/memory"

type processor struct {
	programCounter int
	stackPointer int

	registers *registers
	memory memory.MemoryMap
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.memory.ReadByte(p.programCounter)
	p.programCounter++
	return lookupOpcode(opCodeByte)
}