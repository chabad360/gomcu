package main

import (
	"fmt"
	"github.com/chabad360/gomcu/switches"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
	"strings"
	"time"
)

var (
	Letters = map[rune]byte{
		'@':  0x00,
		'A':  0x01,
		'B':  0x02,
		'C':  0x03,
		'D':  0x04,
		'E':  0x05,
		'F':  0x06,
		'G':  0x07,
		'H':  0x08,
		'I':  0x09,
		'J':  0x0A,
		'K':  0x0B,
		'L':  0x0C,
		'M':  0x0D,
		'N':  0x0E,
		'O':  0x0F,
		'P':  0x10,
		'Q':  0x11,
		'R':  0x12,
		'S':  0x13,
		'T':  0x14,
		'U':  0x15,
		'V':  0x16,
		'W':  0x17,
		'X':  0x18,
		'Y':  0x19,
		'Z':  0x1A,
		'[':  0x1B,
		'\\': 0x1C,
		']':  0x1D,
		'^':  0x1E,
		'_':  0x1F,
		' ':  0x20,
		'!':  0x21,
		'"':  0x22,
		'#':  0x23,
		'$':  0x24,
		'%':  0x25,
		'&':  0x26,
		'\'': 0x27,
		'(':  0x28,
		')':  0x29,
		'*':  0x2A,
		'+':  0x2B,
		',':  0x2C,
		'-':  0x2D,
		'.':  0x2E,
		'/':  0x2F,
		'0':  0x30,
		'1':  0x31,
		'2':  0x32,
		'3':  0x33,
		'4':  0x34,
		'5':  0x35,
		'6':  0x36,
		'7':  0x37,
		'8':  0x38,
		'9':  0x39,
		':':  0x3A,
		';':  0x3B,
		'<':  0x3C,
		'=':  0x3D,
		'>':  0x3E,
		'?':  0x3F,
	}

	Channels = map[int]byte{
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

	IDs = map[string]switches.Switch{
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

	header = []byte{0x00, 0x00, 0x66, 0x14}

	SysExMessages = map[string][]byte{
		"Query":       []byte{0x00},
		"GoOffline":   []byte{0x0F, 0x7F},
		"Version":     []byte{0x13, 0x00},
		"ResetFaders": []byte{0x61},
		"ResetLEDs":   []byte{0x62},
		"Reset":       []byte{0x63},
	}
)

type ControlSurface struct {
	writer *writer.Writer
	reader *reader.Reader
	buffer chan midi.Message
	Serial [5]byte
}

func New(_ midi.In, out midi.Out) (*ControlSurface, error) {
	s := &ControlSurface{}
	s.writer = writer.New(out)
	s.reader = reader.New(s.callback)

	s.buffer = make(chan midi.Message, 1024)

	return s, nil
}

func (s *ControlSurface) callback(*reader.Reader) {}

func (s *ControlSurface) Reset() {
	for i := 0; i <= switches.LenIDs-1; i++ {
		s.LEDOn(switches.Switch(i))
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		s.SetFaderPos(switches.Channel(i), 0x1FFF)
	}
	for i := 0; i <= switches.LenChannels-2; i++ {
		s.SetVPot(switches.Channel(i), switches.VPotMode3, switches.VPot6+switches.VPotDot)
	}
	for i := 0; i <= switches.LenChannels-2; i++ {
		s.SetMeter(switches.Channel(i), switches.MoreThan0)
	}

	time.Sleep(time.Second)

	for i := 0; i <= switches.LenIDs-1; i++ {
		s.LEDOff(switches.Switch(i))
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		s.SetFaderPos(switches.Channel(i), -0x1FFF)
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		s.SetVPot(switches.Channel(i), switches.VPotMode0, switches.VPot0)
	}
}

func (s *ControlSurface) LEDOn(led switches.Switch) error {
	return writer.NoteOn(s.writer, uint8(led), 0xF7)
}

func (s *ControlSurface) LEDOff(led switches.Switch) error {
	return writer.NoteOn(s.writer, uint8(led), 0x00)
}

func (s *ControlSurface) LEDBlink(led switches.Switch) error {
	return writer.NoteOn(s.writer, uint8(led), 0x01)
}

func (s *ControlSurface) SetFaderPos(fader switches.Channel, pos int16) error {
	s.writer.SetChannel(uint8(fader))
	defer s.writer.SetChannel(0)

	return writer.Pitchbend(s.writer, pos)
}

func (s *ControlSurface) SetTimeDisplay(letters string) error {
	bytes := []byte(strings.ToUpper(letters))
	if len(bytes) > 10 {
		return fmt.Errorf("SetTimeDisplay: more than ten characters sent")
	}

	for i, char := range bytes {
		if char >= 0x40 && char <= 0x60 {
			bytes[i] = char - 0x40
		}
	}

	for i := len(bytes)/2 - 1; i >= 0; i-- {
		opp := len(bytes) - 1 - i
		bytes[i], bytes[opp] = bytes[opp], bytes[i]
	}

	for i := uint8(0); int(i) < len(bytes); i++ {
		writer.ControlChange(s.writer, i+0x40, bytes[i])
	}
	return nil
}

func (s *ControlSurface) SetDigit(digit switches.Digit, char switches.Char) error {
	if (char >= 0x40 && char <= 0x60) || (char >= 0x80 && char <= 0xA0) {
		char = char - 0x40
	}
	return writer.ControlChange(s.writer, byte(digit), byte(char))
}

func (s *ControlSurface) SetLCD(offset uint8, text string) error {
	if offset > 112 || (offset+uint8(len(text)) > 112) {
		return fmt.Errorf("SetLCD: text and/or offset to larger than available lcd space")
	}

	return writer.SysEx(s.writer, append(append(header, 0x12, offset), []byte(text)...))
}

func (s *ControlSurface) SetVPot(channel switches.Channel, mode switches.VPotMode, led switches.VPotLED) error {
	return writer.ControlChange(s.writer, byte(channel+0x30), byte(mode)+byte(led))
}

func (s *ControlSurface) SetMeter(channel switches.Channel, value switches.Signal) error {
	writer.SysEx(s.writer, append(header, 0x21, 0x00))
	return writer.Aftertouch(s.writer, byte(channel<<4)+byte(value))
}
