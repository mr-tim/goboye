package main

import (
	"flag"
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/mr-tim/goboye/internal/pkg/goboye"
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
	frameCount := 0

	milliSecondsPerFrame := 1000 / 60

	for running {
		eventStart := time.Now()
		start := int(sdl.GetTicks())
		// handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
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
			if frameCount == 0 {
				log.Printf("PC: %04X", emulator.GetPC())
				log.Printf("events: %s cpu: %s render: %s display: %s", eventTime, cpuTime, renderTime, displayTime)
			}

			sdl.Delay(uint32(sleepTime))
		}
	}
}
