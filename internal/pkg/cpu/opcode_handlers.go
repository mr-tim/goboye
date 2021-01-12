package cpu

import "fmt"

func unimplementedHandler(op opcode, p *processor) {
	panic(fmt.Sprintf("Unimplemented opcode: %#v\n", op))
}

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
	newValue := originalValue + 1
	p.memory.WriteByte(p.registers.hl, newValue)

	flags := FlagC & p.registers.getFlags()
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(originalValue, 1) {
		flags |= FlagH
	}

	p.registers.setFlags(flags)
}

func decrementHLAddr(op opcode, p *processor) {
	originalValue := p.memory.ReadByte(p.registers.hl)
	newValue := originalValue - 1
	p.memory.WriteByte(p.registers.hl, newValue)

	flags := FlagC & p.registers.getFlags() | FlagN
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarrySubtract(originalValue, 1) {
		flags |= FlagH
	}

	p.registers.setFlags(flags)
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
	flags := FlagNoFlags
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(oldValue, 1) {
		flags |= FlagH
	}
	p.registers.setFlags(flags)
}

func doDecrementRegister(p *processor, reg register) {
	oldValue := p.registers.getRegister(reg)
	newValue := oldValue - 1
	p.registers.setRegister(reg, newValue)
	// set the n flag to 1
	flags := FlagN
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarrySubtract(oldValue, 1) {
		flags |= FlagH
	}
	p.registers.setFlags(flags)

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
		if p.registers.getRegister(RegisterF)&uint8(FlagC) > 0 {
			toAdd++
		}
		doAddValueToA(p, toAdd)
	}
}

func addHLAddrAndCarryToA(op opcode, p *processor) {
	toAdd := p.memory.ReadByte(p.registers.hl)
	if p.registers.getRegister(RegisterF)&uint8(FlagC) > 0 {
		toAdd++
	}
	doAddValueToA(p, toAdd)
}

func doAddValueToA(p *processor, toAdd uint8) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := oldValue + toAdd
	p.registers.setRegister(RegisterA, newValue)
	flags := FlagNoFlags
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(oldValue, toAdd) {
		flags |= FlagH
	}
	if newValue < oldValue {
		flags |= FlagC
	}
	p.registers.setFlags(flags)
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
		if p.registers.getRegister(RegisterF)&uint8(FlagC) > 0 {
			toSubtract++
		}
		doSubtractValueFromA(p, toSubtract)
	}
}

func subtractHLAddrAndCarryFromA(op opcode, p *processor) {
	toSubtract := p.memory.ReadByte(p.registers.hl)
	if p.registers.getRegister(RegisterF)&uint8(FlagC) > 0 {
		toSubtract++
	}
	doSubtractValueFromA(p, toSubtract)
}

func doSubtractValueFromA(p *processor, toSubtract uint8) {
	oldValue := p.registers.getRegister(RegisterA)
	newValue := oldValue - toSubtract
	p.registers.setRegister(RegisterA, newValue)
	flags := FlagN
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarrySubtract(oldValue, toSubtract) {
		flags |= FlagH
	}
	if newValue > oldValue {
		flags |= FlagC
	}
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
	if value < regAValue {
		flags |= FlagH
	} else if value == regAValue {
		flags |= FlagZ
	} else if value > regAValue {
		flags |= FlagC
	}
	p.registers.setFlags(flags)
}

func addRegPairToHL(rp RegisterPair) opcodeHandler {
	return func(op opcode, p *processor) {
		p.registers.hl += p.registers.getRegisterPair(rp)
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
	p.registers.setRegister(RegisterA, ^p.registers.getRegister(RegisterA))
}

func setCarryFlag(op opcode, p *processor) {
	flags := p.registers.getFlags() | FlagC
	p.registers.setFlags(flags)
}

func clearCarryFlag(op opcode, p *processor) {
	flags := p.registers.getFlags() | ^FlagC
	p.registers.setFlags(flags)
}

func addImmediate(op opcode, p *processor) {
	original := p.registers.getRegister(RegisterA)
	other := p.Read8BitImmediate()
	result := original + other
	p.registers.setRegister(RegisterA, result)

	flags := FlagNoFlags
	if result == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(original, other) {
		flags |= FlagH
	}
	if result < original {
		flags |= FlagC
	}
	p.registers.setFlags(flags)
}

func subtractImmediate(op opcode, p *processor) {
	original := p.registers.getRegister(RegisterA)
	other := p.Read8BitImmediate()
	result := original - other
	p.registers.setRegister(RegisterA, result)

	flags := FlagN
	if result == 0 {
		flags |= FlagZ
	}
	if isHalfCarrySubtract(original, other) {
		flags |= FlagH
	}
	if result > original {
		flags |= FlagC
	}
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
	original := p.registers.getRegister(RegisterA)
	other := p.Read8BitImmediate()
	if p.registers.getFlagValue(FlagC) {
		other += 1
	}
	result := original + other
	p.registers.setRegister(RegisterA, result)

	flags := FlagNoFlags
	if result == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(original, other) {
		flags |= FlagH
	}
	if result < original {
		flags |= FlagC
	}
	p.registers.setFlags(flags)
}

func subCImmediate(op opcode, p *processor) {
	other := p.Read8BitImmediate()
	if p.registers.getFlagValue(FlagC) {
		other++
	}
	doSubtractValueFromA(p, other)
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
	p.registers = p.savedRegisters
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
	if v > 0 && isHalfCarryAdd16Bit(p.registers.sp, int8(v)) {
		flags |= FlagH
	}
	if v > 0 && isCarryAdd16Bit(p.registers.sp, int8(v)) {
		flags |= FlagC
	}
	p.registers.setFlags(flags)

	p.registers.setRegisterPair(rp, result)
}

func isHalfCarryAdd16Bit(originalValue uint16, operand int8) bool {
	return ((originalValue & 0x0FFF) + uint16(operand)) & 0x1000 == 0x1000
}

func isCarryAdd16Bit(originalValue uint16, operand int8) bool {
	highSet := originalValue & 0x8000 == 0x8000
	highUnsetAfter := (originalValue + uint16(operand)) & 0x8000 == 0x0000
	return highSet && highUnsetAfter
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

	if p.GetFlagValue(FlagH) || (!p.GetFlagValue(FlagN) && (p.GetRegister(RegisterA) & 0x0F) > 9) {
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
