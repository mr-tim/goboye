package cpu

import "fmt"

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

type OpResultFlag uint8

const (
	FlagNoFlags OpResultFlag = 0x00
	FlagZ       OpResultFlag = 0x80
	FlagN       OpResultFlag = 0x40
	FlagH       OpResultFlag = 0x20
	FlagC       OpResultFlag = 0x10
)

type RegisterPair int

type Registers struct {
	af, bc, de, hl uint16
	sp, pc         uint16
}

func (r *Registers) getRegister(reg register) uint8 {
	shift := r.getShift(reg)
	ptr := r.getRegisterPointer(reg)
	result := uint8(*ptr >> shift)
	return result
}

func (r *Registers) setRegister(reg register, value uint8) {
	shift := r.getShift(reg)
	ptr := r.getRegisterPointer(reg)
	if reg == RegisterF {
		value &= 0xF0
	}
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

func (r *Registers) setFlags(flags OpResultFlag) {
	updated := (r.af & 0xFF0F) | (uint16(flags) & 0x00F0)
	r.af = updated
}

func (r *Registers) getFlags() OpResultFlag {
	return OpResultFlag(r.getRegister(RegisterF) & 0xF0)
}

func (r *Registers) getShift(reg register) uint8 {
	shift := uint8(0)
	if reg == RegisterA || reg == RegisterB || reg == RegisterD || reg == RegisterH {
		shift = 8
	}
	return shift
}

func (r *Registers) getRegisterPair(regPair RegisterPair) uint16 {
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

func (r *Registers) setRegisterPair(regPair RegisterPair, value uint16) {
	if regPair == RegisterPairAF {
		value &= 0xFFF0
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

func (r *Registers) getRegisterPointer(reg register) *uint16 {
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

func (r *Registers) getFlagValue(flag OpResultFlag) bool {
	return (uint8(flag) & r.getRegister(RegisterF)) != 0
}

func (r *Registers) String() string {
	return fmt.Sprintf("{af:%04x bc:%04x de:%04x hl:%04x sp:%04x}", r.af, r.bc, r.de, r.hl, r.sp)
}
