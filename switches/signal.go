package switches

type Signal byte

const (
	LessThan60 Signal = iota
	MoreThan60
	MoreThan50
	MoreThan40
	MoreThan30
	MoreThan20
	MoreThan14
	MoreThan10
	MoreThan8
	MoreThan6
	MoreThan4
	MoreThan2
	MoreThan0

	LenSignal = 12
	ClipOn    = 0x0E
	ClipOff   = 0x0F
)
