package cpu

import "fmt"

func unsupportedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unsupported opcode: %#v\n", op))
}

func nopHandler(op opcode, p *processor) {}

func load16BitToRegPair(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doLoad16BitToRegPair(p, rp)
	}
}

func doLoad16BitToRegPair(p *processor, pair RegisterPair) {
	value := p.Read16BitImmediate()
	p.registers.setRegisterPair(pair, value)
}

func load8BitToReg(r register) opcodeHandler {
	return func(op opcode, p *processor) {
		doLoad8BitToReg(p, r)
	}
}

func doLoad8BitToReg(p *processor, reg register) {
	value := p.Read8BitImmediate()
	p.registers.setRegister(reg, value)
}

func load8BitToHLAddr(op opcode, p *processor) {
	value := p.Read8BitImmediate()
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
	addr := p.Read16BitImmediate()
	sp := p.registers.sp
	p.memory.WriteU16(addr, sp)
}

func incrementRegPair(pair RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doIncrementRegPair(p, pair)
	}
}

func decrementRegPair(pair RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		doDecrementRegPair(p, pair)
	}
}

func doIncrementRegPair(p *processor, rp RegisterPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)+1)
}

func doDecrementRegPair(p *processor, rp RegisterPair) {
	p.registers.setRegisterPair(rp, p.registers.getRegisterPair(rp)-1)
}

func incrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue, flags := add(originalValue, 1, false)
	p.memory.WriteByte(p.registers.hl, newValue)
	p.registers.setFlags(updateIncDecFlags(p, flags))
}

func decrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue, flags := subtract(originalValue, 1, false)
	p.memory.WriteByte(p.registers.hl, newValue)
	p.registers.setFlags(updateIncDecFlags(p, flags))
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
	newValue, flags := add(oldValue, 1, false)
	p.registers.setRegister(reg, newValue)
	p.registers.setFlags(updateIncDecFlags(p, flags))
}

func doDecrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue, flags := subtract(oldValue, 1, false)
	p.registers.setRegister(reg, newValue)
	p.registers.setFlags(updateIncDecFlags(p, flags))

}

func updateIncDecFlags(p *processor, flags OpResultFlag) OpResultFlag {
	return p.registers.getFlags()&FlagC + (^FlagC & flags)
}

func isHalfCarryAdd(old, plusValue uint8, carry bool) bool {
	result := (old & 0x0f) + (plusValue & 0x0f)
	if carry {
		result += 1
	}
	return result > 0x0f
}

func isHalfCarrySubtract(old, subtractValue uint8, carry bool) bool {
	result := (old & 0x0f) - (subtractValue & 0x0f)
	if carry {
		result -= 1
	}
	return result > 0xF
}

func addRegToA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toAdd := p.registers.getRegister(reg)
		doAddValueToA(p, toAdd, false)
	}
}

func addHLAddrToA(op opcode, p *processor) {
	toAdd := p.memory.ReadByte(p.registers.hl)
	doAddValueToA(p, toAdd, false)
}

func addRegAndCarryToA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toAdd := p.registers.getRegister(reg)
		doAddValueToA(p, toAdd, p.registers.getFlagValue(FlagC))
	}
}

func addHLAddrAndCarryToA(op opcode, p *processor) {
	toAdd := p.memory.ReadByte(p.registers.hl)
	doAddValueToA(p, toAdd, p.registers.getFlagValue(FlagC))
}

func doAddValueToA(p *processor, toAdd uint8, carry bool) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue, flags := add(oldValue, toAdd, carry)
	p.registers.setRegister(RegisterA, newValue)
	p.registers.setFlags(flags)
}

func subtractRegFromA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toSubtract := p.registers.getRegister(reg)
		doSubtractValueFromA(p, toSubtract, false)
	}
}

func subtractHLAddrFromA(op opcode, p *processor) {
	toSubtract := p.memory.ReadByte(p.registers.hl)
	doSubtractValueFromA(p, toSubtract, false)
}

func subtractRegAndCarryFromA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
		toSubtract := p.registers.getRegister(reg)
		doSubtractValueFromA(p, toSubtract, p.registers.getFlagValue(FlagC))
	}
}

func subtractHLAddrAndCarryFromA(op opcode, p *processor) {
	toSubtract := p.memory.ReadByte(p.registers.hl)
	doSubtractValueFromA(p, toSubtract, p.registers.getFlagValue(FlagC))
}

func doSubtractValueFromA(p *processor, toSubtract uint8, carry bool) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue, flags := subtract(oldValue, toSubtract, carry)
	p.registers.setRegister(RegisterA, newValue)
	p.registers.setFlags(flags)
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
	flags |= FlagH
	p.registers.setFlags(flags)
}

func logicalXorRegAgainstA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
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
	p.registers.setFlags(flags)
}

func logicalOrRegAgainstA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
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
	p.registers.setFlags(flags)
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

func doLogicalOpAgainstA(p *processor, other uint8, op func(a, b uint8) uint8) OpResultFlag {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := op(oldValue, other)
	p.registers.setRegister(RegisterA, newValue)
	flags := FlagNoFlags
	if newValue == 0 {
		flags |= FlagZ
	}
	return flags
}

func compareRegAgainstA(reg register) opcodeHandler {
	return func(op opcode, p *processor) {
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
	flags := FlagN
	result := regAValue - value
	isHalfCarry := isHalfCarrySubtract(regAValue, value, false)
	if isHalfCarry {
		flags |= FlagH
	}
	if result == 0 {
		flags |= FlagZ
	}
	if regAValue < value {
		flags |= FlagC
	}
	p.registers.setFlags(flags)
}

func addRegPairToHL(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		flags := p.registers.getFlags() & FlagZ
		original := p.registers.hl
		toAdd := p.registers.getRegisterPair(rp)
		if isHalfCarryAdd16Bit(original, toAdd) {
			flags |= FlagH
		}
		if isCarryAdd16Bit(original, toAdd) {
			flags |= FlagC
		}
		p.registers.hl = original + toAdd
		p.registers.setFlags(flags)
	}
}

func loadAFromRegPairAddr(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.setRegister(RegisterA, p.memory.ReadByte(p.registers.getRegisterPair(rp)))
	}
}

func loadAFromHLAddrInc(op opcode, p *processor) {
	p.registers.setRegister(RegisterA, p.memory.ReadByte(p.registers.hl))
	p.registers.hl += 1
}

func loadAFromHLAddrDec(op opcode, p *processor) {
	p.registers.setRegister(RegisterA, p.memory.ReadByte(p.registers.hl))
	p.registers.hl -= 1
}

func complementOnA(op opcode, p *processor) {
	flags := p.registers.getFlags() | FlagN | FlagH
	p.registers.setRegister(RegisterA, ^p.registers.getRegister(RegisterA))
	p.registers.setFlags(flags)
}

func setCarryFlag(op opcode, p *processor) {
	flags := p.registers.getFlags() | FlagC
	p.registers.setFlags(flags)
}

func clearCarryFlag(op opcode, p *processor) {
	flags := p.registers.getFlags() & ^FlagC
	p.registers.setFlags(flags)
}

func addImmediate(op opcode, p *processor) {
	original := p.registers.getRegister(RegisterA)
	other := p.Read8BitImmediate()
	result, flags := add(original, other, false)
	p.registers.setRegister(RegisterA, result)
	p.registers.setFlags(flags)
}

func subtractImmediate(op opcode, p *processor) {
	original := p.registers.getRegister(RegisterA)
	other := p.Read8BitImmediate()
	result, flags := subtract(original, other, false)
	p.registers.setRegister(RegisterA, result)
	p.registers.setFlags(flags)
}

func logicalAndImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	result := p.registers.getRegister(RegisterA) & other
	p.registers.setRegister(RegisterA, result)

	flags := FlagH
	if result == 0 {
		flags |= FlagZ
	}
	p.registers.setFlags(flags)

}

func logicalOrImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	result := p.registers.getRegister(RegisterA) | other
	p.registers.setRegister(RegisterA, result)

	flags := FlagNoFlags
	if result == 0 {
		flags |= FlagZ
	}
	p.registers.setFlags(flags)
}

func addCImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	doAddValueToA(p, other, p.registers.getFlagValue(FlagC))
}

func subCImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	doSubtractValueFromA(p, other, p.registers.getFlagValue(FlagC))
}

func logicalXorImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	result := p.registers.getRegister(RegisterA) ^ other
	p.registers.setRegister(RegisterA, result)

	flags := FlagNoFlags
	if result == 0 {
		flags |= FlagZ
	}
	p.registers.setFlags(flags)
}

func compareImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	doCompareValueAgainstA(p, other)
}

func relativeJumpImmediate(op opcode, p *processor) {
	jumpValue := p.Read8BitImmediate()

	doRelativeJump(jumpValue, p)
}

func doRelativeJump(unsignedJumpValue uint8, p *processor) {
	// -128 ... -3, -2, -1, 0, 1, 2, 3, ..., 127
	// -126 ... -1,  0,  1, 2, 3, 4, 5, ..., 129
	// -127 ... -2, -1,  1, 2, 3, 4, 5, ..., 129
	jumpValue := int(unsignedJumpValue)
	if jumpValue > 127 {
		jumpValue = -(^jumpValue & 0xFF) - 1
	}
	if jumpValue > 0 {
		p.registers.pc += uint16(jumpValue)
	} else {
		p.registers.pc -= uint16(-jumpValue)
	}
}

func relativeJumpImmediateIfFlag(f OpResultFlag, value bool) opcodeHandler {
	return func(op opcode, p *processor) {
		jumpValue := p.Read8BitImmediate()

		if p.registers.getFlagValue(f) == value {
			p.cycles += 4
			doRelativeJump(jumpValue, p)
		}
	}
}

func jumpToHLAddr(op opcode, p *processor) {
	p.registers.pc = p.registers.hl
}

func jumpTo16BitAddress(op opcode, p *processor) {
	newAddr := p.memory.ReadU16(p.registers.pc)
	p.registers.pc = newAddr
}

func jumpTo16BitAddressIfFlag(f OpResultFlag, value bool) opcodeHandler {
	return func(op opcode, p *processor) {
		newAddr := p.Read16BitImmediate()

		if p.registers.getFlagValue(f) == value {
			p.cycles += 4
			p.registers.pc = newAddr
		}
	}
}

func pushRegisterPair(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.registers.getRegisterPair(rp)
		p.registers.sp -= 2
		p.memory.WriteU16(p.registers.sp, value)
	}
}

func popRegisterPair(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		value := p.memory.ReadU16(p.registers.sp)
		p.registers.sp += 2
		p.registers.setRegisterPair(rp, value)
	}
}

func call16BitAddress(op opcode, p *processor) {
	address := p.Read16BitImmediate()
	doCall16BitAddress(p, address)
}

func doCall16BitAddress(p *processor, address uint16) {
	p.registers.sp -= 2
	p.memory.WriteU16(p.registers.sp, p.registers.pc)
	p.registers.pc = address
}

func conditionalCall16BitAddress(f OpResultFlag, value bool) opcodeHandler {
	return func(op opcode, p *processor) {
		address := p.Read16BitImmediate()
		if p.registers.getFlagValue(f) == value {
			p.cycles += 12
			doCall16BitAddress(p, address)
		}
	}
}

func doReturn(op opcode, p *processor) {
	returnTo := p.memory.ReadU16(p.registers.sp)
	p.registers.sp += 2
	p.registers.pc = returnTo
}

func doReturnEnablingInterrupts(op opcode, p *processor) {
	p.interruptsEnabled = true
	if p.savedRegisters != nil {
		// TODO: check if this is correct? Fixes a NPE but may be incorrect
		p.registers = p.savedRegisters
		p.savedRegisters = nil
	}
	doReturn(op, p)
}

func conditionalReturn(f OpResultFlag, value bool) opcodeHandler {
	return func(op opcode, p *processor) {
		if p.registers.getFlagValue(f) == value {
			p.cycles += 12
			doReturn(op, p)
		}
	}
}

func callRoutineAtAddress(address uint16) opcodeHandler {
	return func(op opcode, p *processor) {
		doCall16BitAddress(p, address)
	}
}

func saveAToFFPlusImmediateAddr(op opcode, p *processor) {
	address := 0xFF00 + uint16(p.Read8BitImmediate())
	saveAToAddr(address, p)
}

func saveAToFFPlusCAddr(op opcode, p *processor) {
	address := 0xFF00 + uint16(p.registers.getRegister(RegisterC))
	saveAToAddr(address, p)
}

func saveATo16BitAddr(op opcode, p *processor) {
	address := p.Read16BitImmediate()
	saveAToAddr(address, p)
}

func saveAToAddr(address uint16, p *processor) {
	p.memory.WriteByte(address, p.registers.getRegister(RegisterA))
}

func loadAFromFFPlusImmediateAddr(op opcode, p *processor) {
	address := 0xFF00 + uint16(p.Read8BitImmediate())
	doLoadAFromAddr(p, address)
}

func loadAFromFFPlusC(op opcode, p *processor) {
	address := 0xFF00 + uint16(p.registers.getRegister(RegisterC))
	doLoadAFromAddr(p, address)
}

func loadAFromAddr(op opcode, p *processor) {
	address := p.Read16BitImmediate()
	doLoadAFromAddr(p, address)
}

func doLoadAFromAddr(p *processor, address uint16) {
	p.registers.setRegister(RegisterA, p.memory.ReadByte(address))
}

func add8BitSignedImmediateToSP(op opcode, p *processor) {
	doAdd8BitSignedImmediateToSP(p, RegisterPairSP)
}

func add8BitImmediateToSPSaveInHL(op opcode, p *processor) {
	doAdd8BitSignedImmediateToSP(p, RegisterPairHL)
}

func doAdd8BitSignedImmediateToSP(p *processor, rp RegisterPair) {
	v := int(p.Read8BitImmediate())
	if v > 127 {
		v -= 256
	}
	result := p.registers.sp
	if v > 0 {
		result += uint16(v)
	} else {
		result -= uint16(-v)
	}

	flags := FlagNoFlags
	if result&0xF < p.registers.sp&0xF {
		flags |= FlagH
	}
	if result&0xFF < p.registers.sp&0xFF {
		flags |= FlagC
	}
	p.registers.setFlags(flags)

	p.registers.setRegisterPair(rp, result)
}

func isHalfCarryAdd16Bit(originalValue uint16, operand uint16) bool {
	return (originalValue&0x0FFF)+(operand&0x0FFF) > 0x0FFF
}

func isCarryAdd16Bit(originalValue uint16, operand uint16) bool {
	result := uint32(originalValue) + uint32(operand)
	return result > 0xFFFF
}

func copyHLToSP(op opcode, p *processor) {
	p.registers.sp = p.registers.hl
}

func disableInterrupts(op opcode, p *processor) {
	p.interruptsEnabled = false
}

func enableInterrupts(op opcode, p *processor) {
	p.interruptsEnabled = true
}

func adjustAForBCDAddition(op opcode, p *processor) {
	var correction uint8

	flags := FlagNoFlags
	if p.GetFlagValue(FlagN) {
		flags |= FlagN
	}

	if p.GetFlagValue(FlagH) || (!p.GetFlagValue(FlagN) && (p.GetRegister(RegisterA)&0x0F) > 9) {
		correction |= 0x06
	}

	if p.GetFlagValue(FlagC) || (!p.GetFlagValue(FlagN) && (p.GetRegister(RegisterA) > 0x99)) {
		correction |= 0x60
		flags |= FlagC
	}

	var corrected uint8
	if !p.GetFlagValue(FlagN) {
		corrected = p.GetRegister(RegisterA) + correction
		p.registers.setRegister(RegisterA, corrected)
	} else {
		corrected = p.GetRegister(RegisterA) - correction
		p.registers.setRegister(RegisterA, corrected)
	}

	if corrected == 0 {
		flags |= FlagZ
	}

	p.registers.setFlags(flags)
}

func stop(op opcode, p *processor) {
	p.isStopped = true
}

func halt(op opcode, p *processor) {
	p.isHalted = true
}
