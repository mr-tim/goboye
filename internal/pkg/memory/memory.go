package memory

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

const bootRomRegisterAddr uint16 = 0xff50
const bootRomDisabledValue byte = 0x01
const bootRomEnabledValue byte = 0x00

type MemoryMap interface {
	LoadRomImage(filename string) error
	ReadByte(addr uint16) byte
	WriteByte(addr uint16, value byte)
	ReadU16(addr uint16) uint16
	WriteU16(addr, value uint16)
	GetRoRegister(addr uint16) RoRegister
	GetRwRegister(addr uint16) RwRegister
	ReadAll() []byte
	GetBootRomRegister() *BootRomRegister
}

type BootRomRegister struct {
	// TODO: need to invert this relationship - memoryMap should point at register rather than other way round
	m *memoryMap
}

func (brr *BootRomRegister) IsBootRomPageDisabled() bool {
	return brr.m.bootRomPageDisabled()
}

func (brr *BootRomRegister) SetBootRomPageDisabled(disabled bool) {
	value := bootRomEnabledValue
	if disabled {
		value = bootRomDisabledValue
	}
	brr.m.WriteByte(bootRomRegisterAddr, value)
}
