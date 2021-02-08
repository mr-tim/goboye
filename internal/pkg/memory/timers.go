package memory

import "github.com/mr-tim/goboye/internal/pkg/utils"

type divRegister struct {
	value      byte
	cycleCount uint16
}

func (r *divRegister) Read() byte {
	return r.value
}

func (r *divRegister) Write(_value byte) {
	// any write resets the div register to 0
	r.value = 0x00
}

func (r *divRegister) Update(cycles uint8) {
	r.cycleCount += uint16(cycles)
	if r.cycleCount > 255 {
		r.cycleCount = 0
		r.value += 1
	}
}

type timerController struct {
	value           byte
	timerCycleCount int
}

func (c *timerController) Read() byte {
	return c.value
}

func (c *timerController) Write(value byte) {
	previousClockSelect := c.GetInputClockSelect()
	c.value = value & 0x07
	currentClockSelect := c.GetInputClockSelect()

	if previousClockSelect != currentClockSelect {
		c.timerCycleCount = currentClockSelect.Countdown()
	}
}

func (c *timerController) IsStarted() bool {
	return utils.IsBitSet(c.value, 2)
}

func (c *timerController) GetInputClockSelect() ClockSelect {
	return ClockSelect(c.value & 0x03)
}

func (c *timerController) UpdateCountdown(cycles uint8) bool {
	c.timerCycleCount -= int(cycles)
	trigger := c.timerCycleCount <= 0
	if trigger {
		c.timerCycleCount = c.GetInputClockSelect().Countdown()
	}
	return trigger
}

type ClockSelect byte

var (
	Clock4k   ClockSelect = 0
	Clock262k ClockSelect = 1
	Clock65k  ClockSelect = 2
	Clock16k  ClockSelect = 3
)

func (s ClockSelect) Countdown() int {
	switch s {
	case Clock4k:
		return 1024
	case Clock262k:
		return 16
	case Clock65k:
		return 64
	case Clock16k:
		return 256
	}
	panic("Unsupported clock select!")
}
