package display

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/display/register"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/mr-tim/goboye/internal/pkg/utils"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

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
			2: Match flag LYC != LCDC LY (0) or LYC = LCDC LY (1)
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

const FRAMES_PER_SECOND = 60
const CYCLES_PER_FRAME = utils.CPU_CYCLES_PER_SECOND / FRAMES_PER_SECOND
const ROWS = 144
const VBLANK_ROWS = 10
const TOTAL_ROWS = ROWS + VBLANK_ROWS
const CYCLES_PER_LINE = CYCLES_PER_FRAME / TOTAL_ROWS

const outputBgChars = false

func NewDisplay(m *memory.Controller) Display {
	return Display{
		m: m,
	}
}

type Display struct {
	m      *memory.Controller
	cycles int
}

func (d *Display) DebugRenderMemory() image.Image {
	bounds := image.Rect(0, 0, 256, 256)

	colors := []color.RGBA{
		{R: 0x9b, G: 0xbc, B: 0x0f, A: 0xff},
		{R: 0x8b, G: 0xac, B: 0x0f, A: 0xff},
		{R: 0x30, G: 0x62, B: 0x30, A: 0xff},
		{R: 0x0f, G: 0x38, B: 0x0f, A: 0xff},
	}

	palette := color.Palette{}
	palDefinition := d.m.BGP.Read()
	for i := 0; i < 4; i++ {
		idx := (palDefinition >> byte(2*i)) & 0x03
		palette = append(palette, colors[idx])
	}

	// data for characters
	bgCharArea := d.m.LCDCFlags.GetBgCharArea()
	// render the characters into images
	var bgChars = make([]image.PalettedImage, 0)
	charBounds := image.Rect(0, 0, 8, 8)
	for charId := 0; charId < 256; charId++ {
		charImage := image.NewPaletted(charBounds, palette)
		charAddr := bgCharArea.Address(byte(charId))
		for y := 0; y < 8; y++ {
			addr := charAddr + uint16(2*y)
			// first byte (high) is low shade bit
			// second byte (lower) is high shade bit
			row := decodeRow(d.m.ReadU16(addr))
			for x := 0; x < 8; x++ {
				charImage.SetColorIndex(x, y, row[x])
			}
		}
		bgChars = append(bgChars, charImage)
	}

	if outputBgChars {
		for idx, bgChar := range bgChars {
			err := os.MkdirAll("../chars", 0755)
			if err != nil {
				panic(err)
			}
			f, err := os.Create(fmt.Sprintf("../chars/bgchar%000d.png", idx))
			if err != nil {
				panic(err)
			}
			png.Encode(f, bgChar)
			f.Close()
		}
	}

	// position of character codes
	bgCodeArea := d.m.LCDCFlags.GetBgCodeArea()
	offset := bgCodeArea.StartAddress()
	p := image.NewPaletted(bounds, palette)
	for i := 0; i < 1024; i++ {
		tileX := i % 32
		tileY := int(i / 32)
		charCode := d.m.ReadByte(offset + uint16(i))
		charImg := bgChars[charCode]
		draw.Draw(p, image.Rect(tileX*8, tileY*8, (tileX+1)*8, (tileY+1)*8), charImg, image.Point{}, draw.Src)
	}
	return p
}

func (d *Display) Update(cycles uint8) {
	if d.m.LCDCFlags.IsLCDEnabled() {
		d.cycles += int(cycles)
		for d.cycles > CYCLES_PER_LINE {
			d.m.LY.Write(d.m.LY.Read() + 1)
			value := d.m.LY.Read()
			if value >= TOTAL_ROWS {
				// set it to zero
				d.m.LY.Write(0)
			}

			if d.m.LY.Read() >= ROWS {
				// set v-blank flag
				d.m.StatFlags.SetMode(register.VerticalBlank)
				d.m.InterruptFlags.VBlankInterrupt()
			} else {
				d.m.StatFlags.SetMode(register.EnableCPUAccessToDisplayRAM)
			}

			// draw a line
			d.cycles -= CYCLES_PER_LINE
		}
	}
}

func decodeRow(rowData uint16) [8]uint8 {
	var result [8]uint8
	for col := 0; col < 8; col++ {
		shift := 7 - col
		low := uint8(((0x0001 << shift) & rowData) >> shift)
		high := uint8(((0x0100 << shift) & rowData) >> (shift + 7))
		result[col] = low + high
	}

	return result
}
