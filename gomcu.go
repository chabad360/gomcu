package main

import (
	"github.com/chabad360/gomcu/switches"
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

func Reset(wd *writer.Writer) {
	m := []midi.Message{}
	for i := 0; i <= switches.LenIDs-1; i++ {
		m = append(m, LEDOn(switches.Switch(i)))
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		m = append(m, SetFaderPos(switches.Channel(i), 0x1FFF))
	}
	for i := 0; i <= switches.LenChannels-2; i++ {
		m = append(m, SetVPot(switches.Channel(i), switches.VPotMode3, switches.VPot6+switches.VPotDot))
	}
	for i := 0; i <= switches.LenChannels-2; i++ {
		m = append(m, SetMeter(switches.Channel(i), switches.ClipOn))
	}
	for i := 0; i <= switches.LenDigits-1; i++ {
		m = append(m, SetDigit(switches.Digit(uint8(i)+0x40), switches.Char0+switches.DigitDot))
	}

	writer.WriteMessages(wd, m)

	time.Sleep(time.Second)
	m = []midi.Message{}

	for i := 0; i <= switches.LenIDs-1; i++ {
		m = append(m, LEDOff(switches.Switch(i)))
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		m = append(m, SetFaderPos(switches.Channel(i), -0x1FFF))
	}
	for i := 0; i <= switches.LenChannels-1; i++ {
		m = append(m, SetVPot(switches.Channel(i), switches.VPotMode0, switches.VPot0))
	}
	for i := 0; i <= switches.LenChannels-2; i++ {
		m = append(m, SetMeter(switches.Channel(i), switches.ClipOff))
	}
	for i := 0; i <= switches.LenDigits-1; i++ {
		m = append(m, SetDigit(switches.Digit(uint8(i)+0x40), switches.SymbolSpace))
	}
	writer.WriteMessages(wd, m)
}

func LEDOn(led switches.Switch) midi.Message {
	return channel.Channel(0).NoteOn(uint8(led), 0xF7)
}

func LEDOff(led switches.Switch) midi.Message {
	return channel.Channel(0).NoteOff(uint8(led))
}

func LEDBlink(led switches.Switch) midi.Message {
	return channel.Channel(0).NoteOn(uint8(led), 0x01)
}

func SetFaderPos(fader switches.Channel, pos int16) midi.Message {
	return channel.Channel(fader).Pitchbend(pos)
}

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

func SetDigit(digit switches.Digit, char switches.Char) midi.Message {
	if (char >= 0x40 && char <= 0x60) || (char >= 0x80 && char <= 0xA0) {
		char = char - 0x40
	}
	return channel.Channel(15).ControlChange(byte(digit), byte(char))
}

func SetLCD(offset uint8, text string) midi.Message {
	return sysex.SysEx(append(append(header, 0x12, offset), []byte(text)...))
}

func SetVPot(ch switches.Channel, mode switches.VPotMode, led switches.VPotLED) midi.Message {
	return channel.Channel(0).ControlChange(byte(ch+0x30), byte(mode)+byte(led))
}

func SetMeter(ch switches.Channel, value switches.Signal) midi.Message {
	return channel.Channel(0).Aftertouch(byte(ch<<4) + byte(value))
}
