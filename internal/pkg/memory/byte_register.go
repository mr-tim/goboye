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
