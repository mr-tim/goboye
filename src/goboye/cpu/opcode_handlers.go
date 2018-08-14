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

func saveSPToAddr(op opcode, p *processor) {
	addr := p.memory.ReadU16(p.registers.pc)
	sp := p.registers.sp
	p.memory.WriteU16(addr, sp)
	p.registers.pc += 2
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
		flags = flags | uint8(FlagZ)
	}
	if isHalfCarrySubtract(oldValue, 1) {
		flags = flags | uint8(FlagH)
	}
	p.registers.setRegister(RegisterF, flags)

}

func isHalfCarryAdd(old, plusValue uint8) bool {
	return ((old&0x0f)+(plusValue&0x0f))&0x10 == 0x10
}

func isHalfCarrySubtract(old, subtractValue uint8) bool {
	return ((old&0x0f)-(subtractValue&0x0f))&0x80 == 0x80
}

func addRegToA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toAdd := p.registers.getRegister(reg)
		doAddValueToA(p, toAdd)
	}
}

func addHLAddrToA(op opcode, p *processor) {
	toAdd := p.memory.ReadByte(p.registers.hl)
	doAddValueToA(p, toAdd)
}

func addRegAndCarryToA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toAdd := p.registers.getRegister(reg)
		if p.registers.getRegister(RegisterF) & uint8(FlagC) > 0 {
			toAdd++
		}
		doAddValueToA(p, toAdd)
	}
}

func addHLAddrAndCarryToA(op opcode, p *processor) {
	toAdd := p.memory.ReadByte(p.registers.hl)
	if p.registers.getRegister(RegisterF) & uint8(FlagC) > 0 {
		toAdd++
	}
	doAddValueToA(p, toAdd)
}

func doAddValueToA(p *processor, toAdd uint8) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := oldValue + toAdd
	p.registers.setRegister(RegisterA, newValue)
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarryAdd(oldValue, toAdd) {
		flags |= uint8(FlagH)
	}
	if newValue < oldValue {
		flags |= uint8(FlagC)
	}
	p.registers.setRegister(RegisterF, flags)
}

func subtractRegFromA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toSubtract := p.registers.getRegister(reg)
		doSubtractValueFromA(p, toSubtract)
	}
}

func subtractHLAddrFromA(op opcode, p *processor) {
	toSubtract := p.memory.ReadByte(p.registers.hl)
	doSubtractValueFromA(p, toSubtract)
}

func subtractRegAndCarryFromA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toSubtract := p.registers.getRegister(reg)
		if p.registers.getRegister(RegisterF) & uint8(FlagC) > 0 {
			toSubtract++
		}
		doSubtractValueFromA(p, toSubtract)
	}
}

func subtractHLAddrAndCarryFromA(op opcode, p *processor) {
	toSubtract := p.memory.ReadByte(p.registers.hl)
	if p.registers.getRegister(RegisterF) & uint8(FlagC) > 0 {
		toSubtract++
	}
	doSubtractValueFromA(p, toSubtract)
}

func doSubtractValueFromA(p *processor, toSubtract uint8) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := oldValue - toSubtract
	p.registers.setRegister(RegisterA, newValue)
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	if isHalfCarrySubtract(oldValue, toSubtract) {
		flags |= uint8(FlagH)
	}
	if newValue > oldValue {
		flags |= uint8(FlagC)
	}
	p.registers.setRegister(RegisterF, flags)
}

func logicalAndRegAgainstA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		other := p.registers.getRegister(reg)
		doLogicalAndAgainstA(p, other)
	}
}

func logicalAndHLAddrAgainstA(op opcode, p *processor) {
	other := p.memory.ReadByte(p.registers.hl)
	doLogicalAndAgainstA(p, other)
}

func doLogicalAndAgainstA(p *processor, other uint8) {
	flags := doLogicalOpAgainstA(p, other, and)
	flags |= uint8(FlagH)
	p.registers.setRegister(RegisterF, flags)
}

func logicalXorRegAgainstA(reg register) opcodeHandler {
	return func (op opcode, p *processor) {
		other := p.registers.getRegister(reg)
		doLogicalXorAgainstA(p, other)
	}
}

func logicalXorHLAddrAgainstA(op opcode, p *processor) {
	other := p.memory.ReadByte(p.registers.hl)
	doLogicalXorAgainstA(p, other)
}

func doLogicalXorAgainstA(p *processor, other uint8) {
	flags := doLogicalOpAgainstA(p, other, xor)
	p.registers.setRegister(RegisterF, flags)
}

func logicalOrRegAgainstA(reg register) opcodeHandler {
	return func (op opcode, p *processor) {
		other := p.registers.getRegister(reg)
		doLogicalOrAgainstA(p, other)
	}
}

func logicalOrHLAddrAgainstA(op opcode, p *processor) {
	other := p.memory.ReadByte(p.registers.hl)
	doLogicalOrAgainstA(p, other)
}

func doLogicalOrAgainstA(p *processor, other uint8) {
	flags := doLogicalOpAgainstA(p, other, or)
	p.registers.setRegister(RegisterF, flags)
}

func and(a, b uint8) uint8 {
	return a & b
}

func xor(a, b uint8) uint8 {
	return a ^ b
}

func or(a, b uint8) uint8 {
	return a | b
}

func doLogicalOpAgainstA(p *processor, other uint8, op func(a, b uint8) uint8) uint8 {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := oldValue | other
	p.registers.setRegister(RegisterA, newValue)
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if newValue == 0 {
		flags |= uint8(FlagZ)
	}
	return flags
}

func compareRegAgainstA(reg register) opcodeHandler {
	return func (op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		doCompareValueAgainstA(p, value)
	}
}

func compareHLAddrAgainstA(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.hl)
	doCompareValueAgainstA(p, value)
}

func doCompareValueAgainstA(p *processor, value uint8) {
	regAValue := p.registers.getRegister(RegisterA)
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	flags |= uint8(FlagN)
	if value < regAValue {
		flags |= uint8(FlagH)
	} else if value == regAValue {
		flags |= uint8(FlagZ)
	} else if value > regAValue {
		flags |= uint8(FlagC)
	}
}

func testBitOfReg(bit uint8, reg register) opcodeHandler {
	return func (op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		doTestBit(p, bit, value)
	}
}

func testBitOfHLAddr(bit uint8) opcodeHandler {
	return func (op opcode, p *processor) {
		value := p.memory.ReadByte(p.registers.hl)
		doTestBit(p, bit, value)
	}
}

func doTestBit(p *processor, bit, value uint8) {
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	flags |= uint8(FlagH)
	mask := uint8(0x01 << (7-bit))
	if value & mask == 0 {
		flags |= uint8(FlagZ)
	}
}