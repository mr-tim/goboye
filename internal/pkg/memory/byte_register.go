package memory

type ByteRegister interface {
	Read() byte
	Write(value byte)
}

type simpleByteRegister struct {
	value byte
}

func (r *simpleByteRegister) Read() byte {
	return r.value
}

func (r *simpleByteRegister) Write(value byte) {
	r.value = value
}

type bootRomByteRegister struct {
	isDisabled bool
}

func (r *bootRomByteRegister) Read() byte {
	if r.isDisabled {
		return bootRomDisabledValue
	} else {
		return bootRomEnabledValue
	}
}

func (r *bootRomByteRegister) Write(value byte) {
	r.isDisabled = value & bootRomDisabledValue == bootRomDisabledValue
}
