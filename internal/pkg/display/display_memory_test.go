package display

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeTile(t *testing.T) {
	assert.Equal(t, [8]uint8{2, 2, 1, 0, 0, 3, 3, 0}, decodeRow(0xC626))
}
