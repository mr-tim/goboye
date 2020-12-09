package utils

const CPU_CYCLES_PER_SECOND = 4194304

func IsBitSet(b byte, index byte) bool {
	mask := uint8(0x01 << index)
	return b&mask > 0
}

func UnsetBit(b byte, index byte) byte {
	mask := uint8(0xFF) - uint8(0x01<<index)
	return b & mask
}
