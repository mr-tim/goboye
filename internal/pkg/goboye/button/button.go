package button

type Button uint8

const (
	Unbound Button = iota
	Quit
	Left
	Right
	Up
	Down
	A
	B
	Start
	Select
)

