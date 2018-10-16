package display

/*
	0x8000-0x9FFF - display memory
		Dot data - 16 bytes per character
			2 bytes per row - lower then higher

			eg: 0 1 0 0 1 1 1 0
				1 0 0 0 1 0 1 1
				2 1 0 0 3 1 3 2

			0x8000-0x87FF - OBJ dot data 0-127
			0x8800-0x8FFF - OBJ dot data 128-255
						  - BG data -128 to -1
			0x9000-0x97FF - BG data 0-127

			0x9800-0x9BFF - tile map 0
			0x9C00-0x9FFF - tile map 1

	tile map 0 or 1 is determined by bit 3 of LCDC register - 0xFF40

	Registers:
		LCDC	0xFF40
			Bits:
			0: Bg display off (0) or on (1). Always on for CGB
			1: OBJ flag off (0) or on (1)
			2: Obj composition 8x8 (0) or 8x16 (1)
			3: BG code area selection 0x9800-0x9BFF (0) or 0x9C00-0x9FFF (1)
			4: BG char data selection 0x8800-0x97FF (0) or 0x8000-0x8FFF (1)
			5: Windowing flag off (0) or on (1)
			6: Window code area 0x9800-0x9BFF (0) or 0x9C00-0x9FFF (1)
			7: LCD controller op stop flag off (0) or on (1)

		STAT	0xFF41
			Bits:
			0, 1: Mode flag
				00 Enable cpu access to all display ram
				01 In vertical blanking period (cpu has access approx 1ms to display ram)
				10 Searching OAM RAM
				11 Transferring data to LCD driver
			2: Match flag LYC = LCDC LY (0) or LYC = LCDC LY (1) [not a typo...]
			Interrupt selection:
				3: Mode 00 selection
				4: Mode 01 selection 0: not selected
				5: Mode 10 selection 1: selected
				6: LYC = LY matching selection

		SCY		0xFF42 - Scroll Y
		SCX		0xFF43 - Scroll X

		LY		0xFF44 - LCDC (read only)
			LY indicates line bening transferred to lcd driver
			Values 0-153, 144-153 indicate v-blank period
			When bit 7 of LCDC is 1, writing 1 again does not change value LY
			Writing 0 to LCDC(7) when it's 1 resets LY to 0

		LYC		0xFF45 - LYC
			Compare to LY - if matches, STAT(2) is set

		BCPS	0xFF68 - BCPS - Specifies BG Write
			0: Specifies H (1) or L (0)
			1, 2: Palette data no
			3, 4, 5: Palette no
			6: XXX
			7: 	1: with each write, specifies next palette
				0: bits 0-5 fixed

		BCPD	0xFF69 - BG Write data

		OCPS	0xFF6A - Specifies OBJ Write
			0: Specified H (1) or L (0)
			1, 2: Palette data no
			3, 4, 5: Palette no
			7: 	1: with each write, specifies next palette
				0: bits 0-5 fixed

		OCPD	0xFF6B - OBJ Write data

		WY		0xFF4A - Window y-coordinate
						0 <= WY <= 143

		WX		0XFF4B - Window x-coordinate
						7 <= WX <= 166
						0-6 should not be specified.
						7 corresponds to left edge


	0xFE00-0xFE9F - OAM Registers
		Data for up to 40 objects - 10 can be on the same line
		Structure:
			y-coord (8 bits) (y=10 => object displayed from top edge of screen)
			x-coord (8 bits) (x=6 => object displayed from left edge of screen)
			chr code (8 bits)
			bg/ob priority (1 bit) (0: obj, 1: bg)
			vertical flip (1 bit) (0: normal, 1: flipped)
			horizontal flip (1 bit) (0: normal, 1: flipped)
			dmg mode palette (1 bit)
			character bank (1 bit) (cgb)
			color palette (3 bits) (cgb)

 */

 type BgCodeArea byte

 const (
 	BgCodeArea1 BgCodeArea = 1 //0x9800-0x9BFF
 	BgCodeArea2 BgCodeArea = 2 //0x9C00-0x9FFF
 )

 type BgCharDataArea byte

 const (
 	BgCharArea1 BgCharDataArea = 1 // 0x8800-0x97FF
 	BgCharArea2 BgCharDataArea = 2 // 0x8000-0x8FFF
)

 type WindowCodeArea byte

 const (
 	WindowCodeArea1 WindowCodeArea = 1 //0x9800-0x9BFF
 	WindowCodeArea2 WindowCodeArea = 2 //0x9C00-0x9FFF
 )

 type LCDCFlags byte

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
func (b LCDCFlags) IsBgDisplay() bool {
	return IsBitSet(byte(b), 0)
}

func (b LCDCFlags) IsObjFlag() bool {
	return IsBitSet(byte(b), 1)
}

func (b LCDCFlags) IsDoubleObjTiles() bool {
	return IsBitSet(byte(b), 2)
}

func (b LCDCFlags) GetBgCodeArea() BgCodeArea {
	if IsBitSet(byte(b), 3) {
		return BgCodeArea2
	} else {
		return BgCodeArea1
	}
}

func (b LCDCFlags) GetBgCharArea() BgCharDataArea {
	if IsBitSet(byte(b), 4) {
		return BgCharArea2
	} else {
		return BgCharArea1
	}
}

func (b LCDCFlags) IsWindowingFlagSet() bool {
	return IsBitSet(byte(b), 5)
}

func (b LCDCFlags) GetWindowCodeArea() WindowCodeArea {
	if IsBitSet(byte(b), 6) {
		return WindowCodeArea2
	} else {
		return WindowCodeArea1
	}
}

func (b LCDCFlags) IsOpStopped() bool {
	return IsBitSet(byte(b), 7)
}


func IsBitSet(b byte, index byte) bool {
	mask := uint8(0x01 << index)
	return b & mask > 0
}
