package gomcu

// A Digit is a single 7-segment display on the Control Surface
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
)
