package memory

import (
	"io"
	"os"
)

type memoryMap struct {
	mem []byte
}

func NewMemoryMapWithBytes(bytes []byte) *memoryMap {
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
	return m.mem[addr]
}

func (m *memoryMap) WriteByte(addr uint16, value byte) {
	m.mem[addr] = value
}

func (m *memoryMap) ReadAll() []byte {
	result := make([]byte, 0xffff)
	copy(result, m.mem)
	return result
}

func (m *memoryMap) GetRoRegister(addr uint16) RoRegister {
	return &RoRegisterAtAddr{memoryMap: m, addr: addr}
}

func (m *memoryMap) GetRwRegister(addr uint16) RwRegister {
	return &RwRegisterAtAddr{roRegister: RoRegisterAtAddr{memoryMap: m, addr: addr}}
}
