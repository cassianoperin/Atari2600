package Graphics

import (
	"fmt"
	"time"
	"github.com/faiface/pixel/pixelgl"
	"Atari2600/CPU"
)

func keyboard() {

	// Enable Disable Debug
	if win.Pressed(pixelgl.Key9) {
		if CPU.Debug {
			CPU.Debug = false
			fmt.Printf("\t\tDEBUG mode Disabled\n")
			time.Sleep(500 * time.Millisecond)
		} else {
			CPU.Debug = true
			fmt.Printf("\t\tDEBUG mode Enabled\n")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Reset
	if win.Pressed(pixelgl.Key0) {
		// F000 - FFFF
		var ROM_dump = [4096]byte{}

		// Dump the rom from memory prior to clear all information
		for i := 0 ; i < 4096 ; i ++{
			ROM_dump[i] = CPU.Memory[0xF000+i]
		}

		// Workaround for WSYNC before VSYNC
		VSYNC_passed = false

		CPU.Initialize()

		// Restore ROM to memory
		for i := 0 ; i < 4096 ; i ++{
			CPU.Memory[0xF000+i] = ROM_dump[i]
		}

		// Reset graphics
		renderGraphics()
		// Restart Draw from the beginning
		line = 1

		// Players Vertical Positioning
		XPositionP0		= 0
		XFinePositionP0	= 0
		XPositionP1		= 0
		XFinePositionP1	= 0

		// ------------------ Personal Control Flags ------------------ //
		CPU.Beam_index	= 0		// Beam index to control where to draw objects using cpu cycles

		CPU.Reset()
	}

	// CPU.Pause Key
	if win.Pressed(pixelgl.KeyP) {
		if CPU.Pause {
			CPU.Pause = false
			fmt.Printf("\t\tPAUSE mode Disabled\n")
			time.Sleep(500 * time.Millisecond)
		} else {
			CPU.Pause = true
			fmt.Printf("\t\tPAUSE mode Enabled\n")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Step Forward
	if win.Pressed(pixelgl.KeyI) {
		if CPU.Pause {
			fmt.Printf("\t\tStep Forward\n")


			// Runs the interpreter
			CPU.Interpreter()

			// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
			CRT( CPU.TIA_Update )

			// Reset to default value
			CPU.TIA_Update = -1


			win.Update()

			time.Sleep(50 * time.Millisecond)
		}
	}

	// -------------- PLAYER 0 -------------- //
	// P0 Right
	if win.Pressed(pixelgl.KeyRight) {
		CPU.Memory[CPU.SWCHA] = 0x7F // 0111 1111
	}
	// P0 Left
	if win.Pressed(pixelgl.KeyLeft) {
		CPU.Memory[CPU.SWCHA] = 0xBF // 1011 1111
	}
	// P0 Down
	if win.Pressed(pixelgl.KeyDown) {
		CPU.Memory[CPU.SWCHA] = 0xDF // 1101 1111
	}
	// P0 Up
	if win.Pressed(pixelgl.KeyUp) {
		CPU.Memory[CPU.SWCHA] = 0xEF // 1110 1111
	}

	// -------------- PLAYER 1 -------------- //
	// P1 Right
	if win.Pressed(pixelgl.KeyD) {
		CPU.Memory[CPU.SWCHA] = 0xF7 // 1111 0111
	}
	// P1 Left
	if win.Pressed(pixelgl.KeyA) {
		CPU.Memory[CPU.SWCHA] = 0xFB // 1111 1011
	}
	// P1 Down
	if win.Pressed(pixelgl.KeyS) {
		CPU.Memory[CPU.SWCHA] = 0xFD // 1111 1101
	}
	// P1 Up
	if win.Pressed(pixelgl.KeyW) {
		CPU.Memory[CPU.SWCHA] = 0xFE // 1111 1110
	}
}
