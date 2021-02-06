package ui

import (
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/veandco/go-sdl2/sdl"
	"image"
)

const SCALE = 4

type Ui interface {
	Destroy()
	UpdateScreen(i image.Image)
}

type SdlUi struct {
	window *sdl.Window
	drawSurface *sdl.Surface
}

func (d SdlUi) Destroy() {
	d.window.Destroy()
	d.drawSurface.Free()
}

func (d SdlUi) UpdateScreen(i image.Image) {
	for x := 0; x < display.COLS; x+=1 {
		for y := 0; y < display.ROWS; y+=1 {
			d.drawSurface.Set(x, y, i.At(x, y))
		}
	}

	windowSurface, err := d.window.GetSurface()
	if err != nil {
		panic(err)
	}
	rect := wholeScreen()
	scaledRect := wholeScreenScaled()
	err = d.drawSurface.BlitScaled(&rect, windowSurface, &scaledRect)
	if err != nil {
		panic(err)
	}

	err = d.window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func NewSdlUi() (Ui, error) {
	window, err := sdl.CreateWindow("goboye", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCALE * display.COLS, SCALE * display.ROWS, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		return nil, err
	}

	rect := wholeScreenScaled()
	c := display.Shade3
	b := uint32(c.A) << 24 | uint32(c.R) << 16 | uint32(c.G) << 8 | uint32(c.B)

	err = surface.FillRect(&rect, b)
	if err != nil {
		panic(err)
	}

	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}

	drawSurface, err := sdl.CreateRGBSurface(0, SCALE * display.COLS, SCALE * display.ROWS, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}

	return SdlUi{
		window:      window,
		drawSurface: drawSurface,
	}, nil
}

func wholeScreen() sdl.Rect {
	return sdl.Rect{W: display.COLS, H: display.ROWS}
}

func wholeScreenScaled() sdl.Rect {
	return sdl.Rect{W: SCALE * display.COLS, H: SCALE * display.ROWS}
}
