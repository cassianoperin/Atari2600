package Graphics

import (
	"fmt"
	// "os"
	"strconv"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"Atari2600/Palettes"
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
	line_color		= colornames.Red
	//0000-002C - TIA (write)
	COLUP0			byte = 0x06
	COLUP1			byte = 0x07
	COLUPF			byte	= 0x08
	COLUBK			byte	= 0x09
	// GRP0				byte = 0			// Graphic Player 0 position
	// GRP1				byte = 0			// Graphic Player 1 position
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


	// ntsc = [256][3]float64{
	// {0, 0, 0},		{0, 0, 0},		{64, 64, 64},		{64, 64, 64},		{108, 108, 108},	{108, 108, 108},	{144, 144, 144},	{144, 144, 144},	{176, 176, 176},		{176, 176, 176},		{200, 200, 200},		{200, 200, 200},		{220, 220, 220},		{220, 220, 220},		{236, 236, 236},		{236, 236, 236},
	// {68, 68, 0},		{68, 68, 0},		{100, 100, 16},	{100, 100, 16},	{132, 132, 36},	{132, 132, 36},	{160, 160, 52},	{160, 160, 52},	{184, 184, 64},		{184, 184, 64},		{208, 208, 80},		{208, 208, 80},		{232, 232, 92},		{232, 232, 92},		{252, 252, 104},		{252, 252, 104},
	// {112, 40, 0},		{112, 40, 0},		{132, 68, 20},		{132, 68, 20},		{152, 92, 40},		{152, 92, 40},		{172, 120, 60},	{172, 120, 60},	{188, 140, 76},		{188, 140, 76},		{204, 160, 92},		{204, 160, 92},		{220, 180, 104},		{220, 180, 104},		{236, 200, 120},		{236, 200, 120},
	// {132, 24, 0},		{132, 24, 0},		{152, 52, 24},		{152, 52, 24},		{172, 80, 48},		{172, 80, 48},		{192, 104, 72},	{192, 104, 72},	{208, 128, 92},		{208, 128, 92},		{224, 148, 112},		{224, 148, 112},		{236, 168, 128},		{236, 168, 128},		{252, 188, 148},		{252, 188, 148},
	// {136, 0, 0},		{136, 0, 0},		{156, 32, 32},		{156, 32, 32},		{176, 60, 60},		{176, 60, 60},		{192, 88, 88},		{192, 88, 88},		{208, 112, 112},		{208, 112, 112},		{224, 136, 136},		{224, 136, 136},		{236, 160, 160},		{236, 160, 160},		{252, 180, 180},		{252, 180, 180},
	// {120, 0, 92},		{120, 0, 92},		{140, 32, 116},	{140, 32, 116},	{160, 60, 136},	{160, 60, 136},	{176, 88, 156},	{176, 88, 156},	{192, 112, 176},		{192, 112, 176},		{208, 132, 192},		{208, 132, 192},		{220, 156, 208},		{220, 156, 208},		{236, 176, 224},		{236, 176, 224},
	// {72, 0, 120},		{72, 0, 120},		{96, 32, 144},		{96, 32, 144},		{120, 60, 164},	{120, 60, 164},	{140, 88, 184},	{140, 88, 184},	{160, 112, 204},		{160, 112, 204},		{180, 132, 220},		{180, 132, 220},		{196, 156, 236},		{196, 156, 236},		{212, 176, 252},		{212, 176, 252},
	// {20, 0, 132},		{20, 0, 132},		{48, 32, 152},		{48, 32, 152},		{76, 60, 172},		{76, 60, 172},		{104, 88, 192},	{104, 88, 192},	{124, 112, 208},		{124, 112, 208},		{148, 136, 224},		{148, 136, 224},		{168, 160, 236},		{168, 160, 236},		{188, 180, 252},		{188, 180, 252},
	// {0, 0, 136},		{0, 0, 136},		{28, 32, 156},		{28, 32, 156},		{56, 64, 176},		{56, 64, 176},		{80, 92, 192},		{80, 92, 192},		{104, 116, 208},		{104, 116, 208},		{124, 140, 224},		{124, 140, 224},		{144, 164, 236},		{144, 164, 236},		{164, 184, 252},		{164, 184, 252},
	// {0, 24, 124},		{0, 24, 124},		{28, 56, 144},		{28, 56, 144},		{56, 84, 168},		{56, 84, 168},		{80, 112, 188},	{80, 112, 188},	{104, 136, 204},		{104, 136, 204},		{124, 156, 220},		{124, 156, 220},		{144, 180, 236},		{144, 180, 236},		{164, 200, 252},		{164, 200, 252},
	// {0, 44, 92},		{0, 44, 92},		{28, 76, 120},		{28, 76, 120},		{56, 104, 144},	{56, 104, 144},	{80, 132, 172},	{80, 132, 172},	{104, 156, 192},		{104, 156, 192},		{124, 180, 212},		{124, 180, 212},		{144, 204, 232},		{144, 204, 232},		{164, 224, 252},		{164, 224, 252},
	// {0, 60, 44},		{0, 60, 44},		{28, 92, 72},		{28, 92, 72},		{56, 124, 100},	{56, 124, 100},	{80, 156, 128},	{80, 156, 128},	{104, 180, 148},		{104, 180, 148},		{124, 208, 172},		{124, 208, 172},		{144, 228, 192},		{144, 228, 192},		{164, 252, 212},		{164, 252, 212},
	// {0, 60, 0},		{0, 60, 0},		{32, 92, 32},		{32, 92, 32},		{64, 124, 64},		{64, 124, 64},		{92, 156, 92},		{92, 156, 92},		{116, 180, 116},		{116, 180, 116},		{140, 208, 140},		{140, 208, 140},		{164, 228, 164},		{164, 228, 164},		{184, 252, 184},		{184, 252, 184},
	// {20, 56, 0},		{20, 56, 0},		{52, 92, 28},		{52, 92, 28},		{80, 124, 56},		{80, 124, 56},		{108, 152, 80},	{108, 152, 80},	{132, 180, 104},		{132, 180, 104},		{156, 204, 124},		{156, 204, 124},		{180, 228, 144},		{180, 228, 144},		{200, 252, 164},		{200, 252, 164},
	// {44, 48, 0},		{44, 48, 0},		{76, 80, 28},		{76, 80, 28},		{104, 112, 52},	{104, 112, 52},	{132, 140, 76},	{132, 140, 76},	{156, 168, 100},		{156, 168, 100},		{180, 192, 120},		{180, 192, 120},		{204, 212, 136},		{204, 212, 136},		{224, 236, 156},		{224, 236, 156},
	// {68, 40, 0},		{68, 40, 0},		{100, 72, 24},		{100, 72, 24},		{132, 104, 48},	{132, 104, 48},	{160, 132, 68},	{160, 132, 68},	{184, 156, 88},		{184, 156, 88},		{208, 180, 108},		{208, 180, 108},		{232, 204, 124},		{232, 204, 124},		{252, 224, 140},		{252, 224, 140},
	// }
	//
	// ntsc = [256][3]float64{
	// {0, 0, 0},	{68, 68, 0},	{112, 40, 0},	{132, 24, 0},	{136, 0, 0},	{120, 0, 92},	{72, 0, 120},	{20, 0, 132},	{0, 0, 136},	{0, 24, 124},	{0, 44, 92},	{0, 60, 44},	{0, 60, 0},	{20, 56, 0},	{44, 48, 0},	{68, 40, 0},
	// {0, 0, 0},	{68, 68, 0},	{112, 40, 0},	{132, 24, 0},	{136, 0, 0},	{120, 0, 92},	{72, 0, 120},	{20, 0, 132},	{0, 0, 136},	{0, 24, 124},	{0, 44, 92},	{0, 60, 44},	{0, 60, 0},	{20, 56, 0},	{44, 48, 0},	{68, 40, 0},
	//
	// {64, 64, 64},	{100, 100, 16},	{132, 68, 20},	{152, 52, 24},	{156, 32, 32},	{140, 32, 116},	{96, 32, 144},	{48, 32, 152},	{28, 32, 156},	{28, 56, 144},	{28, 76, 120},	{28, 92, 72},	{32, 92, 32},	{52, 92, 28},	{76, 80, 28},	{100, 72, 24},
	// {64, 64, 64},	{100, 100, 16},	{132, 68, 20},	{152, 52, 24},	{156, 32, 32},	{140, 32, 116},	{96, 32, 144},	{48, 32, 152},	{28, 32, 156},	{28, 56, 144},	{28, 76, 120},	{28, 92, 72},	{32, 92, 32},	{52, 92, 28},	{76, 80, 28},	{100, 72, 24},
	//
	// {108, 108, 108},	{132, 132, 36},	{152, 92, 40},	{172, 80, 48},	{176, 60, 60},	{160, 60, 136},	{120, 60, 164},	{76, 60, 172},	{56, 64, 176},	{56, 84, 168},	{56, 104, 144},	{56, 124, 100},	{64, 124, 64},	{80, 124, 56},	{104, 112, 52},	{132, 104, 48},
	// {108, 108, 108},	{132, 132, 36},	{152, 92, 40},	{172, 80, 48},	{176, 60, 60},	{160, 60, 136},	{120, 60, 164},	{76, 60, 172},	{56, 64, 176},	{56, 84, 168},	{56, 104, 144},	{56, 124, 100},	{64, 124, 64},	{80, 124, 56},	{104, 112, 52},	{132, 104, 48},
	//
	// {144, 144, 144},	{160, 160, 52},	{172, 120, 60},	{192, 104, 72},	{192, 88, 88},	{176, 88, 156},	{140, 88, 184},	{104, 88, 192},	{80, 92, 192},	{80, 112, 188},	{80, 132, 172},	{80, 156, 128},	{92, 156, 92},	{108, 152, 80},	{132, 140, 76},	{160, 132, 68},
	// {144, 144, 144},	{160, 160, 52},	{172, 120, 60},	{192, 104, 72},	{192, 88, 88},	{176, 88, 156},	{140, 88, 184},	{104, 88, 192},	{80, 92, 192},	{80, 112, 188},	{80, 132, 172},	{80, 156, 128},	{92, 156, 92},	{108, 152, 80},	{132, 140, 76},	{160, 132, 68},
	//
	// {176, 176, 176},	{184, 184, 64},	{188, 140, 76},	{208, 128, 92},	{208, 112, 112},	{192, 112, 176},	{160, 112, 204},	{124, 112, 208},	{104, 116, 208},	{104, 136, 204},	{104, 156, 192},	{104, 180, 148},	{116, 180, 116},	{132, 180, 104},	{156, 168, 100},	{184, 156, 88},
	// {176, 176, 176},	{184, 184, 64},	{188, 140, 76},	{208, 128, 92},	{208, 112, 112},	{192, 112, 176},	{160, 112, 204},	{124, 112, 208},	{104, 116, 208},	{104, 136, 204},	{104, 156, 192},	{104, 180, 148},	{116, 180, 116},	{132, 180, 104},	{156, 168, 100},	{184, 156, 88},
	//
	// {200, 200, 200},	{208, 208, 80},	{204, 160, 92},	{224, 148, 112},	{224, 136, 136},	{208, 132, 192},	{180, 132, 220},	{148, 136, 224},	{124, 140, 224},	{124, 156, 220},	{124, 180, 212},	{124, 208, 172},	{140, 208, 140},	{156, 204, 124},	{180, 192, 120},	{208, 180, 108},
	// {200, 200, 200},	{208, 208, 80},	{204, 160, 92},	{224, 148, 112},	{224, 136, 136},	{208, 132, 192},	{180, 132, 220},	{148, 136, 224},	{124, 140, 224},	{124, 156, 220},	{124, 180, 212},	{124, 208, 172},	{140, 208, 140},	{156, 204, 124},	{180, 192, 120},	{208, 180, 108},
	//
	// {220, 220, 220},	{232, 232, 92},	{220, 180, 104},	{236, 168, 128},	{236, 160, 160},	{220, 156, 208},	{196, 156, 236},	{168, 160, 236},	{144, 164, 236},	{144, 180, 236},	{144, 204, 232},	{144, 228, 192},	{164, 228, 164},	{180, 228, 144},	{204, 212, 136},	{232, 204, 124},
	// {220, 220, 220},	{232, 232, 92},	{220, 180, 104},	{236, 168, 128},	{236, 160, 160},	{220, 156, 208},	{196, 156, 236},	{168, 160, 236},	{144, 164, 236},	{144, 180, 236},	{144, 204, 232},	{144, 228, 192},	{164, 228, 164},	{180, 228, 144},	{204, 212, 136},	{232, 204, 124},
	//
	// {236, 236, 236},	{252, 252, 104},	{236, 200, 120},	{252, 188, 148},	{252, 180, 180},	{236, 176, 224},	{212, 176, 252},	{188, 180, 252},	{164, 184, 252},	{164, 200, 252},	{164, 224, 252},	{164, 252, 212},	{184, 252, 184},	{200, 252, 164},	{224, 236, 156},	{252, 224, 140},
	// {236, 236, 236},	{252, 252, 104},	{236, 200, 120},	{252, 188, 148},	{252, 180, 180},	{236, 176, 224},	{212, 176, 252},	{188, 180, 252},	{164, 184, 252},	{164, 200, 252},	{164, 224, 252},	{164, 252, 212},	{184, 252, 184},	{200, 252, 164},	{224, 236, 156},	{252, 224, 140},
	//
	// }

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
				// READ COLUBK (Memory[0x09]) - Set the Background Color
				R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
				imd.Color = pixel.RGB(R, G, B)
			} else {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[COLUPF])
				imd.Color = pixel.RGB(R, G, B)
			}

			// If it is rendering the playfield
			if search == 1 {
				// If it is rendering a scoreboard
				if drawing_score {
					// Check D1 status to use color of players in the score
					if D1_Score {
						// READ COLUP0 (Memory[0x06]) - Set the Player 0 Color (On Score)
						R, G, B := Palettes.NTSC(CPU.Memory[COLUP0])
						imd.Color = pixel.RGB(R, G, B)
						// Set P1 Color
						if i<20 {
							// READ COLUP1 (Memory[0x07]) - Set the Player 1 Color (On Score)
							R, G, B := Palettes.NTSC(CPU.Memory[COLUP1])
							imd.Color = pixel.RGB(R, G, B)
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
		// READ COLUBK (Memory[0x09]) - Set the Background Color
		R, G, B := Palettes.NTSC(CPU.Memory[COLUBK])
		imd.Color = pixel.RGB(R, G, B)
	} else {
		// READ COLUPF (Memory[0x08]) - Set the Playfield Color
		R, G, B := Palettes.NTSC(CPU.Memory[COLUPF])
		imd.Color = pixel.RGB(R, G, B)
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
	imd.Color = pixel.RGB(0, 0, 0)

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


			drawPlayfield()
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
