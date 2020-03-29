package Graphics

import (
	"fmt"
	"strconv"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"Atari2600/CPU"

)

var (
	// Window Configuration
	win				* pixelgl.Window
	imd				= imdraw.New(nil)
	cfg				= pixelgl.WindowConfig{}


	// Line draw control
	line				int = 0
	line_max			int = 261

	// Emulate CRT Electron beam
	beam_index		int = 0
	background_color	= colornames.Blue
	playfield_color	= colornames.Yellow
	line_color		= colornames.Red
	p0_color			= colornames.Green
	p1_color			= colornames.Purple
	GRP0				byte = 0			// Graphic Player 0 position
	GRP1				byte = 0			// Graphic Player 1 position
	// Flag used to put different colors in scores
	drawing_score		bool = false
	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]int			//Improve to binary
	pixelSize			float64 = 4.0

	// CTRLPLF (8 bits register)
	// D0 = Reflect, false = Repeat
	D0_Reflect			bool = false
	// D1 = Score == Color of the score will be the same as player
	D1_Score				bool = true
	// D2 = Priority == Player behind the playfield
	// D4-5 = Ball Size (1, 2, 4, 8)


	// Future Wait for Sync implementation
	//WSYNC			bool = false

	// FPS count
	frames			= 0
	draws			= 0

	// Draw 3 pixels per cycle or in optimized mode
	draw_mode_hw		bool = false

	// Debug mode
	debug			bool = false


	Fontset2 			= []byte{
		0x0E, 0x0E, 0x02, 0x02, 0x0E, 0x0E, 0x08, 0x08, 0x0E, 0x0E, // Number2
		0x7E, 0xFF, 0x99, 0xFF, 0xFF, 0xFF, 0xBD, 0xC3, 0xFF, 0x7E, // Smile Face
	}


)

const (
	sizeX			float64	= 228.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY			float64	= 262.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	screenWidth		= float64(sizeX*6)
	screenHeight		= float64(sizeY*3)
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


func drawLine() {
	if debug {
		fmt.Printf("Line: %d\t VERTICAL SYNC - Draw Line\n", line)
	}

	imd.Color = line_color

	imd.Push(pixel.V(	0				, float64(line_max-line)*height ))
	imd.Push(pixel.V(	win.Bounds().W()	, float64(line_max-line)*height ))
	imd.Line(height)
	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}


func drawVisibleModeHW() {
	// Dont draw first 68 pixels (Horizontal Blank)
	if ( beam_index > 67 ) {
		if debug {
			fmt.Printf("Line: %d\t VISIBLE AREA\n", line)
		}

		// 3 pixels per cycle
		for i := 0 ; i < 3 ; i++ {
			//fmt.Printf("Line: %d, i: %d | beam_index: %d\n", line, i, beam_index)

			// Test if the beam is inside the X range and draw
			if ( beam_index + i < int(sizeX) ) {
				imd.Push(pixel.V( width*float64(beam_index + i)        , height*float64(line_max-line) ))
				imd.Push(pixel.V( width*float64(beam_index + i) + width, height*float64(line_max-line) ))
				imd.Line(height)
			}
		}

		// Each CPU Cycle, sends 3 pixels to be drawed
		if (beam_index + 3 < int(sizeX)) {
			beam_index += 3
		} else {
			// Start drawing a new line
			beam_index = 0
			line ++
		}

	// Skip drawing in Horizontal Blank
	} else {
		beam_index ++
	}

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}


func drawVisibleModeLine() {
	if debug {
		fmt.Printf("Line: %d\t VISIBLE AREA\n", line)
	}

	// 1 = Reflect first 20 sprites to the last 20
	if D0_Reflect {
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
	index:=0
	count := 1

	for i := 0 ; i < len(playfield) -1; i++ {
		if playfield[i+1] == search {
			//fmt.Printf("\nI: %d\tRepeated Number\n",i)
			count++
		} else {
			// Set color (0: Background | 1: Playfield)
			if search == 0 {
				imd.Color = background_color
			} else {
				imd.Color = playfield_color
			}

			// If it is rendering the playfield
			if search == 1 {
				// If it is rendering a scoreboard
				if drawing_score {
					// Check D1 status to use color of players in the score
					if D1_Score {
						// Set P0 Color
						imd.Color = p0_color
						// Set P1 Color
						if i<20 {
							imd.Color = p1_color
						}
					}
				}
			}

			// Draw
			//fmt.Printf("\ni: %d\tIndex: %d\tNumber of repeated %d: %d\n", i, index, search,count)
			imd.Push(pixel.V(  (68*width) + (float64(index) *pixelSize)*width								, float64(line_max-line)*height ))
			imd.Push(pixel.V(  (68*width) + (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(line_max-line)*height ))
			//fmt.Printf("%f %f", (68 + (float64(index) *5)),	(68) + (float64(index) *5) +float64(count*5) )
			imd.Line(height)
			count = 1
			index = i+1
			search = playfield[i+1]
		}
	}

	// Process the last value [19]
	if search == 0 {
		imd.Color = background_color
	} else {
		imd.Color = playfield_color
	}

	//fmt.Printf("\ni: 19\tIndex: %d\tNumber of repeated %d: %d\n",  index, search,count)
	imd.Push(pixel.V(  (68*width) + (float64(index) *pixelSize)*width						, float64(line_max-line)*height ))
	imd.Push(pixel.V(  (68*width) + (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(line_max-line)*height ))
	//fmt.Printf("%f %f", (68 + (float64(index) *5)) ,	(68) + (float64(index) *5) +float64(count*5) )
	imd.Line(height)

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
	line ++
}




func drawGraphics() {

	imd	= imdraw.New(nil)

	// Set backgound color
	imd.Color = background_color

	// Skip draw ultil Visible Area
	if line < 40 {

		// 3 first lines = VERTICAL SYNC
		if (line < 2) {
			if debug {
				fmt.Printf("Line: %d\t VERTICAL SYNC\n", line)
			}
		} else if (line == 2) {

			drawLine()

		// Next 37 = VERTICAL BLANK
		} else if line < 39 {
			if debug {
				fmt.Printf("Line: %d\t VERTICAL BLANK\n", line)
			}
		} else if line == 39 {
			if debug {
				fmt.Printf("Line: %d\t VERTICAL BLANK - Draw Line\n", line)
			}

			drawLine()

			// Count draw operations number per second
			draws ++
		}

		line ++

	// DRAW VISIBLE LINES
	} else if line < 232 {

		// Draw 3 pixels each CPU Cycle
		// NTSC TV specification
		// NO PLAYFIELD DRAW IMPLEMENTED
		if draw_mode_hw {
			drawVisibleModeHW()
		// Draw line mode (Optimized)
		} else {
			drawVisibleModeLine()
		}


	// OVERSCAN - Last 30 lines
	} else if line == 232 {
		if debug {
			fmt.Printf("Line: %d\t OVERSCAN - Draw Line\n", line)
		}

		drawLine()

		line ++
	} else {
		if debug {
			fmt.Printf("Line: %d\t OVERSCAN\n", line)
		}
		line ++
	}

	// win.Update()
	// frames++
	// Every Cycle Control the clock!!!
	select {
	case <-CPU.ScreenRefresh.C:
	// When ticker run (60 times in a second, check de DelayTimer)

		win.Update()
		frames++
		default:
			// No timer to handle
	}

}






func Run() {
	// Initialize
	playfield		= [40]int{}

	// // Load some fonts to memory
	// for i := 0 ; i < 20 ; i++ {
	// 	CPU.Memory[i] = Fontset2[i]
	// }
	//
	// fmt.Printf("Memory: %02X \n", memory)

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)

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
			//fmt.Printf("CPU Cycle: %d\n", cycle)
			//fmt.Printf("Beam index: %d\n", beam_index)


			//drawPlayfield()
			//drawScore()


			// DRAW
			drawGraphics()

			CPU.Interpreter()

			// When finished drawing the screen, reset and start a new frame
			if line == line_max + 1 {
				//fmt.Printf("Line 0\t\tFinished the screen height, start a new frame.\n")
				line = 0
			}

			// Reset flag
			drawing_score = false

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

	}

}


func loadP1( value byte) {
	var (
		binary string = ""
	)

	binary = fmt.Sprintf("%.8b", value)
	//fmt.Printf("%d",playfield)
	for i := 0 ; i < 8 ; i++ {

		bit_binary, err := strconv.Atoi(fmt.Sprintf("%c", binary[i]))
		if err == nil {
			playfield[4+i] = bit_binary
			//fmt.Printf("%0d\n",bit_binary)
		}

	}
	//fmt.Printf("%08b\n",value)
	//fmt.Printf("%d\n",playfield)

}

func drawScore() {
	// Draw Playfield

	if line == 0 {
		playfield = [40]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	}


	if line >= 47 && line <= 56 {
		drawing_score = true

		loadP1(CPU.Memory[line-47])
	}

	if line >= 80 && line <= 90 {

		loadP1(CPU.Memory[line-70])
	}

	if line == 57 {
		playfield = [40]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

//		playfield = [40]int{0,0,0,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	}

}

func drawPlayfield() {
	// Draw Playfield

	if line == 0 {
		playfield = [40]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	}

	if line >= 47 && line <= 53 {
		playfield = [40]int{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	}

	if line > 53 && line <= 247 {
		playfield = [40]int{0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
		}

	if line >= 218 && line <= 224 {
		playfield = [40]int{0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	}

	if line > 224 {
		playfield = [40]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	}

}
