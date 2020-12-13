package gomcu

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/midimessage/sysex"
	"gitlab.com/gomidi/midi/writer"
	"strings"
	"time"
)

var (
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

// Reset resets and quickly triggers all the available features on the control surface.
// I recommend running this both to avoid some errors and as a sanity check to make sure that your entire control surface is working.
func Reset(wd *writer.Writer) {
	var m []midi.Message
	for i := 0; i < LenIDs; i++ {
		m = append(m, SetLED(Switch(i), StateOn))
	}
	for i := 0; i < LenChannels; i++ {
		m = append(m, SetFaderPos(Channel(i), 0x1FFF))
	}
	for i := 0; i < LenChannels-1; i++ {
		m = append(m, SetVPot(Channel(i), VPotMode3, VPot6+VPotDot))
	}
	for i := 0; i < LenChannels-1; i++ {
		m = append(m, SetMeter(Channel(i), Clipping))
	}
	for i := 0; i < LenDigits; i++ {
		m = append(m, SetDigit(Digit(uint8(i)+0x40), Char0+DigitDot))
	}
	for i := 0; i < LenLines; i++ {
		m = append(m, SetLCD(uint8(i), "A"))
	}

	writer.WriteMessages(wd, m)

	time.Sleep(time.Second)
	m = []midi.Message{}

	for i := 0; i < LenIDs; i++ {
		m = append(m, SetLED(Switch(i), StateOff))
	}
	for i := 0; i < LenChannels; i++ {
		m = append(m, SetFaderPos(Channel(i), -0x1FFF))
	}
	for i := 0; i < LenChannels-1; i++ {
		m = append(m, SetVPot(Channel(i), VPotMode0, VPot0))
	}
	for i := 0; i < LenChannels-1; i++ {
		m = append(m, SetMeter(Channel(i), ClipOff))
	}
	for i := 0; i < LenDigits; i++ {
		m = append(m, SetDigit(Digit(uint8(i)+0x40), SymbolSpace))
	}
	for i := 0; i < LenLines; i++ {
		m = append(m, SetLCD(uint8(i), " "))
	}
	writer.WriteMessages(wd, m)
}

// SetLED sets the chosen button's LED to the chosen State.
func SetLED(led Switch, state State) midi.Message {
	return channel.Channel(0).NoteOn(uint8(led), byte(state))
}

// SetFaderPos sets the position of the chosen fader to a number between 0 (bottom) and 16382 (top).
func SetFaderPos(fader Channel, pos uint16) midi.Message {
	p := int16(pos) - 8191
	return channel.Channel(fader).Pitchbend(p)
}

// SetTimeDisplay sets multiple characters on the timecode display.
// Note: letters is limited to ten characters and is right aligned.
// Refer to timecode.Digit for valid characters.
func SetTimeDisplay(letters string) (m []midi.Message) {
	bytes := []byte(strings.ToUpper(letters))
	if len(bytes) > 10 {
		bytes = bytes[:10]
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
		m = append(m, channel.Channel(15).ControlChange(i+0x40, bytes[i]))
	}
	return
}

// SetDigit sets an individual digit on the timecode or Assignment section.
// Refer to Char for more information on valid characters.
func SetDigit(digit Digit, char Char) midi.Message {
	if (char >= 0x40 && char <= 0x60) || (char >= 0x80 && char <= 0xA0) {
		char = char - 0x40
	}
	return channel.Channel(15).ControlChange(byte(digit), byte(char))
}

// SetLCD sets the text (an ASCII string) found on the LCD starting from the specified offset.
func SetLCD(offset uint8, text string) midi.Message {
	return sysex.SysEx(append(append(header, 0x12, offset), []byte(text)...))
}

// SetVPot sets the LEDs around the knobs (VPots).
// Refer to VPotMode for an explanation of the various VPot modes.
func SetVPot(ch Channel, mode VPotMode, led VPotLED) midi.Message {
	return channel.Channel(0).ControlChange(byte(ch+0x30), byte(mode)+byte(led))
}

// SetMeter sets the level meter for the selected Channel to the desired value.
func SetMeter(ch Channel, value MeterLevel) midi.Message {
	return channel.Channel(0).Aftertouch(byte(ch<<4) + byte(value))
}
