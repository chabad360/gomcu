package switches

type VPotMode byte
type VPotLED byte

const (
	VPot0 VPotLED = iota
	VPot1
	VPot2
	VPot3
	VPot4
	VPot5
	VPot6
	VPot7
	VPot8
	VPot9
	VPotA
	VPotB

	VPotMode0 VPotMode = 0x00
	VPotMode1 VPotMode = 0x10
	VPotMode2 VPotMode = 0x20
	VPotMode3 VPotMode = 0x30

	VPotDot VPotLED = 0x40
)
