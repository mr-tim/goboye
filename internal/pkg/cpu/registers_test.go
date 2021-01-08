package cpu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetRegister(t *testing.T) {
	r := Registers{}

	rs := []register{RegisterA, RegisterB, RegisterC, RegisterD, RegisterE, RegisterF, RegisterH, RegisterL}

	for _, reg := range rs {
		for i := 0; i < 255; i++ {
			v := uint8(i)
			r.setRegister(reg, v)
			if reg == RegisterF {
				v &= 0xF0
			}
			assert.Equal(t, v, r.getRegister(reg))
		}
	}
}

func TestSetTwoAdjacentRegisters(t *testing.T) {
	r := Registers{}
	r.setRegister(RegisterA, 0xA5)
	assert.Equal(t, uint8(0xA5), r.getRegister(RegisterA))
	r.setRegister(RegisterF, 0x6D)
	assert.Equal(t, uint8(0x60), r.getRegister(RegisterF))
	assert.Equal(t, uint8(0xA5), r.getRegister(RegisterA))
	assert.Equal(t, uint16(0xA560), r.getRegisterPair(RegisterPairAF))
}

func TestSetRegisterPair(t *testing.T) {
	r := Registers{}
	rs := []RegisterPair{RegisterPairAF, RegisterPairBC, RegisterPairDE, RegisterPairHL, RegisterPairSP}
	vs := []uint16{123, 12345, 65355}

	for _, reg := range rs {
		for _, v := range vs {
			r.setRegisterPair(reg, v)
			expected := v
			if reg == RegisterPairAF {
				expected &= 0xFFF0
			}
			assert.Equal(t, expected, r.getRegisterPair(reg))
		}
	}
}

func TestRegisterPairEndianness(t *testing.T) {
	r := &Registers{}
	rs := []RegisterPair{RegisterPairAF, RegisterPairBC, RegisterPairDE, RegisterPairHL}
	rps := [][]register{{RegisterA, RegisterF}, {RegisterB, RegisterC}, {RegisterD, RegisterE}, {RegisterH, RegisterL}}
	v := uint16(0x4698)

	for i, reg := range rs {
		r.setRegisterPair(reg, v)
		expected := v
		if reg == RegisterPairAF {
			expected &= 0xFFF0
		}
		assert.Equal(t, expected, r.getRegisterPair(reg))
		rp := rps[i]
		assert.Equal(t, uint8(0x46), r.getRegister(rp[0]))
		regExpected := uint8(0x98)
		if rp[1] == RegisterF {
			regExpected &= 0xF0
		}
		assert.Equal(t, regExpected, r.getRegister(rp[1]))
	}
}
