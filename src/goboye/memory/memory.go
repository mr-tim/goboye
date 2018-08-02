package memory

import (
	"os"
	"io"
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

const ROM_SIZE = 0x8000
const MEM_SIZE = 0xFFFF

type memoryMap struct {
	mem []byte
}

func (m *memoryMap) loadRomImage(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	romBytes := make([]byte, ROM_SIZE)
	io.ReadAtLeast(f, romBytes, ROM_SIZE)
	m.mem = append(romBytes[:ROM_SIZE], m.mem[ROM_SIZE:]...)
	return nil
}