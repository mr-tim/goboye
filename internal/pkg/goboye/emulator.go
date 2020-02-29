package goboye

import (
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"log"
)

type Emulator struct {
	memoryMap memory.MemoryMap
	processor cpu.Processor
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
