package cpu

import (
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/mr-tim/goboye/internal/pkg/utils"
)

type Processor interface {
	NextInstruction() Opcode
	DoNextInstruction() uint8
	DebugRegisters() Registers
	GetRegister(reg register) uint8
	GetRegisterPair(pair RegisterPair) uint16
	GetFlagValue(flagName OpResultFlag) bool
	Cycles() uint
	IsStopped() bool
	IsHalted() bool
}

type processor struct {
	registers         *Registers
	savedRegisters    *Registers
	memory            *memory.Controller
	cycles            uint
	interruptsEnabled bool
	isHalted          bool
	isStopped         bool
}

func NewProcessor(memory *memory.Controller) Processor {
	p := processor{
		registers: &Registers{},
		memory:    memory,
	}
	p.registers.pc = uint16(0x0000)
	return &p
}

func (p *processor) readNextInstruction() opcode {
	opCodeByte := p.Read8BitImmediate()
	return LookupOpcode(opCodeByte)
}

func (p *processor) peekNextInstruction() opcode {
	b := p.memory.ReadByte(p.registers.pc)
	return LookupOpcode(b)
}

func (p *processor) NextInstruction() Opcode {
	o := p.peekNextInstruction()
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
	if !p.HandleInterrupts() {
		o := p.readNextInstruction()
		o.handler(o, p)
		p.cycles += uint(o.Cycles())
		return o.Cycles()
	} else {
		// TODO: how many cycles?
		return 0
	}
}

func (p *processor) DebugRegisters() Registers {
	return *p.registers
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

func (p *processor) HandleInterrupts() bool {
	if p.interruptsEnabled {
		eif := p.memory.InterruptEnabled.Read() & p.memory.InterruptFlags.Read()
		addr, flagIndex := memory.GetIsrAddress(eif)
		p.memory.InterruptFlags.Write(utils.UnsetBit(eif, flagIndex))
		if addr != 0x0000 {
			p.serviceInterrupt(addr)
			return true
		}
	}
	return false
}

func (p *processor) serviceInterrupt(address memory.InterruptAddress) {
	//save Registers
	p.savedRegisters = p.registers
	//disable interrupts
	p.interruptsEnabled = false
	//call isr
	doCall16BitAddress(p, uint16(address))
}

func (p *processor) IsHalted() bool {
	return p.isHalted
}

func (p *processor) IsStopped() bool {
	return p.isStopped
}
