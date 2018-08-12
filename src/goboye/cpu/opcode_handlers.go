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

func decrementBC(op opcode, p *processor) {
	decrementRegPair(p, BC)
}

func incrementDE(op opcode, p *processor) {
	incrementRegPair(p, DE)
}

func decrementDE(op opcode, p *processor) {
	decrementRegPair(p, DE)
}

func incrementHL(op opcode, p *processor) {
	incrementRegPair(p, HL)
}

func decrementHL(op opcode, p *processor) {
	decrementRegPair(p, HL)
}

func incrementSP(op opcode, p *processor) {
	incrementRegPair(p, SP)
}

func decrementSP(op opcode, p *processor) {
	decrementRegPair(p, SP)
}

func incrementRegPair(p *processor, rp registerPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)+1)
}

func decrementRegPair(p *processor, rp registerPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)-1)
}

func incrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue := originalValue + 1
	p.memory.WriteByte(p.registers.hl, newValue)

	flags := uint8(FlagC) & p.registers.getRegister(F)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarryAdd(originalValue, 1) {
		flags |= uint8(FlagH)
	}

	p.registers.setRegister(F, flags)
}

func decrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue := originalValue - 1
	p.memory.WriteByte(p.registers.hl, newValue)

	flags := uint8(FlagC)&p.registers.getRegister(F) | uint8(FlagN)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarrySubtract(originalValue, 1) {
		flags |= uint8(FlagH)
	}

	p.registers.setRegister(F, flags)
}

func incrementA(op opcode, p *processor) {
	incrementRegister(p, A)
}

func decrementA(op opcode, p *processor) {
	decrementRegister(p, A)
}

func incrementB(op opcode, p *processor) {
	incrementRegister(p, B)
}

func decrementB(op opcode, p *processor) {
	decrementRegister(p, B)
}

func incrementC(op opcode, p *processor) {
	incrementRegister(p, C)
}

func decrementC(op opcode, p *processor) {
	decrementRegister(p, C)
}

func incrementD(op opcode, p *processor) {
	incrementRegister(p, D)
}

func decrementD(op opcode, p *processor) {
	decrementRegister(p, D)
}

func incrementE(op opcode, p *processor) {
	incrementRegister(p, E)
}

func decrementE(op opcode, p *processor) {
	decrementRegister(p, E)
}

func incrementH(op opcode, p *processor) {
	incrementRegister(p, H)
}

func decrementH(op opcode, p *processor) {
	decrementRegister(p, H)
}

func incrementL(op opcode, p *processor) {
	incrementRegister(p, L)
}

func decrementL(op opcode, p *processor) {
	decrementRegister(p, L)
}

func incrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue := oldValue + 1
	p.registers.setRegister(reg, newValue)
	flags := p.registers.getRegister(F)
	// zero the n flag
	flags &= 0xB0
	if newValue == 0 {
		flags |= 0x80
	}
	if isHalfCarryAdd(oldValue, 1) {
		flags |= 0x20
	}
	p.registers.setRegister(F, flags)
}

func decrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue := oldValue - 1
	p.registers.setRegister(reg, newValue)
	// set the n flag to 1
	flags := p.registers.getRegister(F) | 0x50
	if newValue == 0 {
		flags = flags | 0x80
	}
	if isHalfCarrySubtract(oldValue, 1) {
		flags = flags | 0x20
	}
	p.registers.setRegister(F, flags)

}

func isHalfCarryAdd(old, plusValue uint8) bool {
	return ((old&0x0f)+(plusValue&0x0f))&0x10 == 0x10
}

func isHalfCarrySubtract(old, subtractValue uint8) bool {
	return ((old&0x0f)-(subtractValue&0x0f))&0x80 == 0x80
}
