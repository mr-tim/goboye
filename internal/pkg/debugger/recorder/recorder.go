package recorder

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/memory"
	"github.com/rs/zerolog"
	"os"
)

type Snapshot struct {
	registers cpu.Registers
	address   uint16
	op        cpu.OpcodeAndPayload
	payload   string
	cycles    uint
}

func (s Snapshot) String() string {
	return fmt.Sprintf("%d %s 0x%04x: %s", s.cycles, s.registers.String(), s.address, s.op.Disassembly())
}

type Recorder struct {
	snapshots    []Snapshot
	maxSnapshots int
	currentIndex int
	logger       zerolog.Logger
}

func (r *Recorder) TakeSnapshot(processor cpu.Processor, memory *memory.Controller) {
	if r.maxSnapshots > 0 {
		registers := processor.DebugRegisters()
		addr, op, _ := cpu.Disassemble(memory, processor.GetRegisterPair(cpu.RegisterPairPC))
		s := Snapshot{
			cycles:    processor.Cycles(),
			registers: registers,
			address:   addr,
			op:        op,
		}
		r.snapshots[r.currentIndex] = s
		r.currentIndex = (r.currentIndex + 1) % r.maxSnapshots
		r.logger.Info().Msgf("%d %s 0x%04x: %s", s.cycles, s.registers.String(), s.address, s.op.Disassembly())
	}
}

func (r *Recorder) GetSnapshots() []*Snapshot {
	result := make([]*Snapshot, r.maxSnapshots)
	idx := 0
	for i := r.currentIndex; i < r.maxSnapshots; i += 1 {
		result[idx] = &r.snapshots[i]
		idx += 1
	}
	for i := 0; i < r.currentIndex; i += 1 {
		result[idx] = &r.snapshots[i]
		idx += 1
	}
	return result
}

func NewRecorder(maxSnapshots int) *Recorder {
	logger := createLogger(true)
	return &(Recorder{
		maxSnapshots: maxSnapshots,
		snapshots:    make([]Snapshot, maxSnapshots),
		currentIndex: 0,
		logger:       logger,
	})
}

func createLogger(enabled bool) zerolog.Logger {
	if enabled {
		f, err := os.Create("flight-recorder.log")
		if err != nil {
			panic(err)
		}
		return zerolog.New(f).With().Timestamp().Logger()
	} else {
		return zerolog.Nop()
	}
}
