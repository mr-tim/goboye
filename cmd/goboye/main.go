package main

import (
	"flag"
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
	"github.com/mr-tim/goboye/internal/pkg/goboye/button"
	"github.com/mr-tim/goboye/internal/pkg/goboye/ui"
	"github.com/pkg/profile"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

var (
	rom        = flag.String("rom", "", "ROM to run")
	profileCpu = flag.Bool("profileCpu", false, "Profile CPU")
	profileMem = flag.Bool("profileMem", false, "Profile memory")
)

func main() {
	flag.Parse()

	if *rom == "" {
		panic("Please specify a ROM to run")
	}

	if *profileCpu {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if *profileMem {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}

	emulator := goboye.NewEmulator()
	emulator.LoadRomImage(*rom)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	d, err := ui.NewSdlUi()
	if err != nil {
		panic(err)
	}
	defer d.Destroy()

	running := true
	showFrameRate := false
	frameCount := 0

	milliSecondsPerFrame := 1000 / 60

	for running {
		eventStart := time.Now()
		start := int(sdl.GetTicks())
		// handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.KeyboardEvent:
				keyEvent := event.(*sdl.KeyboardEvent)
				btn := buttonMapping(keyEvent)
				switch {
				case btn == button.Quit && keyEvent.Type == sdl.KEYUP:
					running = false
					sdl.Quit()
				case btn == button.Frames && keyEvent.Type == sdl.KEYUP:
					showFrameRate = !showFrameRate
				case btn.IsJoypad():
					emulator.SetButtonState(btn, keyEvent.State == sdl.PRESSED)
				}
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		if !running {
			continue
		}
		eventTime := time.Since(eventStart)
		eventStart = time.Now()

		// Run a single frame's worth of ops
		emulator.StepFrame()

		cpuTime := time.Since(eventStart)
		eventStart = time.Now()

		//redraw
		i := emulator.DebugRender()
		renderTime := time.Since(eventStart)
		eventStart = time.Now()

		d.UpdateScreen(i)

		displayTime := time.Since(eventStart)

		frameTime := int(sdl.GetTicks()) - start
		sleepTime := milliSecondsPerFrame - frameTime

		frameCount = (frameCount + 1) % display.FRAMES_PER_SECOND

		if sleepTime < 0 {
			continue
		} else if sleepTime > 0 {
			if showFrameRate && frameCount == 0 {
				log.Printf("events: %4d\tcpu: %4d\trender: %4d\tdisplay: %4d",
					eventTime.Microseconds(),
					cpuTime.Microseconds(),
					renderTime.Microseconds(),
					displayTime.Microseconds())
			}

			sdl.Delay(uint32(sleepTime))
		}
	}
}

func buttonMapping(ke *sdl.KeyboardEvent) button.Button {
	switch ke.Keysym.Sym {
	// "meta" buttons
	case sdl.K_ESCAPE:
		return button.Quit
	case sdl.K_f:
		return button.Frames
	// joypad
	case sdl.K_a:
		return button.Left
	case sdl.K_d:
		return button.Right
	case sdl.K_s:
		return button.Down
	case sdl.K_w:
		return button.Up
	case sdl.K_p:
		return button.A
	case sdl.K_o:
		return button.B
	case sdl.K_n:
		return button.Select
	case sdl.K_m:
		return button.Start
	}
	return button.Unbound
}
