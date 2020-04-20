package Graphics

import (
	"fmt"
	// "os"
	// "strconv"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"Atari2600/Palettes"
	"Atari2600/CPU"
	"image/color"

)

var (
	// Window Configuration
	win				* pixelgl.Window
	imd				= imdraw.New(nil)
	cfg				= pixelgl.WindowConfig{}

	// Line draw control
	line				int = 1
	line_max			int = 262

	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]byte			//Improve to binary
	pixelSize			float64 = 4.0		// 80 lines (half screen) / 20 PF0, PF1 and PF2 bits

	// FPS count
	frames			= 0
	draws			= 0

	// Debug mode
	debug			bool = false

)

const (
	sizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY			float64	= 192.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	// screenWidth		= float64(sizeX*3)
	// screenHeight		= float64(sizeY*1.5)
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


func readPF0() {
	// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", CPU.Memory[PF0])
	for i := 4 ; i < 8 ; i++ {
		playfield[i-4] = ( CPU.Memory[CPU.PF0] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)

}


func readPF1() {
	// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", CPU.Memory[PF1])
	for i := 0 ; i < 8 ; i++ {
		playfield[11-i] = ( CPU.Memory[CPU.PF1] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)
}


func readPF2() {
	// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", CPU.Memory[PF2])
	for i := 0 ; i < 8 ; i++ {
		playfield[12+i] = ( CPU.Memory[CPU.PF2] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)
}



func drawPlayer0() {
	if CPU.DrawP0 {



		// If a program doesnt use RESP0, initialize
		if CPU.XPositionP0 == 0 {
			CPU.XPositionP0 = 23
		}

		if debug {
			fmt.Printf("\nLine: %d\tGRP0: %08b\tXPositionP0: %d\tXFinePositionP0: %d", line, CPU.Memory[CPU.GRP0], CPU.XPositionP0, CPU.XFinePositionP0)
		}
		// CPU.Pause = true

		for i:=0 ; i <=7 ; i++{
			bit := CPU.Memory[CPU.GRP0] >> (7-byte(i)) & 0x01

			if bit == 1 {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP0])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				imd.Push(pixel.V(  (float64(   (CPU.XPositionP0*3)-68 +byte(i))  +float64(CPU.XFinePositionP0) )*width			, float64(232-line)*height ))
				imd.Push(pixel.V(  (float64(   (CPU.XPositionP0*3)-68 +byte(i))  +float64(CPU.XFinePositionP0) )*width + width	, float64(232-line)*height + height))
				imd.Rectangle(0)

				imd.Draw(win)
				// Count draw operations number per second
				draws ++

				// CPU.Pause = true
			}
		}
		CPU.DrawP0 = false
	}
}


func drawPlayer1() {
	if CPU.DrawP1 {

		// If a program doesnt use RESP0, initialize
		if CPU.XPositionP1 == 0 {
			CPU.XPositionP1 = 30
		}

		if debug {
			fmt.Printf("\nLine: %d\tGRP1: %08b\tXPositionP1: %d\tHMP1: %d", line, CPU.Memory[CPU.GRP1], CPU.XPositionP1, CPU.Memory[CPU.HMP1])
		}
		// CPU.Pause = true

		for i:=0 ; i <=7 ; i++{
			bit := CPU.Memory[CPU.GRP1] >> (7-byte(i)) & 0x01

			if bit == 1 {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP1])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				imd.Push(pixel.V(  (float64(   (CPU.XPositionP1*3)-68 +byte(i))  +float64(CPU.XFinePositionP1) )*width			, float64(232-line)*height ))
				imd.Push(pixel.V(  (float64(   (CPU.XPositionP1*3)-68 +byte(i))  +float64(CPU.XFinePositionP1) )*width + width	, float64(232-line)*height + height))
				imd.Rectangle(0)

				imd.Draw(win)
				// Count draw operations number per second
				draws ++

				// CPU.Pause = true
			}
		}
		CPU.DrawP1 = false
	}
}


// Old Draw Player
// func drawPlayer1() {
// 	if CPU.DrawP1 {
// 		// fmt.Printf("\nLine: %d\tGRP1: %08b\n", line, CPU.Memory[GRP1])
//
// 		for i:=0 ; i <=7 ; i++{
// 			bit := CPU.Memory[GRP1] >> (7-byte(i)) & 0x01
//
// 			if bit == 1 {
// 				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
// 				R, G, B := Palettes.NTSC(CPU.Memory[COLUP1])
// 				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
//
// 				imd.Push(pixel.V(  (float64(CPU.Memory[RESP1]*3+byte(i)) )*width			, float64(232-line)*height ))
// 				imd.Push(pixel.V(  (float64(CPU.Memory[RESP1]*3+byte(i)) )*width + width	, float64(232-line)*height +height))
// 				imd.Rectangle(0)
//
// 				imd.Draw(win)
// 				// Count draw operations number per second
// 				draws ++
// 			}
// 		}
// 		CPU.DrawP1 = false
// 	}
// }

func drawGraphics() {

	imd	= imdraw.New(nil)

	// Draw conten on every WSYNC from CPU
	if CPU.DrawLine {
		// 3 lines VSYNC
		if CPU.Memory[CPU.VBLANK] == 2 && CPU.Memory[CPU.VSYNC] == 2  {
			if debug {
				fmt.Printf("\nLine: %d\tVSYNC: %02X", line, CPU.Memory[CPU.VSYNC])
			}


		// 37 lines VBLANK
		} else if CPU.Memory[CPU.VBLANK] == 2 {
			if debug {
				fmt.Printf("\nLine: %d\tVBLANK: %02X", line, CPU.Memory[CPU.VBLANK])
			}


		// 192 Visible Area
		} else if line <= 232 {
			if debug {
				fmt.Printf("\nLine: %d\tVisible Area: %d", line, line-40)
			}

			readPF0()
			readPF1()
			readPF2()

			drawVisibleModeLine()

			// DRAW PLAYER 0
			if CPU.DrawP0 {
				drawPlayer0()

				CPU.DrawP0 = false
			}

			// DRAW PLAYER 1
			if CPU.DrawP1 {
				drawPlayer1()

				CPU.DrawP1 = false
			}


		// Overscan
		} else {
			if debug {
				fmt.Printf("\nLine: %d\tOVERSCAN", line)
			}
		}

		line ++
		CPU.DrawLine = false

	}



	select {
	case <-CPU.ScreenRefresh.C:
	// When ticker run (60 times in a second, check de DelayTimer)

		win.Update()
		// frames++
		default:
			// No timer to handle
	}


}



func drawVisibleModeLine() {

	// D0 = 1 = Reflect first 20 sprites of the PF to the last 20
	if (CPU.Memory[CPU.CTRLPF] & 0x01) == 1 {
		j := 0
		for i := len(playfield) - 1; i > 19  ; i-- {
			playfield[i] = playfield[j]
			j++
		}
	// Duplicate last 20 sprites with first 20
	}  else {
		for i := 20 ; i < len(playfield) ; i++ {
			playfield[i] = playfield[i-20]
		}
	}

	// Value that Im looking for repetitions
	search := playfield[0]
	// Where to draw
	index := 0
	count := 1

	for i := 0 ; i < len(playfield) -1; i++ {

		if playfield[i+1] == search {
			// fmt.Printf("\nI: %d\tRepeated Number\n",i)
			count++
		} else {
			// Set color (0: Background | 1: Playfield)
			if search == 0 {
				// READ COLUBK (Memory[0x09]) - Set the Background Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUBK])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			} else {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUPF])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			}


			// If it is rendering the playfield
			if search == 1 {
				// Check D1 status to use color of players in the score

				if (CPU.Memory[CPU.CTRLPF] & 0x02) >> 1 == 1  {
					// READ COLUP0 (Memory[0x06]) - Set the Player 0 Color (On Score)
					R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP0])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					// Set P1 Color
					if i >= 20 {
						// READ COLUP1 (Memory[0x07]) - Set the Player 1 Color (On Score)
						R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP1])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					}
				}

			}

			// Draw
			//fmt.Printf("\ni: %d\tIndex: %d\tNumber of repeated %d: %d\n", i, index, search,count)
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width																			, float64(232-line)*height ))
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height + height))
			// fmt.Printf("%f %f", (68 + (float64(index) *5)),	(68) + (float64(index) *5) +float64(count*5) )
			imd.Rectangle(0)
			count = 1
			index = i+1
			search = playfield[i+1]
		}
	}

	// Process the last value [19]
	if search == 0 {
		// READ COLUBK (Memory[0x09]) - Set the Background Color
		R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUBK])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
	} else {
		// READ COLUPF (Memory[0x08]) - Set the Playfield Color
		R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUPF])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
	}

	//fmt.Printf("\ni: 19\tIndex: %d\tNumber of repeated %d: %d\n",  index, search,count)
	imd.Push(pixel.V(  (float64(index) *pixelSize)*width								, float64(232-line)*height ))
	imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height + height))
	//fmt.Printf("%f %f", (68 + (float64(index) *5)) ,	(68) + (float64(index) *5) +float64(count*5) )
	imd.Rectangle(0)

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}


func keyboard() {

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
			CPU.Interpreter()
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
				CPU.Interpreter()
				// CPU.Flags_V_SBC(5,15)

				// Reset Controllers Buttons to 1 (not pressed)
				CPU.Memory[CPU.SWCHA] = 0xFF //1111 11111
			}
			// DRAW

			// When finished drawing the screen, reset and start a new frame
			if line == line_max + 1 {
				if debug {
					fmt.Printf("\nFinished the screen height, start a new frame.\n")
				}
				line = 1
				frames ++
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

		drawGraphics()
		// fmt.Printf("%d",CPU.DecodeTwoComplement(246))

	}

}
