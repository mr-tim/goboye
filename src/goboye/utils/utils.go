package utils

func IsBitSet(b byte, index byte) bool {
	mask := uint8(0x01 << index)
	return b&mask > 0
}
