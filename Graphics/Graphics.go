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

	// Emulate CRT Electron beam
	beam_index		int = 0

	//0000-002C - TIA (write)
	VSYNC 			byte = 0x00		//Vertical Sync Set-Clear
	VBLANK			byte = 0x01		//Vertical Blank Set-Clear
	WSYNC			byte = 0x02		//Wait for Horizontal Blank
	RSYNC			byte = 0x03		//Reset Horizontal Sync Counter
	COLUP0			byte = 0x06
	COLUP1			byte = 0x07
	COLUPF			byte	= 0x08
	COLUBK			byte	= 0x09
	PF0 				byte	= 0x0D		//Playfield Register Byte 0
	PF1 				byte	= 0x0E		//Playfield Register Byte 1
	PF2 				byte	= 0x0F		//Playfield Register Byte 2
	// GRP0				byte = 0			// Graphic Player 0 position
	// GRP1				byte = 0			// Graphic Player 1 position
	// Flag used to put different colors in scores
	drawing_score		bool = false
	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]byte			//Improve to binary
	pixelSize			float64 = 4.0

	// CTRLPLF (8 bits register)
	// D0 = Reflect, false = Repeat
	D0_Reflect			bool = false
	// D1 = Score == Color of the score will be the same as player
	D1_Score				bool = true
	// D2 = Priority == Player behind the playfield
	// D4-5 = Ball Size (1, 2, 4, 8)

	// FPS count
	frames			= 0
	draws			= 0

	// Draw 3 pixels per cycle or in optimized mode
	draw_mode_hw		bool = true

	// Debug mode
	debug			bool = false


	// Flag used to dont draw vblank before VSYNC
	vsync_started		bool = false

)

const (
	sizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY			float64	= 192.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	screenWidth		= float64(sizeX*3)
	screenHeight		= float64(sizeY*1.5)
	width			= screenWidth  / sizeX
	height			= screenHeight / sizeY
)

func renderGraphics() {
	cfg = pixelgl.WindowConfig{
		Title:  "NTSC CRT TV Emulator",
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
		playfield[i-4] = ( CPU.Memory[PF0] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)

}

func readPF1() {
	// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", CPU.Memory[PF1])
	for i := 0 ; i < 8 ; i++ {
		playfield[11-i] = ( CPU.Memory[PF1] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)
}

func readPF2() {
	// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", CPU.Memory[PF2])
	for i := 0 ; i < 8 ; i++ {
		playfield[12+i] = ( CPU.Memory[PF2] >> byte(i) ) & 0x01
	}
	// fmt.Printf("%d", playfield)
}


func drawGraphics() {

	imd	= imdraw.New(nil)

	// Draw conten on every WSYNC from CPU
	if CPU.DrawVSYNC {
		// 3 lines VSYNC
		if CPU.Memory[VBLANK] == 2 && CPU.Memory[VSYNC] == 2  {
			if debug {
				fmt.Printf("\nLine: %d\tVSYNC: %02X", line, CPU.Memory[VSYNC])
			}
			line ++
			CPU.DrawVSYNC = false
		// 37 lines VBLANK
		} else if CPU.Memory[VBLANK] == 2 {
			if debug {
				fmt.Printf("\nLine: %d\tVBLANK: %02X", line, CPU.Memory[VBLANK])
			}
			line ++
			CPU.DrawVSYNC = false
		// 192 Visible Area
		} else {
			if debug {
				fmt.Printf("\nLine: %d\tVisible Area: %d", line, line-40)
			}

			// R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
			// imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			//
			// imd.Push(pixel.V(	0				, float64(232-line)*height ))
			// imd.Push(pixel.V(	win.Bounds().W()	, float64(232-line)*height ))
			// imd.Line(height)
			// imd.Draw(win)
			// draws ++

			readPF0()
			readPF1()
			readPF2()
			drawVisibleModeLine()


			CPU.DrawVSYNC = false
		}


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


func Run() {

	//imd := imdraw.New(nil)

	// Set up render system
	renderGraphics()



	// Main Infinite Loop
	for !win.Closed() {

		// Esc to quit program
		if win.Pressed(pixelgl.KeyEscape) {
			break
		}

		// CPU.Pause Key
		if win.Pressed(pixelgl.KeyP) {
			if CPU.Pause {
				CPU.Pause = false
				fmt.Printf("\t\tPAUSE mode Disabled\n")
			 	// drawGraphics()
				// win.Update()
				time.Sleep(500 * time.Millisecond)
			} else {
				CPU.Pause = true
				fmt.Printf("\t\tPAUSE mode Enabled\n")
				time.Sleep(500 * time.Millisecond)
			}
		}

		// Every Cycle Control the clock!!!
		select {
		case <-CPU.Clock.C:
			//fmt.Printf("CPU Cycle: %d\n", cycle)
			//fmt.Printf("Beam index: %d\n", beam_index)


			if !CPU.Pause {
				if !CPU.DrawVSYNC {
					CPU.Interpreter()
				} else {
					fmt.Printf("\nWAIT FOR SYNC!!!!")
				}

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


	}

}



func drawVisibleModeLine() {
	if debug {
		fmt.Printf("Line: %d\t VISIBLE AREA\n", line)
	}

	// // 1 = Reflect first 20 sprites to the last 20
	// if D0_Reflect {
	// 	j := 0
	// 	for i := len(playfield) - 1; i > 19  ; i-- {
	// 		playfield[i] = playfield[j]
	// 		j++
	// 	}
	// // Duplicate last 20 sprites with first 20
	// }  else {
	// 	for i := 20 ; i < len(playfield) ; i++ {
	// 		playfield[i] = playfield[i-20]
	// 	}
	// }

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
				R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				fmt.Printf("\nBackground 0-39: %d, %d, %d\n\n\n\n\n\n\n\n", R, G, B)
			} else {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[COLUPF])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			}

			// // If it is rendering the playfield
			// if search == 1 {
			// 	// If it is rendering a scoreboard
			// 	if drawing_score {
			// 		// Check D1 status to use color of players in the score
			// 		if D1_Score {
			// 			// READ COLUP0 (Memory[0x06]) - Set the Player 0 Color (On Score)
			// 			R, G, B := Palettes.NTSC(CPU.Memory[COLUP0])
			// 			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			// 			// Set P1 Color
			// 			if i<20 {
			// 				// READ COLUP1 (Memory[0x07]) - Set the Player 1 Color (On Score)
			// 				R, G, B := Palettes.NTSC(CPU.Memory[COLUP1])
			// 				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			// 			}
			// 		}
			// 	}
			// }


			// R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
			// imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			//
			// imd.Push(pixel.V(	0				, float64(232-line)*height ))
			// imd.Push(pixel.V(	win.Bounds().W()	, float64(232-line)*height ))

			// Draw
			//fmt.Printf("\ni: %d\tIndex: %d\tNumber of repeated %d: %d\n", i, index, search,count)
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width								, float64(232-line)*height ))
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height ))
			// fmt.Printf("%f %f", (68 + (float64(index) *5)),	(68) + (float64(index) *5) +float64(count*5) )
			imd.Line(height)
			count = 1
			index = i+1
			search = playfield[i+1]
		}
	}

	// Process the last value [19]
	if search == 0 {
		// READ COLUBK (Memory[0x09]) - Set the Background Color
		R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
	} else {
		// READ COLUPF (Memory[0x08]) - Set the Playfield Color
		R, G, B := Palettes.NTSC(CPU.Memory[COLUPF])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
	}

	//fmt.Printf("\ni: 19\tIndex: %d\tNumber of repeated %d: %d\n",  index, search,count)
	imd.Push(pixel.V(  (float64(index) *pixelSize)*width								, float64(232-line)*height ))
	imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height ))
	//fmt.Printf("%f %f", (68 + (float64(index) *5)) ,	(68) + (float64(index) *5) +float64(count*5) )
	imd.Line(height)

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
	line ++
}







// func drawVisibleModeHW() {
// 	// Dont draw first 68 pixels (Horizontal Blank)
// 	if ( beam_index > 67 ) {
// 		if debug {
// 			fmt.Printf("Line: %d\t VISIBLE AREA\n", line)
// 		}
//
//
// 		// 3 pixels per cycle
// 		for i := 0 ; i < 3 ; i++ {
// 			//fmt.Printf("Line: %d, i: %d | beam_index: %d\n", line, i, beam_index)
//
// 			// Test if the beam is inside the X range and draw
// 			if ( beam_index + i < int(sizeX) ) {
// 				R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
// 				imd.Color = pixel.RGB(R, G, B)
// 				imd.Push(pixel.V( width*float64(beam_index + i)        , height*float64(line_max-line) ))
// 				imd.Push(pixel.V( width*float64(beam_index + i) + width, height*float64(line_max-line) ))
// 				imd.Line(height)
// 			}
// 		}
//
// 		// Each CPU Cycle, sends 3 pixels to be drawed
// 		if (beam_index + 3 < int(sizeX)) {
// 			beam_index += 3
// 		} else {
// 			// Start drawing a new line
// 			beam_index = 0
// 			line ++
// 		}
//
// 	// Skip drawing in Horizontal Blank
// 	} else {
// 		beam_index ++
// 	}
//
// 	imd.Draw(win)
// 	// Count draw operations number per second
// 	draws ++
// }
