package gomcu

type State byte

const (
	StateOff      State = 0x00
	StateBlinking State = 0x01
	StateOn       State = 0x7F
)
