package recorder

import (
	"fmt"
	"github.com/mr-tim/goboye/internal/pkg/cpu"
	"github.com/mr-tim/goboye/internal/pkg/memory"
)

type Snapshot struct {
	registers cpu.Registers
	address   uint16
	op        cpu.Opcode
	payload   string
}

func (s Snapshot) String() string {
	return fmt.Sprintf("%s 0x%04x: %s %s", s.registers.String(), s.address, s.op.Disassembly(), s.payload)
}

type Recorder struct {
	snapshots    []Snapshot
	maxSnapshots int
	currentIndex int
}

func (r *Recorder) TakeSnapshot(processor cpu.Processor, memoryMap memory.MemoryMap) {
	if r.maxSnapshots > 0 {
		registers := processor.DebugRegisters()
		d := cpu.NewDisassembler(memoryMap)
		d.SetPos(processor.GetRegisterPair(cpu.RegisterPairPC))
		// TODO: can probably make this more efficient by just storing PC + any operands?
		addr, op, payload := d.GetNextInstruction()
		r.snapshots[r.currentIndex] = Snapshot{
			registers: registers,
			address:   addr,
			op:        op,
			payload:   payload,
		}
		r.currentIndex = (r.currentIndex + 1) % r.maxSnapshots
	}
}

func (r *Recorder) GetSnapshots() []*Snapshot {
	result := make([]*Snapshot, r.maxSnapshots)
	idx := 0
	for i := r.currentIndex; i<r.maxSnapshots; i+=1 {
		result[idx] = &r.snapshots[i]
		idx += 1
	}
	for i := 0; i<r.currentIndex; i+=1 {
		result[idx] = &r.snapshots[i]
		idx += 1
	}
	return result
}

func NewRecorder(maxSnapshots int) *Recorder {
	return &(Recorder{
		maxSnapshots: maxSnapshots,
		snapshots: make([]Snapshot, maxSnapshots),
		currentIndex: 0,
	})
}
