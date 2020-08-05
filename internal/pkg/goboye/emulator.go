package goboye

import (
	"encoding/base64"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/display"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"image"
	"log"
)

type Emulator struct {
	memoryMap   memory.MemoryMap
	processor   cpu.Processor
	display     display.Display
	breakpoints map[uint16]bool
}

const CYCLES_PER_SECOND = 4194304
const FRAMES_PER_SECOND = 60
const CYCLES_PER_FRAME = CYCLES_PER_SECOND/FRAMES_PER_SECOND

func NewEmulator() *Emulator {
	return &(Emulator{
		breakpoints: make(map[uint16]bool),
	})
}

func (e *Emulator) LoadRomImage(filename string) {
	buf := make([]byte, memory.MEM_SIZE)
	e.memoryMap = memory.NewMemoryMapWithBytes(buf)

	log.Printf("Loading rom: %s", filename)
	err := e.memoryMap.LoadRomImage(filename)
	if err != nil {
		panic(err)
	}

	e.processor = cpu.NewProcessor(e.memoryMap)
	e.display = display.NewDisplay(e.memoryMap)
}

func (e *Emulator) GetDisassembler() *cpu.Disassembler {
	return cpu.NewDisassembler(e.memoryMap)
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

func (e *Emulator) Step() {
	e.processor.DoNextInstruction()
}

func (e *Emulator) StepFrame() {
	cycles := 0
	for cycles < CYCLES_PER_FRAME {
		c := e.processor.DoNextInstruction()
		e.display.Update(c)
	}
}

func (e *Emulator) ContinueDebugging() {
	for {
		e.Step()
		if _, isBreakpoint := e.breakpoints[e.processor.GetRegisterPair(cpu.RegisterPairPC)]; isBreakpoint {
			break
		}
	}
}

func (e *Emulator) AddBreakpoint(addr uint16) {
	e.breakpoints[addr] = true
}

func (e *Emulator) RemoveBreakpoint(addr uint16) {
	delete(e.breakpoints, addr)
}

func (e *Emulator) MemoryBase64() string {
	mem := e.memoryMap.ReadAll()
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
