package cpu

type register int

const (
	A = iota
	F
	B
	C
	D
	E
	H
	L
)

const (
	AF registerPair = iota
	BC
	DE
	HL
	SP
	PC
)

type registerPair int

type registers struct {
	af, bc, de, hl uint16
	sp, pc         uint16
}

func (r *registers) getRegister(reg register) uint8 {
	shift := r.getShift(reg)
	ptr := r.getRegisterPointer(reg)
	result := uint8(*ptr >> shift)
	return result
}

func (r *registers) setRegister(reg register, value uint8) {
	shift := r.getShift(reg)
	ptr := r.getRegisterPointer(reg)
	orig := *ptr
	var x, y uint16
	if shift == 8 {
		//update high byte
		x = uint16(value) << shift
		y = orig & 0x00FF
	} else {
		//update low
		x = orig & 0xFF00
		y = uint16(value)
	}

	*ptr = x | y
}

func (r *registers) getShift(reg register) uint8 {
	shift := uint8(0)
	if reg == A || reg == B || reg == D || reg == H {
		shift = 8
	}
	return shift
}

func (r *registers) getRegisterPair(regPair registerPair) uint16 {
	if regPair == AF {
		return r.af
	} else if regPair == BC {
		return r.bc
	} else if regPair == DE {
		return r.de
	} else if regPair == HL {
		return r.hl
	} else if regPair == PC {
		return r.pc
	} else {
		return r.sp
	}
}

func (r *registers) setRegisterPair(regPair registerPair, value uint16) {
	if regPair == AF {
		r.af = value
	} else if regPair == BC {
		r.bc = value
	} else if regPair == DE {
		r.de = value
	} else if regPair == HL {
		r.hl = value
	} else if regPair == PC {
		r.pc = value
	} else {
		r.sp = value
	}
}

func (r *registers) getRegisterPointer(reg register) *uint16 {
	if reg == A || reg == F {
		return &r.af
	} else if reg == B || reg == C {
		return &r.bc
	} else if reg == D || reg == E {
		return &r.de
	} else {
		return &r.hl
	}
}
