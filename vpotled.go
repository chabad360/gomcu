package gomcu

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
	VPot10
	VPot11

	LenVPots = 12

	// VPotDot is the bottom LED separate from the other 11.
	VPotDot VPotLED = 0x40
)
