package memory

import (
	"github.com/mr-tim/goboye/internal/pkg/display/register"
	"io"
	"log"
	"os"
)

type Controller struct {
	romImage        memoryMap
	ram             memoryMap
	stack           memoryMap
	BootRomRegister bootRomByteRegister
	LCDCFlags       register.LCDCFlags
	StatFlags       register.StatFlags
	SCY             simpleByteRegister
	SCX             simpleByteRegister
	LY              simpleByteRegister
	LYC             simpleByteRegister
	BGP             simpleByteRegister
}

func NewController() Controller {
	return Controller{
		romImage: memoryMap{make([]byte, ROM_SIZE)},
		ram:      memoryMap{make([]byte, 0xFF00-0x8000)},
		stack:    memoryMap{make([]byte, 0xFFFE-0xFF00)},
	}
}

func NewControllerWithBytes(bytes []byte) Controller {
	c := NewController()
	c.romImage.initWithBytes(bytes)
	return c
}

func (c *Controller) getRegister(addr uint16) (ByteRegister, bool) {
	switch addr {
	case bootRomRegisterAddr:
		return &c.BootRomRegister, true
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
	if addr < 0x0100 && !c.BootRomRegister.isDisabled {
		return bootRom[addr]
	} else if addr < ROM_SIZE {
		return c.romImage.ReadByte(addr)
	} else if addr >= 0x8000 && addr < 0xFF00 {
		// working ram (ish)
		// todo: protect against access to forbidden areas?
		return c.ram.ReadByte(addr - 0x8000)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		return reg.Read()
	} else if addr >= 0xFF00 && addr < 0xFFFE {
		return c.stack.ReadByte(addr - 0xFF00)
	} else {
		return 0x00
	}
}

func (c *Controller) WriteByte(addr uint16, value byte) {
	if addr < 0x0100 && !c.BootRomRegister.isDisabled {
		log.Printf("Ignoring request to write to boot rom")
	} else if addr < ROM_SIZE {
		log.Printf("Ignoring request to write to rom")
	} else if addr >= 0x8000 && addr < 0xFF00 {
		// working ram
		// todo: protect against access to forbidden areas?
		c.ram.WriteByte(addr-0x8000, value)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		reg.Write(value)
	} else if addr >= 0xFF00 && addr < 0xFFFE {
		c.stack.WriteByte(addr-0xFF00, value)
	} else {
		panic("Unhandled memory location!")
	}
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
