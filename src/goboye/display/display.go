package display

import "github.com/veandco/go-sdl2/sdl"

type Display interface {
	Destroy()
}

type SdlDisplay struct {
	window *sdl.Window
}

func (d SdlDisplay) Destroy() {
	d.window.Destroy()
}

func NewDisplay() (Display, error) {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		160, 144, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		return nil, err
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff00ff)
	window.UpdateSurface()

	return SdlDisplay{ window: window }, nil
}