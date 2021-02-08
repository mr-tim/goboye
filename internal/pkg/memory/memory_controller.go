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
	Divider          divRegister
	TimerCounter     simpleByteRegister
	TimerModulo      simpleByteRegister
	TimerController  timerController
	BootRomRegister  bootRomByteRegister
	LCDCFlags        register.LCDCFlags
	StatFlags        register.StatFlags
	SCY              simpleByteRegister
	SCX              simpleByteRegister
	LY               simpleByteRegister
	LYC              simpleByteRegister
	BGP              simpleByteRegister
	OBP0             simpleByteRegister
	OBP1             simpleByteRegister
	InterruptFlags   InterruptFlagsRegister
	InterruptEnabled InterruptEnabledRegister
	SerialOutput     string
	serialRequested  bool
	dmaStart         byte
}

func NewController() Controller {
	return Controller{
		romImage:       memoryMap{make([]byte, ROM_SIZE)},
		ram:            memoryMap{make([]byte, STACK_START-ROM_SIZE)},
		stack:          memoryMap{make([]byte, STACK_END-STACK_START+1)},
		ControllerData: NewControllerRegister(),
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
	case 0xFF04:
		return &c.Divider, true
	case 0xFF05:
		return &c.TimerCounter, true
	case 0xFF06:
		return &c.TimerModulo, true
	case 0xFF07:
		return &c.TimerController, true
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
	case 0xFF48:
		return &c.OBP0, true
	case 0xFF49:
		return &c.OBP1, true
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

func (c *Controller) ReadAddr(addr uint16) byte {
	if c.isBootRoomAddr(addr) {
		return bootRom[addr]
	} else if c.isRomAddr(addr) {
		return c.romImage.ReadAddr(addr)
	} else if c.isRamAddr(addr) {
		// working ram (ish)
		// todo: protect against access to forbidden areas?
		return c.ram.ReadAddr(addr - ROM_SIZE)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		return reg.Read()
	} else if addr == 0xFF46 {
		return c.dmaStart
	} else if c.isStackAddr(addr) {
		return c.stack.ReadAddr(addr - STACK_START)
	} else {
		return 0x00
	}
}

func (c *Controller) WriteAddr(addr uint16, value byte) {
	if c.isBootRoomAddr(addr) {
		//panic("Ignoring request to write to boot rom")
	} else if c.isRomAddr(addr) {
		// todo: rom bank switching?
	} else if c.isRamAddr(addr) {
		// working ram
		// todo: protect against access to forbidden areas?
		c.ram.WriteAddr(addr-ROM_SIZE, value)
	} else if reg, hasKey := c.getRegister(addr); hasKey {
		reg.Write(value)
	} else if addr == 0xFF01 {
		c.stack.WriteAddr(addr-STACK_START, value)
		if c.serialRequested {
			c.SerialOutput += string(value)
			c.WriteAddr(0xFF02, 0x00)
		}
	} else if addr == 0xFF02 {
		c.stack.WriteAddr(addr-STACK_START, value)
		c.serialRequested = value == 0x81
	} else if addr == 0xFF46 {
		// do a DMA transfer
		if value >= 0x80 && value < 0xE0 {
			c.dmaStart = value
			for i := 0; i < 0x0100; i += 1 {
				c.WriteAddr(0xFE00+uint16(i), c.ReadAddr(0x100*uint16(c.dmaStart)+uint16(i)))
			}
		}
	} else if c.isStackAddr(addr) {
		c.stack.WriteAddr(addr-STACK_START, value)
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

func (c *Controller) ReadAddrU16(addr uint16) uint16 {
	return (uint16(c.ReadAddr(addr+1)) << 8) | uint16(c.ReadAddr(addr))
}

func (c *Controller) WriteAddrU16(addr, value uint16) {
	l := uint8(0x00FF & value)
	h := uint8((0xFF00 & value) >> 8)
	c.WriteAddr(addr, l)
	c.WriteAddr(addr+1, h)
}

func (c *Controller) ReadAll() []byte {
	result := make([]byte, 0xFFFF)
	for i := 0x0000; i < 0xFFFF; i += 1 {
		result[i] = c.ReadAddr(uint16(i))
	}
	return result
}
