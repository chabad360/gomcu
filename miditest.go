package main

import (
	"fmt"
	driver "gitlab.com/gomidi/rtmididrv"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// This example expects the first input and output port to be connected
// somehow (are either virtual MIDI through ports or physically connected).
// We write to the out port and listen to the in port.
func main() {
	drv, err := driver.New()
	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	in, out := ins[1], outs[1]

	must(in.Open())
	must(out.Open())

	control, _ := New(in, out)
	//control.Reset()

	//control.SetFaderPos(switches.Channel1, 0x1FFF)
	fmt.Println(control.SetTimeDisplay("helloworld"))

	//control.writer.SetChannel(15)
	//fmt.Println(control.SetDigit(switches.AssignRight, '_' + switches.DigitDot))
}
