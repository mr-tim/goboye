package memory

type controllerRegister struct {

}

func (r *controllerRegister) Read() byte {
	// output is pinned high by default
	return 0x0F
}

func (r *controllerRegister) Write(value byte) {

}
