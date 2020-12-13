package gomcu

// VPot Mode is the display mode that should be used when setting the amount of VPot LEDs to display.
type VPotMode byte

const (
	VPotMode0 VPotMode = 0x00
	VPotMode1 VPotMode = 0x10
	VPotMode2 VPotMode = 0x20
	VPotMode3 VPotMode = 0x30
)
