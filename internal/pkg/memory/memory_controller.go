package memory

import (
	"github.com/mr-tim/goboye/internal/pkg/display/register"
	"io"
	"os"
)

type Controller struct {
	romImage         memoryMap
	ram              memoryMap
	stack            memoryMap
	ControllerData   controllerRegister
	BootRomRegister  bootRomByteRegister
	LCDCFlags        register.LCDCFlags
	StatFlags        register.StatFlags
	SCY              simpleByteRegister
	SCX              simpleByteRegister
	LY               simpleByteRegister
	LYC              simpleByteRegister
	BGP              simpleByteRegister
	InterruptFlags   InterruptFlagsRegister
	InterruptEnabled InterruptEnabledRegister
}

func NewController() Controller {
	return Controller{
		romImage: memoryMap{make([]byte, ROM_SIZE)},
		ram:      memoryMap{make([]byte, STACK_START-ROM_SIZE)},
		stack:    memoryMap{make([]byte, STACK_END-STACK_START+1)},
	}
}

func NewControllerWithBytes(bytes []byte) Controller {
	c := NewController()
	c.romImage.initWithBytes(bytes)
	return c
}

func (c *Controller) getRegister(addr uint16) (ByteRegister, bool) {
	switch addr {
	case 0xFF00:
		return &c.ControllerData, true
	case 0xFF0F:
		return &c.InterruptFlags, true
	case 0xFF40:
		return &c.LCDCFlags, true
	case 0xFF41:
		return &c.StatFlags, true
	case 0xFF42:
		return &c.SCY, true
	case 0xFF43:
		return &c.SCX, true
	case 0xFF44:
		return &c.LY, true
	case 0xFF45:
		return &c.LYC, true
	case 0xFF47:
		return &c.BGP, true
	case bootRomRegisterAddr:
		return &c.BootRomRegister, true
	case 0xFFFF:
		return &c.InterruptEnabled, true
	default:
		return nil, false
	}
}

func (c *Controller) LoadRomImage(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	romBytes := make([]byte, ROM_SIZE)
	_, err = io.ReadAtLeast(f, romBytes, ROM_SIZE)
	if err != nil {
		return err
	}

	c.romImage.initWithBytes(romBytes)

	return nil
}

func (c *Controller) ReadByte(addr uint16) byte {
	if c.isBootRoomAddr(addr) {
		return bootRom[addr]
	} else if c.isRomAddr(addr) {
		return c.romImage.ReadByte(addr)
	} else if c.isRamAddr(addr) {
		// working ram (ish)
		// todo: protect against access to forbidden areas?
		return c.ram.ReadByte(addr - ROM_SIZE)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		return reg.Read()
	} else if c.isStackAddr(addr) {
		return c.stack.ReadByte(addr - STACK_START)
	} else {
		return 0x00
	}
}

func (c *Controller) WriteByte(addr uint16, value byte) {
	if c.isBootRoomAddr(addr) {
		//panic("Ignoring request to write to boot rom")
	} else if c.isRomAddr(addr) {
		// todo: rom bank switching?
	} else if c.isRamAddr(addr) {
		// working ram
		// todo: protect against access to forbidden areas?
		c.ram.WriteByte(addr-ROM_SIZE, value)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		reg.Write(value)
	} else if c.isStackAddr(addr) {
		c.stack.WriteByte(addr-STACK_START, value)
	} else {
		panic("Unhandled memory location!")
	}
}

func (c *Controller) isStackAddr(addr uint16) bool {
	return addr >= STACK_START
}

func (c *Controller) isRamAddr(addr uint16) bool {
	return addr >= ROM_SIZE && addr < STACK_START
}

func (c *Controller) isRomAddr(addr uint16) bool {
	return addr < ROM_SIZE
}

func (c *Controller) isBootRoomAddr(addr uint16) bool {
	return addr < BOOT_ROM_SIZE && !c.BootRomRegister.isDisabled
}

func (c *Controller) ReadU16(addr uint16) uint16 {
	return (uint16(c.ReadByte(addr+1)) << 8) | uint16(c.ReadByte(addr))
}

func (c *Controller) WriteU16(addr, value uint16) {
	l := uint8(0x00FF & value)
	h := uint8((0xFF00 & value) >> 8)
	c.WriteByte(addr, l)
	c.WriteByte(addr+1, h)
}

func (c *Controller) ReadAll() []byte {
	result := make([]byte, 0xFFFF)
	for i := 0x0000; i < 0xFFFF; i += 1 {
		result[i] = c.ReadByte(uint16(i))
	}
	return result
}
