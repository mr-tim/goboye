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

func TestSetRegisterPair(t *testing.T) {
	r := registers{}

	rs := []registerPair{AF, BC, DE, HL}
	vs := []uint16{123, 12345, 65355}

	for _, reg := range rs {
		for _, v := range vs {
			r.setRegisterPair(reg, v)
			assert.Equal(t, v, r.getRegisterPair(reg))
		}
	}
}
