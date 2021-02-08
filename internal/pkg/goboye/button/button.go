package button

type Button uint8

func (b Button) IsJoypad() bool {
	switch b {
	case Left, Right, Up, Down, A, B, Start, Select:
		return true
	default:
		return false
	}
}

const (
	Unbound Button = iota
	Quit
	Frames

	Left
	Right
	Up
	Down
	A
	B
	Start
	Select
)
