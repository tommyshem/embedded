package main

import (
	"machine"
	"time"

	"image/color"

	"tinygo.org/x/drivers/st7789"
)

func main() {

	println("Graphic Lcd Program V1.0")
	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
		SCK:       machine.GP2,
		SDO:       machine.GP4,
	})
	if err != nil {
		println("error setting up")
	} else {
		println("Setup ok")
	}
	display := st7789.New(machine.SPI0,
		machine.GP6, // TFT_RESET
		machine.GP7, // TFT_DC
		machine.GP5, // TFT_CS
		machine.GP8) // TFT_BL

	display.Configure(st7789.Config{
		Rotation:   st7789.NO_ROTATION,
		RowOffset:  80,
		FrameRate:  st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
		Width:      240,
		Height:     135,
	})

	width, height := display.Size()

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	display.FillScreen(black)

	display.FillRectangle(0, 0, width/2, height/2, white)
	display.FillRectangle(width/2, 0, width/2, height/2, red)
	display.FillRectangle(0, height/2, width/2, height/2, green)
	display.FillRectangle(width/2, height/2, width/2, height/2, blue)
	display.FillRectangle(width/4, height/4, width/2, height/2, black)

	for {
		machine.LED.Low()
		time.Sleep(time.Millisecond * 500)
		print("Graphic Lcd - Program is running.\n")
		machine.LED.High()
		time.Sleep(time.Millisecond * 100)
	}
}
