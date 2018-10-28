package cpu

import (
	"fmt"
	"goboye/memory"
)

type Processor interface {
	NextInstruction() Opcode
	DoNextInstruction()
	DebugRegisters() string
	GetRegister(reg register) uint8
	GetRegisterPair(pair registerPair) uint16
	Cycles() uint
}

type processor struct {
	registers         *registers
	memory            memory.MemoryMap
	cycles            uint
	interruptsEnabled bool
}

func NewProcessor(memory memory.MemoryMap) Processor {
	p := processor{
		registers: &registers{},
		memory:    memory,
	}
	p.registers.pc = uint16(0x0000)
	return &p
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.Read8BitImmediate()
	return lookupOpcode(opCodeByte)
}

func (p *processor) NextInstruction() Opcode {
	b := p.memory.ReadByte(p.registers.pc)
	o := lookupOpcode(b)
	return &o
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

func (p *processor) DoNextInstruction() {
	o := p.readNextInstruction()
	o.handler(o, p)
	p.cycles += uint(o.Cycles())
}

func (p *processor) DebugRegisters() string {
	return fmt.Sprintf("{af:%04x bc:%04x de:%04x hl:%04x sp:%04x}",
		p.registers.af, p.registers.bc, p.registers.de, p.registers.hl, p.registers.sp)
}

func (p *processor) GetRegister(reg register) uint8 {
	return p.registers.getRegister(reg)
}

func (p *processor) GetRegisterPair(regPair registerPair) uint16 {
	return p.registers.getRegisterPair(regPair)
}

func (p *processor) Cycles() uint {
	return p.cycles
}
