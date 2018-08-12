package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) {}

func load16BitToBC(op opcode, p *processor) {
	load16BitToRegPair(p, RegisterPairBC)
}

func load16BitToDE(op opcode, p *processor) {
	load16BitToRegPair(p, RegisterPairDE)
}

func load16BitToHL(op opcode, p *processor) {
	load16BitToRegPair(p, RegisterPairHL)
}

func load16BitToSP(op opcode, p *processor) {
	load16BitToRegPair(p, RegisterPairSP)
}

func load16BitToRegPair(p *processor, pair registerPair) {
	value := p.memory.ReadU16(p.registers.pc)
	p.registers.pc += 2
	p.registers.setRegisterPair(pair, value)
}

func load8BitToA(op opcode, p *processor) {
	load8BitToReg(p, RegisterA)
}

func load8BitToB(op opcode, p *processor) {
	load8BitToReg(p, RegisterB)
}

func load8BitToC(op opcode, p *processor) {
	load8BitToReg(p, RegisterC)
}

func load8BitToD(op opcode, p *processor) {
	load8BitToReg(p, RegisterD)
}

func load8BitToE(op opcode, p *processor) {
	load8BitToReg(p, RegisterE)
}

func load8BitToH(op opcode, p *processor) {
	load8BitToReg(p, RegisterH)
}

func load8BitToL(op opcode, p *processor) {
	load8BitToReg(p, RegisterL)
}

func load8BitToReg(p *processor, reg register) {
	value := p.memory.ReadByte(p.registers.pc)
	p.registers.pc++
	p.registers.setRegister(reg, value)
}

func load8BitToHLAddr(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.pc)
	p.registers.pc++
	p.memory.WriteByte(p.registers.hl, value)
}

func loadRegToReg(to, from register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(to, p.registers.getRegister(from))
	}
}

func loadHLAddrToReg(to register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.memory.ReadByte(p.registers.hl)
		p.registers.setRegister(to, value)
	}
}

func loadRegToHLAddr(from register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.memory.WriteByte(p.registers.hl, p.registers.getRegister(from))
	}
}

func saveAToBCAddr(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.bc, p.registers.getRegister(RegisterA))
}

func saveAToDEAddr(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.de, p.registers.getRegister(RegisterA))
}

func saveAToHLAddrInc(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, p.registers.getRegister(RegisterA))
	p.registers.hl++
}

func saveAToHLAddrDec(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, p.registers.getRegister(RegisterA))
	p.registers.hl--
}

func incrementBC(op opcode, p *processor) {
	incrementRegPair(p, RegisterPairBC)
}

func decrementBC(op opcode, p *processor) {
	decrementRegPair(p, RegisterPairBC)
}

func incrementDE(op opcode, p *processor) {
	incrementRegPair(p, RegisterPairDE)
}

func decrementDE(op opcode, p *processor) {
	decrementRegPair(p, RegisterPairDE)
}

func incrementHL(op opcode, p *processor) {
	incrementRegPair(p, RegisterPairHL)
}

func decrementHL(op opcode, p *processor) {
	decrementRegPair(p, RegisterPairHL)
}

func incrementSP(op opcode, p *processor) {
	incrementRegPair(p, RegisterPairSP)
}

func decrementSP(op opcode, p *processor) {
	decrementRegPair(p, RegisterPairSP)
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

	flags := uint8(FlagC) & p.registers.getRegister(RegisterF)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarryAdd(originalValue, 1) {
		flags |= uint8(FlagH)
	}

	p.registers.setRegister(RegisterF, flags)
}

func decrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue := originalValue - 1
	p.memory.WriteByte(p.registers.hl, newValue)

	flags := uint8(FlagC)&p.registers.getRegister(RegisterF) | uint8(FlagN)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarrySubtract(originalValue, 1) {
		flags |= uint8(FlagH)
	}

	p.registers.setRegister(RegisterF, flags)
}

func incrementA(op opcode, p *processor) {
	incrementRegister(p, RegisterA)
}

func decrementA(op opcode, p *processor) {
	decrementRegister(p, RegisterA)
}

func incrementB(op opcode, p *processor) {
	incrementRegister(p, RegisterB)
}

func decrementB(op opcode, p *processor) {
	decrementRegister(p, RegisterB)
}

func incrementC(op opcode, p *processor) {
	incrementRegister(p, RegisterC)
}

func decrementC(op opcode, p *processor) {
	decrementRegister(p, RegisterC)
}

func incrementD(op opcode, p *processor) {
	incrementRegister(p, RegisterD)
}

func decrementD(op opcode, p *processor) {
	decrementRegister(p, RegisterD)
}

func incrementE(op opcode, p *processor) {
	incrementRegister(p, RegisterE)
}

func decrementE(op opcode, p *processor) {
	decrementRegister(p, RegisterE)
}

func incrementH(op opcode, p *processor) {
	incrementRegister(p, RegisterH)
}

func decrementH(op opcode, p *processor) {
	decrementRegister(p, RegisterH)
}

func incrementL(op opcode, p *processor) {
	incrementRegister(p, RegisterL)
}

func decrementL(op opcode, p *processor) {
	decrementRegister(p, RegisterL)
}

func incrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue := oldValue + 1
	p.registers.setRegister(reg, newValue)
	flags := p.registers.getRegister(RegisterF)
	// zero the n flag
	flags &= 0xB0
	if newValue == 0 {
		flags |= 0x80
	}
	if isHalfCarryAdd(oldValue, 1) {
		flags |= 0x20
	}
	p.registers.setRegister(RegisterF, flags)
}

func decrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue := oldValue - 1
	p.registers.setRegister(reg, newValue)
	// set the n flag to 1
	flags := p.registers.getRegister(RegisterF) | 0x50
	if newValue == 0 {
		flags = flags | 0x80
	}
	if isHalfCarrySubtract(oldValue, 1) {
		flags = flags | 0x20
	}
	p.registers.setRegister(RegisterF, flags)

}

func isHalfCarryAdd(old, plusValue uint8) bool {
	return ((old&0x0f)+(plusValue&0x0f))&0x10 == 0x10
}

func isHalfCarrySubtract(old, subtractValue uint8) bool {
	return ((old&0x0f)-(subtractValue&0x0f))&0x80 == 0x80
}
