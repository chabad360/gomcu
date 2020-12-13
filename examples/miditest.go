package main

import (
	"github.com/chabad360/gomcu"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
	"time"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// This example demonstrates using github.com/chabad360/gomcu to send and receive signals using the Mackie Control Protocol.
func main() {
	drv, err := driver.New()
	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	wr := writer.New(out)
	rd := reader.New(
		reader.NoteOn(noteon(wr)))
	gomcu.Reset(wr)

	go rd.ListenTo(in)

	m := []midi.Message{gomcu.SetDigit(gomcu.AssignLeft, 'H'), gomcu.SetDigit(gomcu.AssignRight, 'W'), gomcu.SetLCD(0, "Hello,"), gomcu.SetLCD(56, "World")}
	m = append(m, gomcu.SetTimeDisplay("helloWorld")...)
	writer.WriteMessages(wr, m)

	time.Sleep(time.Second * 20)
}

func noteon(wr *writer.Writer) func(p *reader.Position, c, k, v uint8) {
	return func(p *reader.Position, c, k, v uint8) {
		switch gomcu.Switch(k) {
		case gomcu.Play:
			wr.Write(gomcu.SetLED(gomcu.Play, gomcu.StateOn))
		}
	}
}
