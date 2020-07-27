package main

import (
	sdl2 "github.com/mr-tim/goboye/internal/pkg/display/sdl"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	d, err := sdl2.NewDisplay()
	if err != nil {
		panic(err)
	}
	defer d.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}

		}
	}
}
