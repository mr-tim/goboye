package goboye

import (
	"encoding/base64"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/debugger/recorder"
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"image"
	"log"
)

type Emulator struct {
	memory      *memory.Controller
	processor   cpu.Processor
	display     display.Display
	breakpoints map[uint16]bool
	recorder    *recorder.Recorder
}

func NewEmulator() *Emulator {
	return &(Emulator{
		breakpoints: make(map[uint16]bool),
		recorder:    recorder.NewRecorder(100),
	})
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
	e.recorder.TakeSnapshot(e.processor, e.memory)
	c := e.processor.DoNextInstruction()
	e.display.Update(c)
	return c
}

func (e *Emulator) StepFrame() {
	e.ContinueDebugging(true)
}

func (e *Emulator) ContinueDebugging(stopOnFrame bool) {
	stepCount := 0
	cycleCount := 0

	defer func() {
		r := recover()
		for _, s := range e.recorder.GetSnapshots() {
			log.Printf("%s\n", s)
		}
		log.Printf("%d steps", stepCount)
		if r != nil {
			panic(r)
		}
	}()

	for {
		cycleCount += int(e.Step())
		if _, isBreakpoint := e.breakpoints[e.processor.GetRegisterPair(cpu.RegisterPairPC)]; isBreakpoint {
			break
		}
		if stopOnFrame && cycleCount >= display.CYCLES_PER_FRAME {
			break
		}
		stepCount += 1
	}

}

func (e *Emulator) AddBreakpoint(addr uint16) {
	e.breakpoints[addr] = true
}

func (e *Emulator) RemoveBreakpoint(addr uint16) {
	delete(e.breakpoints, addr)
}

func (e *Emulator) MemoryBase64() string {
	mem := e.memory.ReadAll()
	return base64.StdEncoding.EncodeToString(mem)
}

func (e *Emulator) GetBreakpoints() []uint16 {
	breakpoints := make([]uint16, 0)
	for k := range e.breakpoints {
		breakpoints = append(breakpoints, k)
	}
	return breakpoints
}

func (e *Emulator) DebugRender() image.Image {
	return e.display.DebugRenderMemory()
}
