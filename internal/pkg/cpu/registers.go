package cpu

type register int

const (
	RegisterA register = iota
	RegisterF
	RegisterB
	RegisterC
	RegisterD
	RegisterE
	RegisterH
	RegisterL
)

func (r register) String() string {
	return "Register" + string("AFBCDEHL"[r])
}

const (
	RegisterPairAF RegisterPair = iota
	RegisterPairBC
	RegisterPairDE
	RegisterPairHL
	RegisterPairSP
	RegisterPairPC
)

func (rp RegisterPair) String() string {
	return "RegisterPair" + []string{
		"AF", "BC", "DE", "HL", "SP", "PC",
	}[rp]
}

type opResultFlag uint8

const (
	FlagNoFlags opResultFlag = 0x00
	FlagZ       opResultFlag = 0x80
	FlagN       opResultFlag = 0x40
	FlagH       opResultFlag = 0x20
	FlagC       opResultFlag = 0x10
)

type RegisterPair int

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
	if reg == RegisterA || reg == RegisterB || reg == RegisterD || reg == RegisterH {
		shift = 8
	}
	return shift
}

func (r *registers) getRegisterPair(regPair RegisterPair) uint16 {
	if regPair == RegisterPairAF {
		return r.af
	} else if regPair == RegisterPairBC {
		return r.bc
	} else if regPair == RegisterPairDE {
		return r.de
	} else if regPair == RegisterPairHL {
		return r.hl
	} else if regPair == RegisterPairPC {
		return r.pc
	} else {
		return r.sp
	}
}

func (r *registers) setRegisterPair(regPair RegisterPair, value uint16) {
	if regPair == RegisterPairAF {
		r.af = value
	} else if regPair == RegisterPairBC {
		r.bc = value
	} else if regPair == RegisterPairDE {
		r.de = value
	} else if regPair == RegisterPairHL {
		r.hl = value
	} else if regPair == RegisterPairPC {
		r.pc = value
	} else {
		r.sp = value
	}
}

func (r *registers) getRegisterPointer(reg register) *uint16 {
	if reg == RegisterA || reg == RegisterF {
		return &r.af
	} else if reg == RegisterB || reg == RegisterC {
		return &r.bc
	} else if reg == RegisterD || reg == RegisterE {
		return &r.de
	} else {
		return &r.hl
	}
}

func (r *registers) getFlagValue(flag opResultFlag) bool {
	return (uint8(flag) & r.getRegister(RegisterF)) != 0
}
