package cpu

func extendedOps(op opcode, p *processor) {
	opCodeByte := p.memory.ReadByte(p.registers.pc)
	p.registers.pc++
	extendedOpcode := opcodeMapExt[opCodeByte]
	extendedOpcode.handler(extendedOpcode, p)
}

func testBitOfReg(bit uint8, reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		doTestBit(p, bit, value)
	}
}

func testBitOfHLAddr(bit uint8) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.memory.ReadByte(p.registers.hl)
		doTestBit(p, bit, value)
	}
}

func doTestBit(p *processor, bit, value uint8) {
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	flags |= uint8(FlagH)
	mask := uint8(0x01 << bit)
	if value&mask == 0 {
		flags |= uint8(FlagZ)
	}
	p.registers.setRegister(RegisterF, flags)
}

func clearBitOfReg(bit uint8, reg register) opcodeHandler {
	mask := uint8(0xFF) - uint8(0x01<<bit)
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, p.registers.getRegister(reg)&mask)
	}
}

func clearBitOfHLAddr(bit uint8) opcodeHandler {
	mask := uint8(0xFF) - uint8(0x01<<bit)
	return func(op opcode, p *processor) {
		p.memory.WriteByte(p.registers.hl, p.memory.ReadByte(p.registers.hl)&mask)
	}
}

func setBitOfReg(bit uint8, reg register) opcodeHandler {
	mask := uint8(0x01 << bit)
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, p.registers.getRegister(reg)|mask)
	}
}

func setBitOfHLAddr(bit uint8) opcodeHandler {
	mask := uint8(0x01 << bit)
	return func(op opcode, p *processor) {
		p.memory.WriteByte(p.registers.hl, p.memory.ReadByte(p.registers.hl)|mask)
	}
}

func rotateRegLeftWithCarry(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, doRotateLeft(p, p.registers.getRegister(reg), true))
	}
}

func rotateHLAddrLeftWithCarry(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, doRotateLeft(p, p.memory.ReadByte(p.registers.hl), true))
}

func doRotateLeft(p *processor, value uint8, carry bool) uint8 {
	result := value << 1
	if carry {
		result |= value >> 7
	} else if p.GetFlagValue(FlagC) {
		result |= 0x01
	}
	setLeftShiftFlags(p, result, value)
	return result
}

func setLeftShiftFlags(p *processor, result uint8, value uint8) {
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if result == 0 {
		flags |= uint8(FlagZ)
	}
	if value&uint8(0x80) != 0 {
		flags |= uint8(FlagC)
	}
	p.registers.setRegister(RegisterF, flags)
}

func rotateRegRightWithCarry(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, doRotateRight(p, p.registers.getRegister(reg), true))
	}
}

func rotateHLAddrRightWithCarry(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, doRotateRight(p, p.memory.ReadByte(p.registers.hl), true))
}

func doRotateRight(p *processor, value uint8, carry bool) uint8 {
	result := value >> 1
	if carry {
		result |= value << 7
	} else if p.GetFlagValue(FlagC) {
		result |= 0x80
	}

	setRightShiftFlags(p, result, value)
	return result
}

func setRightShiftFlags(p *processor, result uint8, value uint8) {
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if result == 0 {
		flags |= uint8(FlagZ)
	}
	if value&uint8(0x01) != 0 {
		flags |= uint8(FlagC)
	}
	p.registers.setRegister(RegisterF, flags)
}

func rotateRegLeft(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, doRotateLeft(p, p.registers.getRegister(reg), false))
	}
}

func rotateHLAddrLeft(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, doRotateLeft(p, p.memory.ReadByte(p.registers.hl), false))
}

func rotateRegRight(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(reg, doRotateRight(p, p.registers.getRegister(reg), false))
	}
}

func rotateHLAddrRight(op opcode, p *processor) {
	p.memory.WriteByte(p.registers.hl, doRotateRight(p, p.memory.ReadByte(p.registers.hl), false))
}

func shiftRegLeftPreservingSign(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		p.registers.setRegister(reg, doShiftLeftPreservingSign(p, value))
	}
}

func shiftHLAddrLeftPreservingSign(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.hl)
	p.memory.WriteByte(p.registers.hl, doShiftLeftPreservingSign(p, value))
}

func doShiftLeftPreservingSign(p *processor, value uint8) uint8 {
	result := value << 1
	setLeftShiftFlags(p, result, value)
	return result
}

func shiftRegRightPreservingSign(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		p.registers.setRegister(reg, doShiftRightPreservingSign(p, value))
	}
}

func shiftHLAddrRightPreservingSign(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.hl)
	p.memory.WriteByte(p.registers.hl, doShiftRightPreservingSign(p, value))
}

func doShiftRightPreservingSign(p *processor, value uint8) uint8 {
	result := uint8(value&0x80) | (value >> 1)
	setRightShiftFlags(p, result, value)
	return result
}

func shiftRegRight(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		p.registers.setRegister(reg, doShiftRight(p, value))
	}
}

func shiftHLAddrRight(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.hl)
	p.memory.WriteByte(p.registers.hl, doShiftRight(p, value))
}

func doShiftRight(p *processor, value uint8) uint8 {
	result := value >> 1
	setRightShiftFlags(p, result, value)
	return result
}

func swapRegNybbles(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegister(reg)
		p.registers.setRegister(reg, doSwapNybbles(p, value))
	}
}

func swapHLAddrNybbles(op opcode, p *processor) {
	value := p.memory.ReadByte(p.registers.hl)
	p.memory.WriteByte(p.registers.hl, doSwapNybbles(p, value))
}

func doSwapNybbles(p *processor, value uint8) uint8 {
	h := 0xF0 & value
	l := 0x0F & value
	result := (l << 4) | (h >> 4)
	flags := p.registers.getRegister(RegisterF) & uint8(0x0F)
	if result == 0 {
		flags |= uint8(FlagZ)
	}
	p.registers.setRegister(RegisterF, flags)
	return result
}
