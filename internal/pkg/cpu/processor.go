package cpu

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/mr-tim/goboye/internal/pkg/utils"
)

const CYCLES_PER_SECOND = 4194304

type Processor interface {
	NextInstruction() Opcode
	DoNextInstruction() uint8
	DebugRegisters() string
	GetRegister(reg register) uint8
	GetRegisterPair(pair RegisterPair) uint16
	GetFlagValue(flagName OpResultFlag) bool
	Cycles() uint
}

type processor struct {
	registers         *registers
	savedRegisters    *registers
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
	return LookupOpcode(opCodeByte)
}

func (p *processor) NextInstruction() Opcode {
	b := p.memory.ReadByte(p.registers.pc)
	o := LookupOpcode(b)
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

func (p *processor) DoNextInstruction() uint8 {
	o := p.readNextInstruction()
	o.handler(o, p)
	p.cycles += uint(o.Cycles())
	return o.Cycles()
}

func (p *processor) DebugRegisters() string {
	return fmt.Sprintf("{af:%04x bc:%04x de:%04x hl:%04x sp:%04x}",
		p.registers.af, p.registers.bc, p.registers.de, p.registers.hl, p.registers.sp)
}

func (p *processor) GetRegister(reg register) uint8 {
	return p.registers.getRegister(reg)
}

func (p *processor) GetRegisterPair(regPair RegisterPair) uint16 {
	return p.registers.getRegisterPair(regPair)
}

func (p *processor) GetFlagValue(flagName OpResultFlag) bool {
	return p.registers.getFlagValue(flagName)
}

func (p *processor) Cycles() uint {
	return p.cycles
}

func (p *processor) HandleInterrupts() {
	if p.interruptsEnabled {
		eif := interruptRegister(p.memory.ReadByte(interruptsEnabledAddress) & p.memory.ReadByte(interruptFlagsAddress))
		addr, flagIndex := eif.GetIsrAddress()
		p.memory.WriteByte(interruptFlagsAddress, utils.UnsetBit(byte(eif), flagIndex))
		p.serviceInterrupt(addr)
	}
}

func (p *processor) serviceInterrupt(address interruptAddress) {
	//save registers
	p.savedRegisters = p.registers
	//disable interrupts
	p.interruptsEnabled = false
	//call isr
	doCall16BitAddress(p, uint16(address))
}
