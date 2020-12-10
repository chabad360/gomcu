package switches

type Digit byte

const (
	Thousandth Digit = iota + 0x40
	Hundredth
	Tenth
	Second
	TenSecond
	Minute
	TenMinute
	Hour
	TenHour
	HundredHour
	AssignRight
	AssignLeft

	LenDigits = 12
	DigitDot  = 0x40
)
