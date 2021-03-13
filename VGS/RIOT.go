package VGS

import (
	"os"
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	// "github.com/faiface/pixel/imdraw"

)

// ----------------------------------- Keyboard ----------------------------------- //

func Keyboard(target *pixelgl.Window) {


	// Console switches - Reset
	if target.Pressed(pixelgl.Key8) {

		// Set Reset switch (put 0 on position 0 of SWCHB)
		Memory[SWCHB] &^= (1 << 0)
		fmt.Printf("\tReset - Console Switch\n")
	}

	// Console switches - Game Select
	if target.Pressed(pixelgl.Key7) {

		// Set Game Select switch (put 0 on position 1 of SWCHB)
		Memory[SWCHB] &^= (1 << 1)
		fmt.Printf("\tGame Select - Console Switch\n")
	}

	// Console switches - Game Select
	if target.JustPressed(pixelgl.Key6) {

		// Test if bit 3 is set or not (val = 0 no, val = 1 yes)
		bit := Memory[SWCHB] & (1 << 3)

		if bit > 0 {
			// Set Game Select switch (put 0 on position 3 of SWCHB)
			Memory[SWCHB] &^= (1 << 3)
			fmt.Printf("\tColor mode - Console Switch\n")
		} else {
			// Disable Game Select switch (put 1 on position 3 of SWCHB)
			Memory[SWCHB] |= (1 << 3)
			fmt.Printf("\tBlack & White mode - Console Switch\n")
		}

	}

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

					// Draw and update every time for debug
					imd.Draw(target)
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
	// P0 Button
	if target.Pressed(pixelgl.KeySpace) {
		// Memory[INPT4] &= 0x7F // 0111 1111
		update_Memory_TIA_RO(INPT4, Memory[INPT4] & 0x7F)
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
	// P1 Button
	if target.Pressed(pixelgl.KeyX) {
		// Memory[INPT5] &= 0x7F // 0111 1111
		update_Memory_TIA_RO(INPT5, Memory[INPT5] & 0x7F)
	}
}

// ------------------------------------- Timer ------------------------------------ //


func riot_update_timer(addr uint16) {

	// ------------ Update the Timer ------------ //

	// TIM1T
	if addr == 0x294 {
		// Set Timer
		riot_timer = Memory_RIOT_RW[ addr - 0x280]
		// Set Multiplier
		riot_timer_mult = 1
		// Clear TIMINT (interrupt flag)
		for i := 0 ; i < 4 ; i++ {
			// Address and Mirrors
			Memory[ ( TIMINT - 0x280 ) + uint16(i) * 8 ] = 0
			// Address + 2 and Mirrors
			Memory[ ( TIMINT - 0x280 + 2) + uint16(i) * 8 ] = 0
			// fmt.Println( ( TIMINT - 0x280 + 2) + uint16(i) * 8 )
		}
		// Reset the internal cycle counter
		riot_timer_counter = 0
		// Update Timer Output
		Memory[INTIM] = riot_timer

	// TIM8T
	} else if addr == 0x295 {
		// Set Timer
		riot_timer = Memory_RIOT_RW[ addr - 0x280]
		// Set Multiplier
		riot_timer_mult = 8
		// Clear TIMINT (interrupt flag)
		for i := 0 ; i < 4 ; i++ {
			// Address and Mirrors
			Memory[ ( TIM8T - 0x280 ) + uint16(i) * 8 ] = 0
			// Address + 2 and Mirrors
			Memory[ ( TIM8T - 0x280 + 2) + uint16(i) * 8 ] = 0
			// fmt.Println( ( TIMINT - 0x280 + 2) + uint16(i) * 8 )
		}
		// Reset the internal cycle counter
		riot_timer_counter = 0
		// Update Timer Output
		Memory[INTIM] = riot_timer

	// TIM64T
	} else if addr == 0x296 {
		// Set Timer
		riot_timer = Memory_RIOT_RW[ addr - 0x280]
		// Set Multiplier
		riot_timer_mult = 64
		// Clear TIMINT (interrupt flag)
		for i := 0 ; i < 4 ; i++ {
			// Address and Mirrors
			Memory[ ( TIM64T - 0x280 ) + uint16(i) * 8 ] = 0
			// Address + 2 and Mirrors
			Memory[ ( TIM64T - 0x280 + 2) + uint16(i) * 8 ] = 0
			// fmt.Println( ( TIMINT - 0x280 + 2) + uint16(i) * 8 )
		}
		// Reset the internal cycle counter
		riot_timer_counter = 0
		// Update Timer Output
		Memory[INTIM] = riot_timer

	// T1024T
	} else if addr == 0x297 {
		// Set Timer
		riot_timer = Memory_RIOT_RW[ addr - 0x280]
		// Set Multiplier
		riot_timer_mult = 1024
		// Clear TIMINT (interrupt flag)
		for i := 0 ; i < 4 ; i++ {
			// Address and Mirrors
			Memory[ ( T1024T - 0x280 ) + uint16(i) * 8 ] = 0
			// Address + 2 and Mirrors
			Memory[ ( T1024T - 0x280 + 2) + uint16(i) * 8 ] = 0
		}
		// Reset the internal cycle counter
		riot_timer_counter = 0
		// Update Timer Output
		Memory[INTIM] = riot_timer

	} else {
		fmt.Printf("\nriot_update_timer() - Memory address not mapped: %02X! Exiting!\n", addr)
		os.Exit(2)
	}

	// fmt.Printf("Timer Set to: %d\t Multiplier: %d\n", riot_timer, riot_timer_mult)
	// os.Exit(2)


	// &&  memAddr != 0x295 && memAddr != 0x296 && memAddr != 0x297

}
