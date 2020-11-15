package VGS

import (
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Keyboard() {

	// Debug
	if win.JustPressed(pixelgl.Key9) {
		if Debug {
			stopDebug()
		} else {
			startDebug()
		}
	}

	// Reset
	if win.Pressed(pixelgl.Key0) {
		// F000 - FFFF
		var ROM_dump = [4096]byte{}

		// Dump the rom from memory prior to clear all information
		for i := 0 ; i < 4096 ; i ++{
			ROM_dump[i] = Memory[0xF000+i]
		}

		Initialize()

		// Restore ROM to memory
		for i := 0 ; i < 4096 ; i ++{
			Memory[0xF000+i] = ROM_dump[i]
		}

		// Reset PC
		Reset()
		// Restart CPU
		// CPU_Interpreter()

		// Reset graphics
		win.Clear(colornames.Black)

		// Draw Debug Screen
		if Debug {
			startDebug()
		}

		win.Update()

		// Control repetition
		win.UpdateInputWait(time.Second)
	}

	// Pause Key
	if win.Pressed(pixelgl.KeyP) {
		if Pause {
			Pause = false
			fmt.Printf("\t\tPAUSE mode Disabled\n")
			// Control repetition
			win.UpdateInputWait(time.Second)
		} else {
			Pause = true
			dbg_running_opc = true
			fmt.Printf("\t\tPAUSE mode Enabled\n")
			win.UpdateInputWait(time.Second)
		}
	}

	// Step Forward
	if win.Pressed(pixelgl.KeyI) {
		if Pause {
			if Debug {
				for dbg_running_opc == true {
					fmt.Printf("\t\tStep Forward\n")

					win.UpdateInput()
					// Runs the interpreter
					CPU_Interpreter()

					// Update Debug Screen
					if Debug {
						updateDebug()
					}

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					TIA( TIA_Update )

					// Reset Controllers Buttons to 1 (not pressed)
					Memory[SWCHA] = 0xFF //1111 11111
				}

				// After being paused by the end of opcode, set again to start the new one
				dbg_running_opc = true

				// Control repetition
				win.UpdateInputWait(time.Second)
			}

		}

	}




	// Change video resolution
	if win.JustPressed(pixelgl.KeyM) {

		if !Debug {
			// If the mode is smaller than the number of resolutions available increment (-4 to avoid the biggest ones)
			if resolutionCounter < len(settings) -4  {
				resolutionCounter ++
			} else {
				resolutionCounter = 0	// reset resolutionCounter
			}

			activeSetting = &settings[resolutionCounter]

			if isFullScreen {
				win.SetMonitor(activeSetting.Monitor)
			} else {
				win.SetMonitor(nil)
			}
			win.SetBounds(pixel.R(0, 0, float64(activeSetting.Mode.Width), float64(activeSetting.Mode.Height)))

			// Show messages
			if Debug {
				fmt.Printf("\t\tResolution mode[%d]: %dx%d @ %dHz\n",resolutionCounter ,activeSetting.Mode.Width, activeSetting.Mode.Height, activeSetting.Mode.RefreshRate)
			}
			TextMessageStr=fmt.Sprintf("Resolution mode[%d]: %dx%d @ %dHz",resolutionCounter ,activeSetting.Mode.Width, activeSetting.Mode.Height, activeSetting.Mode.RefreshRate)
			ShowMessage = true

			// Update Width and Height values accordingly to new resolutions
			screenWidth	= win.Bounds().W()
			screenHeight	= win.Bounds().H()
			width		= screenWidth/sizeX
			height		= screenHeight/sizeY * sizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

			win.Update()

			CenterWindow()
		} else {
			TextMessageStr = "Cannot change resolution in Debug Mode"
			ShowMessage = true
		}


	}


	// // Fullscreen
	// if win.JustPressed(pixelgl.KeyN) {
	// 	if isFullScreen {
	// 		// Switch to windowed and backup the correct monitor.
	// 		win.SetMonitor(nil)
	// 		isFullScreen = false
	//
	// 		CenterWindow()
	//
	// 		// Show messages
	// 		if Debug {
	// 			fmt.Printf("\n\t\tFullscreen Disabled\n")
	// 		}
	// 		TextMessageStr = "Fullscreen Disabled"
	// 		ShowMessage = true
	//
	// 	} else {
	// 		// Switch to fullscreen.
	// 		win.SetMonitor(activeSetting.Monitor)
	// 		isFullScreen = true
	//
	// 		// Show messages
	// 		if Debug {
	// 			fmt.Printf("\n\t\tFullscreen Enabled\n")
	// 		}
	// 		TextMessageStr = "Fullscreen Enabled"
	// 		ShowMessage = true
	// 	}
	// 	win.SetBounds(pixel.R(0, 0, float64(activeSetting.Mode.Width), float64(activeSetting.Mode.Height)))
	//
	// }


	// -------------- PLAYER 0 -------------- //
	// P0 Right
	if win.Pressed(pixelgl.KeyRight) {
		Memory[SWCHA] = 0x7F // 0111 1111
	}
	// P0 Left
	if win.Pressed(pixelgl.KeyLeft) {
		Memory[SWCHA] = 0xBF // 1011 1111
	}
	// P0 Down
	if win.Pressed(pixelgl.KeyDown) {
		Memory[SWCHA] = 0xDF // 1101 1111
	}
	// P0 Up
	if win.Pressed(pixelgl.KeyUp) {
		Memory[SWCHA] = 0xEF // 1110 1111
	}

	// -------------- PLAYER 1 -------------- //
	// P1 Right
	if win.Pressed(pixelgl.KeyD) {
		Memory[SWCHA] = 0xF7 // 1111 0111
	}
	// P1 Left
	if win.Pressed(pixelgl.KeyA) {
		Memory[SWCHA] = 0xFB // 1111 1011
	}
	// P1 Down
	if win.Pressed(pixelgl.KeyS) {
		Memory[SWCHA] = 0xFD // 1111 1101
	}
	// P1 Up
	if win.Pressed(pixelgl.KeyW) {
		Memory[SWCHA] = 0xFE // 1111 1110
	}
}
