package register

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBgCharAddressArea1(t *testing.T) {
	assert.Equal(t, uint16(0x9000), BgCharArea1.Address(0x00))
	assert.Equal(t, uint16(0x9010), BgCharArea1.Address(0x01))
	assert.Equal(t, uint16(0x9020), BgCharArea1.Address(0x02))
	assert.Equal(t, uint16(0x97F0), BgCharArea1.Address(0x7F))
	assert.Equal(t, uint16(0x8800), BgCharArea1.Address(0x80))
	assert.Equal(t, uint16(0x8820), BgCharArea1.Address(0x82))
	assert.Equal(t, uint16(0x8FF0), BgCharArea1.Address(0xFF))
}

func TestBgCharAddressArea2(t *testing.T) {
	assert.Equal(t, uint16(0x8000), BgCharArea2.Address(0x00))
	assert.Equal(t, uint16(0x8010), BgCharArea2.Address(0x01))
	assert.Equal(t, uint16(0x8020), BgCharArea2.Address(0x02))
	assert.Equal(t, uint16(0x87F0), BgCharArea2.Address(0x7F))
	assert.Equal(t, uint16(0x8800), BgCharArea2.Address(0x80))
	assert.Equal(t, uint16(0x8820), BgCharArea2.Address(0x82))
	assert.Equal(t, uint16(0x8FF0), BgCharArea2.Address(0xFF))
}
