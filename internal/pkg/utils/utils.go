package utils

const CPU_CYCLES_PER_SECOND = 4194304

func IsBitSet(b byte, index byte) bool {
	mask := uint8(0x01 << index)
	return b&mask > 0
}

func SetBit(b byte, index byte) byte {
	return b | uint8(0x01<<index)
}

func UnsetBit(b byte, index byte) byte {
	mask := uint8(0xFF) - uint8(0x01<<index)
	return b & mask
}
