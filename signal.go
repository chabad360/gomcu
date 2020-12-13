package gomcu

// MeterLevel is the height (in decibels) that the meter should be set to.
type MeterLevel byte

const (
	LessThan60 MeterLevel = iota
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
	_
	// Clipping enables the clipping warning LED on the meter.
	Clipping
	// ClipOff disables the clipping warning LED on the meter.
	ClipOff

	LenMeterLevel = 15
)
