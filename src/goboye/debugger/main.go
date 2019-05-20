package main

import (
	"fmt"
	"goboye/cpu"
	"goboye/memory"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

type Disassembler struct {
	m memory.MemoryMap
	pos uint16
}

func (d *Disassembler) SetPos(pos uint16) {
	d.pos = pos
}

func (d *Disassembler) GetNextInstruction() (uint16, cpu.Opcode, string) {
	addr := d.pos
	opcodeByte := d.m.ReadByte(d.pos)
	d.pos += 1
	o := cpu.LookupOpcode(opcodeByte)

	if o.Code() == cpu.OpcodeExtOps.Code() {
		// load the extended code
		opcodeByte = d.m.ReadByte(d.pos)
		d.pos += 1
		o = cpu.LookupExtOpcode(opcodeByte)
	}

	payloadStr := ""
	argWidth := o.PayloadLength()
	if argWidth == 1 {
		payloadStr = fmt.Sprintf("0x%02x", d.m.ReadByte(d.pos))
	} else if argWidth == 2 {
		payloadStr = fmt.Sprintf("0x%04x", d.m.ReadU16(d.pos))
	}
	d.pos += uint16(argWidth)

	return addr, &o, payloadStr
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	buf := make([]byte, memory.MEM_SIZE)
	mm := memory.NewMemoryMapWithBytes(buf)

	if len(os.Args) != 2 {
		fmt.Printf("Please specify a rom to run.\n")
		os.Exit(1)
	}

	e := mm.LoadRomImage(os.Args[1])
	if e != nil {
		panic(e)
	}

	p := cpu.NewProcessor(mm)
	pc := p.GetRegisterPair(cpu.RegisterPairPC)


	da := Disassembler{m: mm, pos: pc}

	dw := &DisassemblyWidget{
		da: da,
		currentPc: 0x0000,
	}
	g.SetManager(dw)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	increment := func (gui *gocui.Gui, view *gocui.View) error {
		dw.SetPc(dw.currentPc + 1)
		return nil
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, increment); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

type DisassemblyWidget struct {
	da Disassembler
	currentPc uint16
}

func (d *DisassemblyWidget) Layout(g *gocui.Gui) error {
	_, maxY := g.Size()
	v, err := g.SetView("disassembly", 0, 0, 30, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Disassembly"
	v.Clear()
	// TODO: Handle scrolling in the disassembly view
	d.da.SetPos(0)
	for i := 0; i<maxY-1; i++ {
		addr, o, payload := d.da.GetNextInstruction()
		pointer := " "
		if addr == d.currentPc {
			pointer = ">"
		}
		fmt.Fprintf(v, " %s 0x%04x %s\n", pointer, addr, o.DisassemblyWithArg(payload))
	}
	return nil
}

func (d *DisassemblyWidget) SetPc(newPc uint16) {
	d.currentPc = newPc
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}