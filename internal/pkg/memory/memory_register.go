package memory

import "github.com/mr-tim/goboye/internal/pkg/utils"

type RoRegister interface {
	IsBitSet(index byte) bool
	GetByte() byte
}

type RwRegister interface {
	RoRegister
	SetValue(value byte)
}

type RoRegisterAtAddr struct {
	memoryMap *memoryMap
	addr      uint16
}

type RwRegisterAtAddr struct {
	roRegister RoRegisterAtAddr
}

func (ro *RoRegisterAtAddr) IsBitSet(index byte) bool {
	return utils.IsBitSet(ro.GetByte(), index)
}

func (ro *RoRegisterAtAddr) GetByte() byte {
	return ro.memoryMap.ReadByte(ro.addr)
}

func (rw *RwRegisterAtAddr) IsBitSet(index byte) bool {
	return rw.roRegister.IsBitSet(index)
}

func (rw *RwRegisterAtAddr) GetByte() byte {
	return rw.roRegister.GetByte()
}

func (rw *RwRegisterAtAddr) SetValue(value byte) {
	rw.roRegister.memoryMap.WriteByte(rw.roRegister.addr, value)
}
