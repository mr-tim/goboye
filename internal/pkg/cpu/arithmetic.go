package cpu

func add(a, b uint8, carry bool) (uint8, OpResultFlag) {
	newValue := a + b
	if carry {
		newValue += 1
	}
	flags := FlagNoFlags
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarryAdd(a, b, carry) {
		flags |= FlagH
	}
	if newValue < a || b != 0 && newValue == a {
		flags |= FlagC
	}
	return newValue, flags
}

func subtract(a, b uint8, carry bool) (uint8, OpResultFlag) {
	newValue := a - b
	if carry {
		newValue -= 1
	}
	flags := FlagN
	if newValue == 0 {
		flags |= FlagZ
	}
	if isHalfCarrySubtract(a, b, carry) {
		flags |= FlagH
	}
	if newValue > a || b != 0 && a == newValue {
		flags |= FlagC
	}
	return newValue, flags
}
