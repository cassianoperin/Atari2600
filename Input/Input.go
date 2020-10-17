package Input

import (
	"fmt"
	"time"
	"Atari2600/CPU"
	"Atari2600/Global"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Keyboard() {

	// Debug
	if Global.Win.JustPressed(pixelgl.Key9) {
		if CPU.Debug {
			CPU.Debug = false
			Global.SizeYused = 1.0
			// Show messages
			if CPU.Debug {
				fmt.Printf("\t\tDEBUG mode Disabled\n")
			}
			Global.TextMessageStr = "DEBUG mode Disabled"
			Global.ShowMessage = true

			// Update Width and Height values accordingly to new resolutions
			Global.ScreenWidth	= Global.Win.Bounds().W()
			Global.ScreenHeight	= Global.Win.Bounds().H()
			Global.Width		= Global.ScreenWidth/Global.SizeX
			Global.Height		= Global.ScreenHeight/Global.SizeY * Global.SizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

			Global.Win.Update()
		} else {
			CPU.Debug = true

			InitializeDebug()
		}
	}

	// Reset
	if Global.Win.Pressed(pixelgl.Key0) {
		// F000 - FFFF
		// var ROM_dump = [4096]byte{}
		//
		// // Dump the rom from memory prior to clear all information
		// for i := 0 ; i < 4096 ; i ++{
		// 	ROM_dump[i] = CPU.Memory[0xF000+i]
		// }
		//
		// // Workaround for WSYNC before VSYNC
		// Global.VSYNC_passed = false
		//
		// CPU.Initialize()
		//
		// // Restore ROM to memory
		// for i := 0 ; i < 4096 ; i ++{
		// 	CPU.Memory[0xF000+i] = ROM_dump[i]
		// }
		//
		// // Reset graphics
		// //renderGraphics()
		// // Restart Draw from the beginning
		// line = 1
		//
		// // Players Vertical Positioning
		// XPositionP0		= 0
		// XFinePositionP0	= 0
		// XPositionP1		= 0
		// XFinePositionP1	= 0
		//
		// // ------------------ Personal Control Flags ------------------ //
		// CPU.Beam_index	= 0		// Beam index to control where to draw objects using cpu cycles
		//
		// CPU.Reset()
	}

	// CPU.Pause Key
	if Global.Win.Pressed(pixelgl.KeyP) {
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
	if Global.Win.Pressed(pixelgl.KeyI) {
		// if CPU.Pause {
		// 	fmt.Printf("\t\tStep Forward\n")
		//
		// 	// Runs the interpreter
		// 	CPU.Interpreter()
		//
		// 	// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
		// 	TIA( CPU.TIA_Update )
		//
		// 	// Reset Controllers Buttons to 1 (not pressed)
		// 	CPU.Memory[CPU.SWCHA] = 0xFF //1111 11111
		//
		// 	// Draw Debug Screen
		// 	if CPU.Debug {
		// 		// Background
		// 		drawDebugScreen(imd)
		// 		// Info
		// 		drawDebugInfo()
		// 	}
		//
		// 	time.Sleep(50 * time.Millisecond)
		// }

	}




	// Change video resolution
	if Global.Win.JustPressed(pixelgl.KeyM) {

		if !CPU.Debug {
			// If the mode is smaller than the number of resolutions available increment (-4 to avoid the biggest ones)
			if Global.ResolutionCounter < len(Global.Settings) -4  {
				Global.ResolutionCounter ++
			} else {
				Global.ResolutionCounter = 0	// reset Global.ResolutionCounter
			}

			Global.ActiveSetting = &Global.Settings[Global.ResolutionCounter]

			if Global.IsFullScreen {
				Global.Win.SetMonitor(Global.ActiveSetting.Monitor)
			} else {
				Global.Win.SetMonitor(nil)
			}
			Global.Win.SetBounds(pixel.R(0, 0, float64(Global.ActiveSetting.Mode.Width), float64(Global.ActiveSetting.Mode.Height)))

			// Show messages
			if CPU.Debug {
				fmt.Printf("\t\tResolution mode[%d]: %dx%d @ %dHz\n",Global.ResolutionCounter ,Global.ActiveSetting.Mode.Width, Global.ActiveSetting.Mode.Height, Global.ActiveSetting.Mode.RefreshRate)
			}
			Global.TextMessageStr=fmt.Sprintf("Resolution mode[%d]: %dx%d @ %dHz",Global.ResolutionCounter ,Global.ActiveSetting.Mode.Width, Global.ActiveSetting.Mode.Height, Global.ActiveSetting.Mode.RefreshRate)
			Global.ShowMessage = true

			// Update Width and Height values accordingly to new resolutions
			Global.ScreenWidth	= Global.Win.Bounds().W()
			Global.ScreenHeight	= Global.Win.Bounds().H()
			Global.Width		= Global.ScreenWidth/Global.SizeX
			Global.Height		= Global.ScreenHeight/Global.SizeY * Global.SizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

			Global.Win.Update()

			Global.CenterWindow()
		} else {
			Global.TextMessageStr = "Cannot change resolution in Debug Mode"
			Global.ShowMessage = true
		}


	}


	// // Fullscreen
	// if Global.Win.JustPressed(pixelgl.KeyN) {
	// 	if Global.IsFullScreen {
	// 		// Switch to windowed and backup the correct monitor.
	// 		Global.Win.SetMonitor(nil)
	// 		Global.IsFullScreen = false
	//
	// 		Global.CenterWindow()
	//
	// 		// Show messages
	// 		if CPU.Debug {
	// 			fmt.Printf("\n\t\tFullscreen Disabled\n")
	// 		}
	// 		Global.TextMessageStr = "Fullscreen Disabled"
	// 		Global.ShowMessage = true
	//
	// 	} else {
	// 		// Switch to fullscreen.
	// 		Global.Win.SetMonitor(Global.ActiveSetting.Monitor)
	// 		Global.IsFullScreen = true
	//
	// 		// Show messages
	// 		if CPU.Debug {
	// 			fmt.Printf("\n\t\tFullscreen Enabled\n")
	// 		}
	// 		Global.TextMessageStr = "Fullscreen Enabled"
	// 		Global.ShowMessage = true
	// 	}
	// 	Global.Win.SetBounds(pixel.R(0, 0, float64(Global.ActiveSetting.Mode.Width), float64(Global.ActiveSetting.Mode.Height)))
	//
	// }








	// -------------- PLAYER 0 -------------- //
	// P0 Right
	if Global.Win.Pressed(pixelgl.KeyRight) {
		CPU.Memory[CPU.SWCHA] = 0x7F // 0111 1111
	}
	// P0 Left
	if Global.Win.Pressed(pixelgl.KeyLeft) {
		CPU.Memory[CPU.SWCHA] = 0xBF // 1011 1111
	}
	// P0 Down
	if Global.Win.Pressed(pixelgl.KeyDown) {
		CPU.Memory[CPU.SWCHA] = 0xDF // 1101 1111
	}
	// P0 Up
	if Global.Win.Pressed(pixelgl.KeyUp) {
		CPU.Memory[CPU.SWCHA] = 0xEF // 1110 1111
	}

	// -------------- PLAYER 1 -------------- //
	// P1 Right
	if Global.Win.Pressed(pixelgl.KeyD) {
		CPU.Memory[CPU.SWCHA] = 0xF7 // 1111 0111
	}
	// P1 Left
	if Global.Win.Pressed(pixelgl.KeyA) {
		CPU.Memory[CPU.SWCHA] = 0xFB // 1111 1011
	}
	// P1 Down
	if Global.Win.Pressed(pixelgl.KeyS) {
		CPU.Memory[CPU.SWCHA] = 0xFD // 1111 1101
	}
	// P1 Up
	if Global.Win.Pressed(pixelgl.KeyW) {
		CPU.Memory[CPU.SWCHA] = 0xFE // 1111 1110
	}
}


func InitializeDebug() {
	Global.Win.Clear(colornames.Black)
	Global.SizeYused = 0.3
	// Show messages
	if CPU.Debug {
		fmt.Printf("\t\tDEBUG mode Enabled\n")
	}
	// Global.Win.Clear(colornames.Black)
	Global.TextMessageStr = "DEBUG mode Enabled"
	Global.ShowMessage = true

	// Set Initial resolution
	Global.ActiveSetting = &Global.Settings[3]

	if Global.IsFullScreen {
		Global.Win.SetMonitor(Global.ActiveSetting.Monitor)
	} else {
		Global.Win.SetMonitor(nil)
	}
	Global.Win.SetBounds(pixel.R(0, 0, float64(Global.ActiveSetting.Mode.Width), float64(Global.ActiveSetting.Mode.Height)))

	// Update Width and Height values accordingly to new resolutions
	Global.ScreenWidth	= Global.Win.Bounds().W()
	Global.ScreenHeight	= Global.Win.Bounds().H()
	Global.Width		= Global.ScreenWidth/Global.SizeX
	Global.Height		= Global.ScreenHeight/Global.SizeY * Global.SizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

	Global.Win.Update()
}
