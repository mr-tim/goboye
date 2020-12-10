package memory

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
	r.isDisabled = value&bootRomDisabledValue == bootRomDisabledValue
}
