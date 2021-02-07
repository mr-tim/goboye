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
const COLS = 160
const ROWS = 144
const VBLANK_ROWS = 10
const TOTAL_ROWS = ROWS + VBLANK_ROWS
const CYCLES_PER_LINE = CYCLES_PER_FRAME / TOTAL_ROWS

func NewDisplay(m *memory.Controller) Display {
	return Display{
		m: m,
	}
}

type Display struct {
	m         *memory.Controller
	cycles    int
	bgPalette color.Palette
	pal0      color.Palette
	pal1      color.Palette
	bgChars   []image.PalettedImage
	pal0Chars []image.PalettedImage
	pal1Chars []image.PalettedImage
	oams      []Oam
}

var Shade0 = color.RGBA{R: 0x9b, G: 0xbc, B: 0x0f, A: 0xff}
var Shade1 = color.RGBA{R: 0x8b, G: 0xac, B: 0x0f, A: 0xff}
var Shade2 = color.RGBA{R: 0x30, G: 0x62, B: 0x30, A: 0xff}
var Shade3 = color.RGBA{R: 0x0f, G: 0x38, B: 0x0f, A: 0xff}

var colors = [4]color.RGBA{Shade0, Shade1, Shade2, Shade3}

type Oam struct {
	X      byte
	Y      byte
	CharID byte
	Attrs  CharAttrs
}

func (d *Display) DebugRenderMemory() image.Image {
	bounds := image.Rect(0, 0, 256, 256)

	d.bgPalette = decodePalette(d.m.BGP.Read(), false)

	// data for characters
	bgCharArea := d.m.LCDCFlags.GetBgCharArea()
	// render the characters into images
	charCount := 256
	rowsPerChar := 8
	addrForChar := func(char byte) uint16 {
		return bgCharArea.Address(char)
	}
	d.bgChars = renderChars(charCount, rowsPerChar, d.bgPalette, addrForChar, d)

	// position of character codes
	bgCodeArea := d.m.LCDCFlags.GetBgCodeArea()
	offset := bgCodeArea.StartAddress()
	p := image.NewPaletted(bounds, d.bgPalette)
	for i := 0; i < 1024; i++ {
		tileX := i % 32
		tileY := i / 32
		charCode := d.m.ReadByte(offset + uint16(i))
		charImg := d.bgChars[charCode]
		draw.Draw(p, image.Rect(tileX*8, tileY*8, (tileX+1)*8, (tileY+1)*8), charImg, image.Point{}, draw.Src)
	}

	if d.m.LCDCFlags.IsObjFlag() {
		// render objs
		charCount := 256
		rowsPerChar := 8
		if d.m.LCDCFlags.IsDoubleObjTiles() {
			charCount = 128
			rowsPerChar = 16
		}
		addrForChar := func(char byte) uint16 {
			return 0x8000 + uint16(char)*uint16(rowsPerChar)*2
		}

		d.pal0 = decodePalette(d.m.OBP0.Read(), true)
		d.pal0Chars = renderChars(charCount, rowsPerChar, d.pal0, addrForChar, d)

		d.pal1 = decodePalette(d.m.OBP1.Read(), true)
		d.pal1Chars = renderChars(charCount, rowsPerChar, d.pal1, addrForChar, d)

		d.oams = make([]Oam, 0)

		for objIdx := 0; objIdx < 40; objIdx += 1 {
			offset := uint16(0xFE00 + objIdx*4)
			oam := Oam{
				X: d.m.ReadByte(offset),
				Y: d.m.ReadByte(offset + 1),
				CharID: d.m.ReadByte(offset + 2),
				Attrs: CharAttrs(d.m.ReadByte(offset + 3)),
			}
			d.oams = append(d.oams, oam)

			left := int(oam.X - 8)
			top := int(oam.Y - 10)
			right := left + 8
			bottom := top + rowsPerChar

			if oam.Attrs.HorizontalFlip() {
				left, right = right, left
			}
			if oam.Attrs.VerticalFlip() {
				top, bottom = bottom, top
			}

			if oam.CharID == 0x88 {
				fmt.Printf("char 88: %#v\n", oam)
			}

			char := d.pal0Chars[oam.CharID]
			if oam.Attrs.IsPal1() {
				char = d.pal1Chars[oam.CharID]
			}

			draw.Draw(p, image.Rect(left, top, right, bottom), char, image.Point{}, draw.Src)
		}
	}

	scx := int(d.m.SCX.Read())
	scy := int(d.m.SCY.Read())
	window := image.Rect(scx, scy, scx+COLS, scy+ROWS)

	return p.SubImage(window)
}

func (d *Display) OutputChars() {
	saveChars(d.bgChars, "../chars/bg_")
	saveChars(d.pal0Chars, "../chars/pal0_")
	saveChars(d.pal1Chars, "../chars/pal1_")
}

func saveChars(chars []image.PalettedImage, prefix string) {
	for i, c := range chars {
		filename := fmt.Sprintf("%s%02d.png", prefix, i)
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, c)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func decodePalette(palDefinition byte, isObj bool) color.Palette {
	palette := make(color.Palette, 4)
	for i := 0; i < 4; i++ {
		idx := (palDefinition >> byte(2*i)) & 0x03
		if isObj && idx == 0 {
			palette[i] = color.Transparent
		} else {
			palette[i] = colors[idx]
		}
	}
	return palette
}

func renderChars(charCount int, rowsPerChar int, palette color.Palette, addrForChar func(byte) uint16,
	d *Display) []image.PalettedImage {
	var chars = make([]image.PalettedImage, charCount)
	charBounds := image.Rect(0, 0, 8, rowsPerChar)
	for charId := 0; charId < charCount; charId++ {
		charImage := image.NewPaletted(charBounds, palette)
		charAddr := addrForChar(byte(charId))
		for y := 0; y < rowsPerChar; y++ {
			addr := charAddr + uint16(2*y)
			cols := decodeRow(d.m.ReadU16(addr))
			for x := 0; x < 8; x++ {
				charImage.SetColorIndex(x, y, cols[x])
			}
		}
		chars[charId] = charImage
	}
	return chars
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

			// TODO: should this be ROWS-1?
			if d.m.LY.Read() == ROWS {
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
		// first byte (high) is low shade bit
		// second byte (lower) is high shade bit
		low := uint8(((0x0001 << shift) & rowData) >> shift)
		high := uint8(((0x0100 << shift) & rowData) >> (shift + 7))
		result[col] = low + high
	}

	return result
}

type CharAttrs byte

func (a CharAttrs) IsPal1() bool {
	return utils.IsBitSet(byte(a), 4)
}

func (a CharAttrs) HorizontalFlip() bool {
	return utils.IsBitSet(byte(a), 5)
}

func (a CharAttrs) VerticalFlip() bool {
	return utils.IsBitSet(byte(a), 6)
}

func (a CharAttrs) BgPriority() bool {
	return utils.IsBitSet(byte(a), 7)
}
