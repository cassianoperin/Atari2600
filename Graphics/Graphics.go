package Graphics

import (
	"fmt"
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

	// Debug mode
	debug			bool = false
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


// Infinte Loop
func Run() {

	CD_P0_P1		= [160]byte{}

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
				// Runs the interpreter
				CPU.Interpreter()

				// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
				TIA( CPU.TIA_Update )

				// Reset Controllers Buttons to 1 (not pressed)
				CPU.Memory[CPU.SWCHA] = 0xFF //1111 11111
			}

			select {
				case <-CPU.Second: // Second
					win.SetTitle(fmt.Sprintf("%s IPS: %d| FPS: %d | Draws: %d", cfg.Title, CPU.IPS, frames, draws))
					CPU.IPS = 0
					frames = 0
					draws  = 0
				default:
			}

			default:
				// No timer to handle
		}

		select {
		case <-CPU.ScreenRefresh.C:
		// When ticker run (60 times in a second, Refresh the screen)

			win.Update()
			// frames++
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

	}

}
