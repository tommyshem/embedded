package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/pcd8544"
)

func main() {
	// configure led on the board
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// configure lcd screen
	dcPin := machine.GPIO5
	dcPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rstPin := machine.GPIO6
	rstPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	scePin := machine.GPIO7
	scePin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.SPI0.Configure(machine.SPIConfig{})
	lcd := pcd8544.New(machine.SPI0, dcPin, rstPin, scePin)
	lcd.Configure(pcd8544.Config{})
	// main loop
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)
		print("Blink Program is running.\n")
		led.High()
		time.Sleep(time.Millisecond * 100)
	}
}
