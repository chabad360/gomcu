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

var (
	Channels = map[int]Channel{
		1: 0x10,
		2: 0x11,
		3: 0x12,
		4: 0x13,
		5: 0x14,
		6: 0x15,
		7: 0x16,
		8: 0x17,
		9: 0x18,
	}
)
