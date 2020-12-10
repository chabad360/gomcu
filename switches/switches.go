package switches

type Switch byte

const (
	Rec1 Switch = iota
	Rec2
	Rec3
	Rec4
	Rec5
	Rec6
	Rec7
	Rec8
	Solo1
	Solo2
	Solo3
	Solo4
	Solo5
	Solo6
	Solo7
	Solo8
	Mute1
	Mute2
	Mute3
	Mute4
	Mute5
	Mute6
	Mute7
	Mute8
	Select1
	Select2
	Select3
	Select4
	Select5
	Select6
	Select7
	Select8
	V1
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	AssignTrack
	AssignSend
	AssignPan
	AssignPlugin
	AssignEQ
	AssignInstrument
	BankL
	BankR
	ChannelL
	ChannelR
	Flip
	GlobalView
	NameValue
	SMPTEBeats
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	MIDITracks
	Inputs
	AudioTracks
	AudioInstrument
	Aux
	Busses
	Outputs
	User
	Shift
	Option
	Control
	CMDAlt
	Read
	Write
	Trim
	Touch
	Latch
	Group
	Save
	Undo
	Cancel
	Enter
	Marker
	Nudge
	Cycle
	Drop
	Replace
	Click
	Solo
	Rewind
	FastFwd
	Stop
	Play
	Record
	Up
	Down
	Left
	Right
	Zoom
	Scrub
	UserA
	UserB
	Fader1
	Fader2
	Fader3
	Fader4
	Fader5
	Fader6
	Fader7
	Fader8
	FaderMaster
	STMPELED
	BeatsLED
	RudeSoloLED
	RelayClickLED
	LenIDs = 120
)

var (
	IDs = map[string]Switch{
		"Rec1":             0x00,
		"Rec2":             0x01,
		"Rec3":             0x02,
		"Rec4":             0x03,
		"Rec5":             0x04,
		"Rec6":             0x05,
		"Rec7":             0x06,
		"Rec8":             0x07,
		"Solo1":            0x08,
		"Solo2":            0x09,
		"Solo3":            0x0A,
		"Solo4":            0x0B,
		"Solo5":            0x0C,
		"Solo6":            0x0D,
		"Solo7":            0x0E,
		"Solo8":            0x0F,
		"Mute1":            0x10,
		"Mute2":            0x11,
		"Mute3":            0x12,
		"Mute4":            0x13,
		"Mute5":            0x14,
		"Mute6":            0x15,
		"Mute7":            0x16,
		"Mute8":            0x17,
		"Select1":          0x18,
		"Select2":          0x19,
		"Select3":          0x1A,
		"Select4":          0x1B,
		"Select5":          0x1C,
		"Select6":          0x1D,
		"Select7":          0x1E,
		"Select8":          0x1F,
		"V1":               0x20,
		"V2":               0x21,
		"V3":               0x22,
		"V4":               0x23,
		"V5":               0x24,
		"V6":               0x25,
		"V7":               0x26,
		"V8":               0x27,
		"AssignTrack":      0x28,
		"AssignSend":       0x29,
		"AssignPan":        0x2A,
		"AssignPlugin":     0x2B,
		"AssignEQ":         0x2C,
		"AssignInstrument": 0x2D,
		"BankL":            0x2E,
		"BankR":            0x2F,
		"ChannelL":         0x30,
		"ChannelR":         0x31,
		"Flip":             0x32,
		"GlobalView":       0x33,
		"Name/Value":       0x34,
		"SMPTE/Beats":      0x35,
		"F1":               0x36,
		"F2":               0x37,
		"F3":               0x38,
		"F4":               0x39,
		"F5":               0x3A,
		"F6":               0x3B,
		"F7":               0x3C,
		"F8":               0x3D,
		"MIDITracks":       0x3E,
		"Inputs":           0x3F,
		"AudioTracks":      0x40,
		"AudioInstrument":  0x41,
		"Aux":              0x42,
		"Busses":           0x43,
		"Outputs":          0x44,
		"User":             0x45,
		"Shift":            0x46,
		"Option":           0x47,
		"Control":          0x48,
		"CMD":              0x49,
		"Read":             0x4A,
		"Write":            0x4B,
		"Trim":             0x4C,
		"Touch":            0x4D,
		"Latch":            0x4E,
		"Group":            0x4F,
		"Save":             0x50,
		"Undo":             0x51,
		"Cancel":           0x52,
		"Enter":            0x53,
		"Marker":           0x54,
		"Nudge":            0x55,
		"Cycle":            0x56,
		"Drop":             0x57,
		"Replace":          0x58,
		"Click":            0x59,
		"Solo":             0x5A,
		"Rewind":           0x5B,
		"FastFwd":          0x5C,
		"Stop":             0x5D,
		"Play":             0x5E,
		"Record":           0x5F,
		"Up":               0x60,
		"Down":             0x61,
		"Left":             0x62,
		"Right":            0x63,
		"Zoom":             0x64,
		"Scrub":            0x65,
		"UserA":            0x66,
		"UserB":            0x67,
		"Fader1":           0x68,
		"Fader2":           0x69,
		"Fader3":           0x6A,
		"Fader4":           0x6B,
		"Fader5":           0x6C,
		"Fader6":           0x6D,
		"Fader7":           0x6E,
		"Fader8":           0x6F,
		"FaderMaster":      0x70,
		"STMPELED":         0x71,
		"BeatsLED":         0x72,
		"RudeSoloLED":      0x73,
		"RelayClickLED":    0x74,
	}
)
