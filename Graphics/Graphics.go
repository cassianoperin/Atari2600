package Graphics

import (
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"Atari2600/CPU"
)

var (
	// Window Configuration
	win				* pixelgl.Window
	imd				= imdraw.New(nil)
	cfg				= pixelgl.WindowConfig{}

	// // Line draw control
	// line				int = 1
	// line_max			int = 262

	// // PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	// playfield			[40]byte			//Improve to binary
	// pixelSize			float64 = 4.0		// 80 lines (half screen) / 20 PF0, PF1 and PF2 bits
	//
	// // FPS count
	// frames			= 0
	// draws			= 0
	//
	// // Workaround to avoid  WSYNC before VSYNC
	// VSYNC_passed		bool = false

	// Debug mode
	debug			bool = false

	// old_BeamIndex	byte = 0	// Used to draw the beam updates every cycle on the CRT
	// visibleArea		bool		// Not used yet, but will be used to just draw in visible area
)

const (
	sizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY			float64	= 192.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	// screenWidth		= float64(sizeX*3)
	// screenHeight	= float64(sizeY*1.5)
	screenWidth		= float64(sizeX*6)
	screenHeight		= float64(sizeY*3)
	width			= screenWidth  / sizeX
	height			= screenHeight / sizeY

)


func renderGraphics() {
	cfg = pixelgl.WindowConfig{
		Title:  "Atari 2600",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  false,
	}
	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
}




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
		// Draw instuctions
		CPU.DrawP0		= false	// Instruct Graphics to draw Player 0 sprite
		CPU.DrawP1		= false	// Instruct Graphics to draw Player 1 sprite


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
			fmt.Printf("\t\tStep Forward - MISSING IMPLEMENT TIA UPDATE STA like in RUN()\n")


			switch CPU.Memory[CPU.PC] {

				// Zeropage: STX, STA, STY
				case 0x86, 0x85, 0x84:

					CPU.Beam_index += 3
					// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

					memAddr, mode := CPU.Addr_mode_Zeropage(CPU.PC+1)
					_ = mode	// not used

					if memAddr < 128 {
						CPU.TIA_Update = int8(memAddr)
					}

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					CRT( CPU.TIA_Update )

					// Reset to default value
					CPU.TIA_Update = -1

					// Runs the interpreter
					CPU.Interpreter()

				// Zeropage,X: STA
				case 0x95:

					CPU.Beam_index += 4
					// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

					memAddr, mode := CPU.Addr_mode_ZeropageX(CPU.PC+1)
					_ = mode	// not used

					if memAddr < 128 {
						CPU.TIA_Update = int8(memAddr)
					}

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					CRT( CPU.TIA_Update )

					// Reset to default value
					CPU.TIA_Update = -1

					// Runs the interpreter
					CPU.Interpreter()

				// Zeropage,X: STA
			case 0x99:

					CPU.Beam_index += 5
					// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

					memAddr, mode := CPU.Addr_mode_AbsoluteY(CPU.PC+1)
					_ = mode	// not used

					if memAddr < 128 {
						CPU.TIA_Update = int8(memAddr)
					}

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					CRT( CPU.TIA_Update )

					// Reset to default value
					CPU.TIA_Update = -1

					// Runs the interpreter
					CPU.Interpreter()

				default:

					// Runs the interpreter
					CPU.Interpreter()

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					CRT( CPU.TIA_Update )

					// Reset to default value
					CPU.TIA_Update = -1

			}


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


// Infinte Loop
func Run() {

	// imd = imdraw.New(nil)

	// Set up render system
	renderGraphics()


	// Main Infinite Loop
	for !win.Closed() {

		// Esc to quit program
		if win.Pressed(pixelgl.KeyEscape) {
			break
		}




		// Every Cycle Control the clock!!!
		select {
		case <-CPU.Clock.C:

			// Handle Input
			keyboard()

			if !CPU.Pause {
				// Call a CPU Cycle


				// MAP STA, STX and STY that needs to first increment the beamer for correctly TIA rendering
				// The opcode spends 2 or 3 cycles to update Memory (TIA NEEDS TO DRAW THIS cycles) prior to use the updated value
				// Ex.: If updated COLUBK, TIA needs to draw the TIA color cycles with the current Background color, and after this, can use the new
				switch CPU.Memory[CPU.PC] {

					// Zeropage: STX, STA, STY
					case 0x86, 0x85, 0x84:

						CPU.Beam_index += 3
						// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

						memAddr, mode := CPU.Addr_mode_Zeropage(CPU.PC+1)
						_ = mode	// not used

						if memAddr < 128 {
							CPU.TIA_Update = int8(memAddr)
						}

						// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
						CRT( CPU.TIA_Update )

						// Reset to default value
						CPU.TIA_Update = -1

						// Runs the interpreter
						CPU.Interpreter()

					// Zeropage,X: STA
					case 0x95:

						CPU.Beam_index += 4
						// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

						memAddr, mode := CPU.Addr_mode_ZeropageX(CPU.PC+1)
						_ = mode	// not used

						if memAddr < 128 {
							CPU.TIA_Update = int8(memAddr)
						}

						// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
						CRT( CPU.TIA_Update )

						// Reset to default value
						CPU.TIA_Update = -1

						// Runs the interpreter
						CPU.Interpreter()

					// Zeropage,X: STA
				case 0x99:

						CPU.Beam_index += 5
						// fmt.Printf("Opcode: %02X\n",CPU.Opcode)

						memAddr, mode := CPU.Addr_mode_AbsoluteY(CPU.PC+1)
						_ = mode	// not used

						if memAddr < 128 {
							CPU.TIA_Update = int8(memAddr)
						}

						// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
						CRT( CPU.TIA_Update )

						// Reset to default value
						CPU.TIA_Update = -1

						// Runs the interpreter
						CPU.Interpreter()

					default:

						// Runs the interpreter
						CPU.Interpreter()

						// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
						CRT( CPU.TIA_Update )

						// Reset to default value
						CPU.TIA_Update = -1

				}

				// Reset Controllers Buttons to 1 (not pressed)
				CPU.Memory[CPU.SWCHA] = 0xFF //1111 11111

			}


			select {
				case <-CPU.Second: // Second
					win.SetTitle(fmt.Sprintf("%s | FPS: %d | Draws: %d", cfg.Title, frames, draws))
					frames = 0
					draws  = 0
				default:
			}

			default:
				// No timer to handle
		}




		// -------------------------- SBC, CARRY and Overflow tests -------------------------- //
		// CPU.P[0] = 1

		// CPU.A = 135
		// CPU.Memory[61832] = 15

		// Test 1
		// CPU.A = 0xF8
		// CPU.Memory[61832] = 0x0A
		// Test 2
		// CPU.A = 0x81
		// CPU.Memory[61832] = 0x07
		// Test 3
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x02
		// Test 4
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0xFE // Hexadecimal de -2 (twos)
		// Test 5
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x09
		// Test 6
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x90
		// Test 7
		// CPU.A = 0x10
		// CPU.Memory[61832] = 0x90
		// Test 8
		// CPU.A = 0x10
		// CPU.Memory[61832] = 0x91

		// CPU.A = 80
		// CPU.Memory[61832] = 240
		// CPU.A = 80
		// CPU.Memory[61832] = 176
		// CPU.A = 208
		// CPU.Memory[61832] = 48
		//
		// fmt.Printf("\n\nOPERATION: A: %d (0x%02X) - %d (0x%02X)", CPU.A, CPU.A, CPU.Memory[61832], CPU.Memory[61832])
		// CPU.PC = 0xF187
		// CPU.Interpreter()
		// fmt.Printf("\n\nResult: %d (0x%02X)\tTwo's Complement: %d (0x%02X)\tOverflow: %d\tCarry: %d\n\n", CPU.A, CPU.A, CPU.DecodeTwoComplement(CPU.A), CPU.DecodeTwoComplement(CPU.A), CPU.P[6], CPU.P[0])
		//
		// os.Exit(2)
		// ----------------------------------------------------------------------------------- //


		select {
		case <-CPU.ScreenRefresh.C:
		// When ticker run (60 times in a second, Refresh the screen)

			win.Update()
			// frames++
			default:
				// No timer to handle
		}

	}

}
