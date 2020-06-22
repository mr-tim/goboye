package memory

import (
	"io"
	"os"
)

/*
	see p13/14 in programming manual
	0x0000-0x8000 - ROM
		0x0000-0x00FF - destination address for RST and starting address for interrupts
		0x0100-0x014F - ROM area for game name etc
		0x0150-0x7FFF - program area
	0x8000-0x9FFF - ram for LCD display
	0xA000-0xBFFF - expansion ram
	0xC000-0xDFFF - work area ram
	0xE000-0xFDFF - prohibited
	0xFE00-0xFFFF - cpu internal
		0xFE00-0xFE9f - OAM-RAM - sprite attributes
		0xFF00-0xFF7F + 0xFFFF instruction registers etc
		0xFF80 - 0xFFFE - CPU work ram/stack ram
		0xFFFF - Interrupt Enable Register
*/

const ROM_SIZE = 0x08000
const MEM_SIZE = 0x10000

type memoryMap struct {
	mem []byte
}

type MemoryMap interface {
	LoadRomImage(filename string) error
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
	ReadU16(addr uint16) uint16
	WriteU16(addr, value uint16)
}

func NewMemoryMapWithBytes(bytes []byte) MemoryMap {
	m := memoryMap{make([]byte, MEM_SIZE)}
	m.initWithBytes(bytes)
	return &m
}

func (m *memoryMap) LoadRomImage(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	romBytes := make([]byte, ROM_SIZE)
	io.ReadAtLeast(f, romBytes, ROM_SIZE)
	m.initWithBytes(romBytes)
	return nil
}

func (m *memoryMap) initWithBytes(bytes []byte) {
	m.mem = append(bytes[:], m.mem[len(bytes):]...)
}

func (m *memoryMap) ReadByte(addr uint16) byte {
	if addr < 0xFF && !m.bootRomPageDisabled() {
		return bootRom[addr]
	} else {
		return m.mem[addr]
	}
}

func (m *memoryMap) WriteByte(addr uint16, value byte) {
	m.mem[addr] = value
}

func (m *memoryMap) ReadU16(addr uint16) uint16 {
	return (uint16(m.ReadByte(addr+1)) << 8) | uint16(m.ReadByte(addr))
}

func (m *memoryMap) WriteU16(addr, value uint16) {
	l := uint8(0x00FF & value)
	h := uint8((0xFF00 & value) >> 8)
	m.WriteByte(addr, l)
	m.WriteByte(addr+1, h)
}

func (m *memoryMap) bootRomPageDisabled() bool {
	return m.ReadByte(0xff50) == 0x01
}