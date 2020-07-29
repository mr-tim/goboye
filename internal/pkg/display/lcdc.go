package display

import "github.com/mr-tim/goboye/internal/pkg/memory"

type BgCodeArea byte

const (
	BgCodeArea1 BgCodeArea = 1 //0x9800-0x9BFF
	BgCodeArea2 BgCodeArea = 2 //0x9C00-0x9FFF
)

func (a BgCodeArea) StartAddress() uint16 {
	if a == BgCodeArea1 {
		return 0x9800
	} else if a == BgCodeArea2 {
		return 0x9C00
	} else {
		panic("invalid bg code area specified!")
	}
}

type BgCharDataArea byte

const (
	BgCharArea1 BgCharDataArea = 1 // 0x8800-0x97FF
	BgCharArea2 BgCharDataArea = 2 // 0x8000-0x8FFF
)

func (a BgCharDataArea) Address(id byte) uint16 {
	if a == BgCharArea1 && id < 0x80 {
		// for BgCharArea1, 0x00-0x7F are in 0x9000 (as is BgCharArea2)
		// 0x80+ is in 0x8800-0x8FFF
		return 0x9000 + uint16(id)*0x0010
	} else {
		return 0x8000 + uint16(id)*0x0010
	}
}

type WindowCodeArea byte

const (
	WindowCodeArea1 WindowCodeArea = 1 //0x9800-0x9BFF
	WindowCodeArea2 WindowCodeArea = 2 //0x9C00-0x9FFF
)

type LCDCFlags struct {
	r memory.RwRegister
}

/*
	Bits:
	0: Bg display off (0) or on (1). Always on for CGB
	1: OBJ flag off (0) or on (1)
	2: Obj composition 8x8 (0) or 8x16 (1)
	3: BG code area selection 0x9800-0x9BFF (0) or 0x9C00-0x9FFF (1)
	4: BG char data selection 0x8800-0x97FF (0) or 0x8000-0x8FFF (1)
	5: Windowing flag off (0) or on (1)
	6: Window code area 0x9800-0x9BFF (0) or 0x9C00-0x9FFF (1)
	7: LCD controller op stop flag off (0) or on (1)
*/
func (f LCDCFlags) IsBgDisplay() bool {
	return f.r.IsBitSet(0)
}

func (f LCDCFlags) IsObjFlag() bool {
	return f.r.IsBitSet(1)
}

func (f LCDCFlags) IsDoubleObjTiles() bool {
	return f.r.IsBitSet(2)
}

func (f LCDCFlags) GetBgCodeArea() BgCodeArea {
	if f.r.IsBitSet(3) {
		return BgCodeArea2
	} else {
		return BgCodeArea1
	}
}

func (f LCDCFlags) GetBgCharArea() BgCharDataArea {
	if f.r.IsBitSet(4) {
		return BgCharArea2
	} else {
		return BgCharArea1
	}
}

func (f LCDCFlags) IsWindowingFlagSet() bool {
	return f.r.IsBitSet(5)
}

func (f LCDCFlags) GetWindowCodeArea() WindowCodeArea {
	if f.r.IsBitSet(6) {
		return WindowCodeArea2
	} else {
		return WindowCodeArea1
	}
}

func (f LCDCFlags) IsOpStopped() bool {
	return f.r.IsBitSet(7)
}
