package VGS

import (
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	// "github.com/faiface/pixel/imdraw"

)

func Keyboard(target *pixelgl.Window) {

	// Debug
	if target.JustPressed(pixelgl.Key9) {
		// if Debug {
		// 	stopDebug()
		// } else {
		// 	startDebug()
		// }
		if Debug {
			Debug = false

		} else {
			Debug = true

		}
	}

	// Reset
	if target.Pressed(pixelgl.Key0) {
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
		target.Clear(colornames.Black)

		// Draw Debug Screen
		if Debug {
			startDebug()
		}

		target.Update()

		// Control repetition
		target.UpdateInputWait(time.Second)
	}

	// Pause Key
	if target.Pressed(pixelgl.KeyP) {
		if Pause {
			Pause = false
			fmt.Printf("\t\tPAUSE mode Disabled\n")
			// Control repetition
			target.UpdateInputWait(time.Second)
		} else {
			Pause = true
			dbg_running_opc = true
			fmt.Printf("\t\tPAUSE mode Enabled\n")
			target.UpdateInputWait(time.Second)
		}
	}

	// Step Forward
	if target.Pressed(pixelgl.KeyI) {
		if Pause {
			// if Debug {
				// for dbg_running_opc == true {
					fmt.Printf("\t\tStep Forward\n")

					target.UpdateInput()
					// Runs the interpreter
					if CPU_Enabled {
						CPU_Interpreter()
					}

					// Update Debug Screen
					if Debug {
						updateDebug()
					}

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					TIA( TIA_Update, target )
					target.Update()

				}

				// After being paused by the end of opcode, set again to start the new one
				// dbg_running_opc = true

				// Control repetition
				target.UpdateInputWait(time.Second)
			// }

		// }

	}

	// Breakpoint
	if target.Pressed(pixelgl.KeyO) {
		if Debug {
			if Pause {
				Pause = false
				dbg_break = true
				dbg_break_cycle = counter_F_Cycle + 1000
				fmt.Printf("\t\tBREAKPOINT set to cycle %d\n", counter_F_Cycle+1000)
				// Control repetition
				target.UpdateInputWait(time.Second)
			}
		}
	}


	// Change video resolution
	if target.JustPressed(pixelgl.KeyM) {

		if !Debug {
			// If the mode is smaller than the number of resolutions available increment (-4 to avoid the biggest ones)
			if resolutionCounter < len(settings) -4  {
				resolutionCounter ++
			} else {
				resolutionCounter = 0	// reset resolutionCounter
			}

			activeSetting = &settings[resolutionCounter]

			if isFullScreen {
				target.SetMonitor(activeSetting.Monitor)
			} else {
				target.SetMonitor(nil)
			}
			target.SetBounds(pixel.R(0, 0, float64(activeSetting.Mode.Width), float64(activeSetting.Mode.Height)))

			// Show messages
			if Debug {
				fmt.Printf("\t\tResolution mode[%d]: %dx%d @ %dHz\n",resolutionCounter ,activeSetting.Mode.Width, activeSetting.Mode.Height, activeSetting.Mode.RefreshRate)
			}
			TextMessageStr=fmt.Sprintf("Resolution mode[%d]: %dx%d @ %dHz",resolutionCounter ,activeSetting.Mode.Width, activeSetting.Mode.Height, activeSetting.Mode.RefreshRate)
			ShowMessage = true

			// Update Width and Height values accordingly to new resolutions
			screenWidth	= target.Bounds().W()
			screenHeight	= target.Bounds().H()
			width		= screenWidth/sizeX
			height		= screenHeight/sizeY * sizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

			target.Update()

			// CenterWindow()
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
	if target.Pressed(pixelgl.KeyRight) {
		Memory[SWCHA] &= 0x7F // 0111 1111
	}
	// P0 Left
	if target.Pressed(pixelgl.KeyLeft) {
		Memory[SWCHA] &= 0xBF // 1011 1111
	}
	// P0 Down
	if target.Pressed(pixelgl.KeyDown) {
		Memory[SWCHA] &= 0xDF // 1101 1111
	}
	// P0 Up
	if target.Pressed(pixelgl.KeyUp) {
		Memory[SWCHA] &= 0xEF // 1110 1111
	}

	// -------------- PLAYER 1 -------------- //
	// P1 Right
	if target.Pressed(pixelgl.KeyD) {
		Memory[SWCHA] &= 0xF7 // 1111 0111
	}
	// P1 Left
	if target.Pressed(pixelgl.KeyA) {
		Memory[SWCHA] &= 0xFB // 1111 1011
	}
	// P1 Down
	if target.Pressed(pixelgl.KeyS) {
		Memory[SWCHA] &= 0xFD // 1111 1101
	}
	// P1 Up
	if target.Pressed(pixelgl.KeyW) {
		Memory[SWCHA] &= 0xFE // 1111 1110
	}
}
