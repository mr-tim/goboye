package memory

import (
	"io"
	"os"
)

type memoryMap struct {
	mem []byte
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
	if addr < 0x0100 && !m.bootRomPageDisabled() {
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
	return m.ReadByte(bootRomRegisterAddr) == 0x01
}

func (m *memoryMap) ReadAll() []byte {
	result := make([]byte, 0xffff)
	copy(result, m.mem)
	if !m.bootRomPageDisabled() {
		copy(result[:0xff], bootRom)
	}
	return result
}

func (m *memoryMap) GetRoRegister(addr uint16) RoRegister {
	return &RoRegisterAtAddr{memoryMap: m, addr: addr}
}

func (m *memoryMap) GetRwRegister(addr uint16) RwRegister {
	return &RwRegisterAtAddr{roRegister: RoRegisterAtAddr{memoryMap: m, addr: addr}}
}

func (m *memoryMap) GetBootRomRegister() *BootRomRegister {
	return &BootRomRegister{m: m}
}
