package goboye

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/debugger/recorder"
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"image"
	"log"
	"os"
)

type Emulator struct {
	memory      *memory.Controller
	processor   cpu.Processor
	display     display.Display
	breakpoints [0xFFFF]bool
	recorder    *recorder.Recorder
	debug       bool
}

func NewEmulator() *Emulator {
	breakpoints := loadBreakpoints()
	return &(Emulator{
		breakpoints: breakpoints,
		recorder:    recorder.NewRecorder(),
	})
}

func loadBreakpoints() [0xFFFF]bool {
	f, err := os.Open("breakpoints.txt")

	breakpoints := [0xFFFF]bool{}

	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	} else {
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := s.Text()
			bs, err := hex.DecodeString(line[2:])
			if err == nil && len(bs) == 2 {
				addr := uint16(0)
				addr |= uint16(bs[0]) << 8
				addr |= uint16(bs[1])
				breakpoints[addr] = true
			}
		}
	}

	return breakpoints
}

func (e *Emulator) LoadRomImage(filename string) {
	m := memory.NewController()
	e.memory = &m

	log.Printf("Loading rom: %s", filename)
	err := e.memory.LoadRomImage(filename)
	if err != nil {
		panic(err)
	}

	e.processor = cpu.NewProcessor(e.memory)
	e.display = display.NewDisplay(e.memory)
}

func (e *Emulator) GetDisassembler() cpu.Disassembler {
	return cpu.NewDisassembler(e.memory)
}

func (e *Emulator) GetPC() uint16 {
	return e.GetRegisterPair(cpu.RegisterPairPC)
}

func (e *Emulator) GetRegisterPair(registerPair cpu.RegisterPair) uint16 {
	return e.processor.GetRegisterPair(registerPair)
}

func (e *Emulator) GetFlagValue(flagName cpu.OpResultFlag) bool {
	return e.processor.GetFlagValue(flagName)
}

func (e *Emulator) Step() uint8 {
	if e.processor.IsStopped() || e.processor.IsHalted() {
		return 0
	}
	e.recorder.TakeSnapshot(e.processor, e.memory)
	pc := e.GetPC()
	c := e.processor.DoNextInstruction()
	// break on infinite loops (PC isn't advancing because of JrN -1
	if pc == e.GetPC() {
		// infinite loop
		e.breakpoints[e.GetPC()] = true
	}
	e.display.Update(c)
	return c
}

func (e *Emulator) StepFrame() {
	e.ContinueDebugging(true)
}

func (e *Emulator) ContinueDebugging(stopOnFrame bool) {
	stepCount := 0
	cycleCount := 0

	if e.debug {
		defer func() {
			r := recover()
			if e.recorder.IsEnabled() {
				for _, s := range e.recorder.GetSnapshots() {
					log.Printf("%s\n", s)
				}
			}
			log.Printf("%d steps", stepCount)
			if r != nil {
				panic(r)
			}
		}()
	}

	for {
		cycleCount += int(e.Step())
		if e.processor.IsHalted() || e.processor.IsStopped() {
			break
		}

		if e.debug {
			pc := e.processor.GetRegisterPair(cpu.RegisterPairPC)

			if e.breakpoints[pc] {
				fmt.Printf("At %04X\n", pc)
				break
			}
		}

		if stopOnFrame {
			if e.memory.LCDCFlags.IsLCDEnabled() {
				if cycleCount >= display.CYCLES_PER_FRAME && e.memory.LYC.Read() == 0 {
					break
				}
			}
		}
		stepCount += 1
	}
}

func (e *Emulator) AddBreakpoint(addr uint16) {
	e.breakpoints[addr] = true
}

func (e *Emulator) RemoveBreakpoint(addr uint16) {
	e.breakpoints[addr] = false
}

func (e *Emulator) MemoryBase64() string {
	mem := e.memory.ReadAll()
	return base64.StdEncoding.EncodeToString(mem)
}

func (e *Emulator) GetBreakpoints() []uint16 {
	breakpoints := make([]uint16, 0)
	for i := range e.breakpoints {
		if e.breakpoints[i] {
			breakpoints = append(breakpoints, uint16(i))
		}
	}
	return breakpoints
}

func (e *Emulator) DebugRender() image.Image {
	return e.display.DebugRenderMemory()
}

func (e *Emulator) SerialOutput() string {
	return e.memory.SerialOutput
}

func (e *Emulator) SetDebug(debug bool) {
	e.debug = debug
}
