package switches

type Channel byte

const (
	Channel1 Channel = iota
	Channel2
	Channel3
	Channel4
	Channel5
	Channel6
	Channel7
	Channel8
	Master
	LenChannels = 9
	FaderMax    = 0x1FFF
	FaderMin    = -0x1FFF
)
