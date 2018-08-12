package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) {}

func load16BitToRegPair(rp registerPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doLoad16BitToRegPair(p, rp)
	}
}

func doLoad16BitToRegPair(p *processor, pair registerPair) {
	value := p.memory.ReadU16(p.registers.pc)
	p.registers.pc += 2
	p.registers.setRegisterPair(pair, value)
}

func load8BitToReg(r register) opcodeHandler {
	return func(op opcode, p *processor) {
		doLoad8BitToReg(p, r)
	}
}

func doLoad8BitToReg(p *processor, reg register) {
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

func incrementRegPair(pair registerPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doIncrementRegPair(p, pair)
	}
}

func decrementRegPair(pair registerPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doDecrementRegPair(p, pair)
	}
}

func doIncrementRegPair(p *processor, rp registerPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)+1)
}

func doDecrementRegPair(p *processor, rp registerPair) {
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

func incrementReg(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		doIncrementRegister(p, reg)
	}
}

func decrementReg(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		doDecrementRegister(p, reg)
	}
}

func doIncrementRegister(p *processor, reg register) {
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

func doDecrementRegister(p *processor, reg register) {
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
