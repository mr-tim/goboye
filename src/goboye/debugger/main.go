package main

import (
	"fmt"
	"goboye/cpu"
	"goboye/memory"
	"os"
)

func main() {
	fmt.Println("Running...")

	buf := make([]byte, memory.MEM_SIZE)
	mm := memory.NewMemoryMapWithBytes(buf)
	e := mm.LoadRomImage(os.Args[1])
	if e != nil {
		panic(e)
	}

	p := cpu.NewProcessor(mm)
	for i := 0; i < 20000; i++ {
		o := p.NextInstruction()
		pc := p.GetRegisterPair(cpu.RegisterPairPC)
		argWidth := o.PayloadLength()
		if argWidth == 1 {
			arg := mm.ReadByte(pc + 1)
			fmt.Printf("0x%04x: %s %s 0x%02x ", pc, p.DebugRegisters(), o.Disassembly(), arg)
		} else if argWidth == 2 {
			arg := mm.ReadU16(pc + 1)
			fmt.Printf("0x%04x: %s %s 0x%04x ", pc, p.DebugRegisters(), o.Disassembly(), arg)
		} else {
			fmt.Printf("0x%04x: %s %s ", pc, p.DebugRegisters(), o.Disassembly())
		}
		fmt.Println()
		p.DoNextInstruction()
	}
}
