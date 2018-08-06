package cpu

import "goboye/memory"

type processor struct {
	programCounter uint16
	stackPointer uint16

	registers *registers
	memory memory.MemoryMap
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.memory.ReadByte(p.programCounter)
	p.programCounter++
	return lookupOpcode(opCodeByte)
}