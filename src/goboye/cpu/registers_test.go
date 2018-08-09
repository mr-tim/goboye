package cpu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetRegister(t *testing.T) {
	r := registers{}

	rs := []register{A, B, C, D, E, F, H, L}

	for _, reg := range rs {
		for i := 0; i < 255; i++ {
			v := uint8(i)
			r.setRegister(reg, v)
			assert.Equal(t, v, r.getRegister(reg))
		}
	}
}

func TestSetTwoAdjacentRegisters(t *testing.T) {
	r := registers{}
	r.setRegister(A, 0xA5)
	assert.Equal(t, uint8(0xA5), r.getRegister(A))
	r.setRegister(F, 0x6D)
	assert.Equal(t, uint8(0x6D), r.getRegister(F))
	assert.Equal(t, uint8(0xA5), r.getRegister(A))
}

func TestSetRegisterPair(t *testing.T) {
	r := registers{}
	rs := []registerPair{AF, BC, DE, HL, SP}
	vs := []uint16{123, 12345, 65355}

	for _, reg := range rs {
		for _, v := range vs {
			r.setRegisterPair(reg, v)
			assert.Equal(t, v, r.getRegisterPair(reg))
		}
	}
}

func TestRegisterPairEndianness(t *testing.T) {
	r := &registers{}
	rs := []registerPair{AF, BC, DE, HL}
	rps := [][]register{{A, F}, {B, C}, {D, E}, {H, L}}
	v := uint16(0x4698)

	for i, reg := range rs {
		r.setRegisterPair(reg, v)
		assert.Equal(t, v, r.getRegisterPair(reg))
		rp := rps[i]
		assert.Equal(t, uint8(0x46), r.getRegister(rp[0]))
		assert.Equal(t, uint8(0x98), r.getRegister(rp[1]))
	}
}
