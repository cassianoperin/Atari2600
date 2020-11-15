package VGS

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel/imdraw"
)

// Change blocks position HERE
var (
	// Memory
	block_memory_X	float64 = 380
	block_memory_Y	float64 = 558
	// Title
	block_title_X	float64 = 20
	block_title_Y	float64 = 460
	// Frame and Frame Cycle
	block_1_X	float64 = 85
	block_1_Y	float64 = 444
	// Scan Line, Scan Cycle, Pixel Pos, Color Clock
	block_2_X	float64 = 250
	block_2_Y	float64 = 444
	// PC, SP, X, Y, A, P
	block_3_X	float64 = 358
	block_3_Y	float64 = 444
	// #PC, #SP, #X, #Y, #A
	block_4_X	float64 = 475
	block_4_Y	float64 = 444
	// Opcode
	block_opcode_X	float64 = 20
	block_opcode_Y	float64 = 290
)

func drawDebugScreen(imd *imdraw.IMDraw) {

	// Draw forms
	dbg_draw_background()
	dbg_draw_block_memory( block_memory_X , block_memory_Y)
	dbg_draw_block_1( block_1_X , block_1_Y)
	dbg_draw_block_2( block_2_X , block_2_Y)
	dbg_draw_block_3( block_3_X , block_3_Y)
	dbg_draw_block_4( block_4_X , block_4_Y)
	dbg_draw_opcode( block_opcode_X , block_opcode_Y)
	imd.Draw(win)
}

func drawDebugInfo() {
	// Draw texts
	dbg_draw_text_memory  ( block_memory_X , block_memory_Y )
	dbg_draw_text_title   ( block_title_X , block_title_Y )
	dbg_draw_text_block_1 ( block_1_X , block_1_Y )
	dbg_draw_text_block_2 ( block_2_X , block_2_Y )
	dbg_draw_text_block_3 ( block_3_X , block_3_Y )
	dbg_draw_text_block_4 ( block_4_X , block_4_Y )
	dbg_draw_text_opcode  ( block_opcode_X , block_opcode_Y)
}


func startDebug() {

	Debug = true

	win.Clear(colornames.Black)
	sizeYused = 0.3
	sizeXused = 0.3
	// Show messages
	if Debug {
		fmt.Printf("\t\tDEBUG mode Enabled\n")
	}
	// win.Clear(colornames.Black)
	TextMessageStr = "DEBUG mode Enabled"
	ShowMessage = true

	// Set Initial resolution
	activeSetting = &settings[3]

	if isFullScreen {
		win.SetMonitor(activeSetting.Monitor)
	} else {
		win.SetMonitor(nil)
	}
	win.SetBounds(pixel.R(0, 0, float64(activeSetting.Mode.Width), float64(activeSetting.Mode.Height)))

	// Update Width and Height values accordingly to new resolutions
	screenWidth	= win.Bounds().W()
	screenHeight	= win.Bounds().H()
	width		= screenWidth/sizeX * sizeXused		// Define the width of the pixel, considering the percentage of screen reserved for emulator
	height		= screenHeight/sizeY * sizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

	// Draw Debug Screen
	drawDebugScreen(imd)	// Background
	drawDebugInfo()			// Info

	win.Update()
}


func stopDebug() {

	Debug = false

	sizeYused = 1.0
	sizeXused = 1.0
	// Show messages
	fmt.Printf("\t\tDEBUG mode Disabled\n")
	TextMessageStr = "DEBUG mode Disabled"
	ShowMessage = true

	// Update Width and Height values accordingly to new resolutions
	screenWidth	= win.Bounds().W()
	screenHeight	= win.Bounds().H()
	width		= screenWidth/sizeX
	height		= screenHeight/sizeY * sizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator

	// Reset graphics
	win.Clear(colornames.Black)

	win.Update()
}

func updateDebug() {
	drawDebugScreen(imd)	// Background
	drawDebugInfo()			// Info
}


//------------------//

// Draw the background of debug screen
func dbg_draw_background() {
	basePositionX := screenWidth  * sizeXused	// Value reserved for debug on screen
	basePositionY := screenHeight * (1 - sizeYused)	// Value reserved for debug on screen

	// ------------------------ Draw Debug Rectangles ----------------------- //
	// Background
	imd.Color = colornames.Lightgray
	// imd.Push(pixel.V ( screenWidth , basePositionX  ) )
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( screenWidth , basePositionY ) )
	imd.Rectangle(0)
	// imd.Push(pixel.V ( screenWidth , basePositionY ) )
	// imd.Push(pixel.V ( basePositionX , screenHeight  ) )
	imd.Push(pixel.V ( basePositionX , basePositionY ) )
	imd.Push(pixel.V ( screenWidth , screenHeight  ) )
	imd.Rectangle(0)

	// Borders
	imd.Color = colornames.Gray
	// Up bar
	// imd.Push(pixel.V ( 0 , basePositionY  ) )
	// imd.Push(pixel.V ( screenWidth , basePositionY -2 ) )
	imd.Push(pixel.V ( 0 , screenHeight  ) )
	imd.Push(pixel.V ( screenWidth , screenHeight -2 ) )
	imd.Rectangle(0)
	// Down bar
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( screenWidth , 2 ) )
	imd.Rectangle(0)
	// Left bar
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( 2 , basePositionY ) )
	imd.Rectangle(0)
	// Right bar
	imd.Push(pixel.V ( screenWidth , 0  ) )
	imd.Push(pixel.V ( screenWidth -2 , basePositionY ) )
	imd.Rectangle(0)
}


func dbg_draw_block_memory(x, y float64) {
	var (
		grade_pos_X float64 = x + 25
		grade_pos_Y float64 = y + 140
	)

	imd.Color = colornames.Black
	imd.Push(pixel.V ( x		, y )       )
	imd.Push(pixel.V ( x + 444	, y + 157 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( x + 1 		, y + 1 )       )
	imd.Push(pixel.V ( x + 444 - 1	, y + 157 - 1 ) )
	imd.Rectangle(0)
	imd.Color = colornames.Gray

	// imd.Push(pixel.V(100, 200), pixel.V(100, 500))
	// Vertical Grade
	for i:= 0 ; i < 15 ; i++ {
		imd.Push(pixel.V ( grade_pos_X , y+157) )
		imd.Push(pixel.V ( grade_pos_X , y    ) )
		imd.Line(1)
		grade_pos_X += 28
	}
	// Horizontal Grade
	for i:= 0 ; i < 7 ; i++ {
		imd.Push(pixel.V ( x+443, grade_pos_Y ) )
		imd.Push(pixel.V ( x	, grade_pos_Y ) )
		imd.Line(1)
		grade_pos_Y -= 20
	}

}


func dbg_draw_text_memory(x, y float64) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	x -= 30
	y += 182

	// ------------------------------- Memory ------------------------------- //

	cpuMessage = text.New(pixel.V(x, y), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// Frame
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "RAM (0x0080 ~ 0x00FF)\n00xx  0   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F")
	// cpuMessage.Color = colornames.White
	txt = ""
	label := 8

	for i := 0 ; i < 128 ; i++ {
		if i % 16 == 0 {
			txt += fmt.Sprintf("\n%X   ", label)
			label++
		}
		txt += fmt.Sprintf(" %02X ", Memory[0x80+i])
	}


	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}


// Debug Text
func dbg_draw_text_title(x , y float64) {
	var fontSize	float64 = 1

	cpuMessage = text.New(pixel.V(x, y), atlas)	// X, Y
	cpuMessage.Clear()
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "DEBUG")
	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize * 1.3))
}

// Shape Block 1
func dbg_draw_block_1(x , y float64) {

	// Frames, Frame Cycle
	for i:= 0 ; i < 2 ; i++ {
		imd.Color = colornames.Black
		imd.Push(pixel.V ( x      , y      ) )
		imd.Push(pixel.V ( x + 50 , y - 18 ) )
		imd.Rectangle(0)
		imd.Color = colornames.White
		imd.Push(pixel.V ( x + 1  , y - 1  ) )
		imd.Push(pixel.V ( x + 49 , y - 17 ) )
		imd.Rectangle(0)

		// Decrement Y
		y -=20
	}
}

func dbg_draw_text_block_1(x , y float64) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	x-=65
	y-=14

	cpuMessage = text.New(pixel.V(x, y), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// Frame
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Frame:            ")
	// cpuMessage.Color = colornames.White
	txt = ""
	// if counter_F_Cycle == 0 {
		txt = fmt.Sprintf("%d  \n", counter_Frame)
	// } else {
	// 	txt = fmt.Sprintf("%d  \n",counter_Frame - 1)
	// }
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Frame Cycle
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "F. Cycle:         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	// if counter_F_Cycle == 0 {
		txt = fmt.Sprintf("%d  \n", counter_F_Cycle)
	// } else {
	// 	txt = fmt.Sprintf("%d  \n",counter_F_Cycle - 1)
	// }
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}


func dbg_draw_block_2(x , y float64) {

	// Scan Line, Scan Cycle, Pixel Pos, Color Clock
	for i := 0 ; i < 4 ; i++ {
		imd.Color = colornames.Black
		imd.Push(pixel.V ( x      , y      ) )
		imd.Push(pixel.V ( x + 50 , y - 18 ) )
		imd.Rectangle(0)
		imd.Color = colornames.White
		imd.Push(pixel.V ( x + 1  , y - 1  ) )
		imd.Push(pixel.V ( x + 49 , y - 17 ) )
		imd.Rectangle(0)

		// Decrement Y
		y -=20
	}
}

func dbg_draw_text_block_2( x , y float64 ) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	x-=80
	y-=14

	cpuMessage = text.New(pixel.V(x, y), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// Scan Line
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Scan Line:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n", line)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Scan Cycle
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Scan Cycle:         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n", beamIndex)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Pixel Pos
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Pixel Pos:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",  (int16(beamIndex) * 3) - 68)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Color Clock
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Color Clk:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n", beamIndex * 3)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}

func dbg_draw_block_3( x , y float64 ) {

	var grade_pos_X float64 = x + 21

	// Scan Line, Scan Cycle, Pixel Pos, Color Clock
	for i := 0 ; i < 5 ; i++ {
		imd.Color = colornames.Black
		imd.Push(pixel.V ( x      , y      ) )
		imd.Push(pixel.V ( x + 52 , y - 18 ) )
		imd.Rectangle(0)
		imd.Color = colornames.White
		imd.Push(pixel.V ( x + 1  , y - 1  ) )
		imd.Push(pixel.V ( x + 51 , y - 17 ) )
		imd.Rectangle(0)

		// Decrement Y
		y -=20
	}

	// P
	imd.Color = colornames.Black
	imd.Push(pixel.V ( x      , y    +1  ) )
	imd.Push(pixel.V ( x + 169 , y - 17 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( x + 1  , y   ) )
	imd.Push(pixel.V ( x + 168 , y - 16 ) )
	imd.Rectangle(0)

	// Grade
	imd.Color = colornames.Gray
	for i:= 0 ; i < 7 ; i++ {
		imd.Push(pixel.V ( grade_pos_X, y ) )
		imd.Push(pixel.V ( grade_pos_X, y - 17) )
		imd.Line(1)
		grade_pos_X+=21
	}

}

func dbg_draw_text_block_3(x, y float64) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	// // PC, SP, X, Y, A, P
	// block_3_X float64 = 358
	// block_3_Y float64 = 444

	x -= 28
	y -= 14

	cpuMessage = text.New(pixel.V( x, y ), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// PC
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "PC:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%X  \n", PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// SP
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "SP:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%X  \n", SP)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// A
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "A:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%X  \n", A)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// X
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "X:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%X  \n", X)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Y
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Y:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%X  \n", Y)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// P
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "P:                         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  %d  -  %d  %d  %d  %d  %d\n", P[7], P[6], P[4], P[3], P[2], P[1], P[0] )
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)
	fmt.Fprintf(cpuMessage, "     N  V  -  B  D  I  Z  C   ")


	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}

func dbg_draw_block_4( x, y float64) {

	// #PC, #SP, #X, #Y, #A
	for i := 0 ; i < 5 ; i++ {
		imd.Color = colornames.Black
		imd.Push(pixel.V ( x      , y      ) )
		imd.Push(pixel.V ( x + 52 , y - 18 ) )
		imd.Rectangle(0)
		imd.Color = colornames.White
		imd.Push(pixel.V ( x + 1  , y - 1  ) )
		imd.Push(pixel.V ( x + 51 , y - 17 ) )
		imd.Rectangle(0)

		// Decrement Y
		y -=20
	}
}


func dbg_draw_text_block_4( x, y float64 ) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	x -= 38
	y -= 14

	cpuMessage = text.New(pixel.V(x, y), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// PC #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "PC#:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// SP #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "SP#:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",SP)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// A #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "A #:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",A)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// X #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "X #:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",X)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Y #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Y #:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",Y)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}


func dbg_draw_text_opcode(x, y float64) {
	var (
		fontSize	float64 = 1
		txt			string
	)

	cpuMessage = text.New(pixel.V(x, y), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// Opcode
	cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "Opcode:                                         ")
	fmt.Fprintf(cpuMessage, "Opcode:\n")
	txt = ""

	// Initialize the slice
	if len(dbg_opc_messages) < 10 {
		if debug_opc_text != "" {
			dbg_opc_messages = append(dbg_opc_messages, debug_opc_text)
		}
	// Remove the first one and add the new last one
	} else {
		if debug_opc_text != "" {
			copy(dbg_opc_messages[0:], dbg_opc_messages[1:]) // Shift a[i+1:] left one index.
			dbg_opc_messages[9] = debug_opc_text
		}
	}

	// Print Opcode Slice
	for i := 0 ; i < len(dbg_opc_messages) ; i++ {
		txt += dbg_opc_messages[i] + "\n"
	}

	fmt.Fprintf(cpuMessage, txt)
	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))


	// Opcode Detailed Message
	cpuMessage = text.New(pixel.V(0, 40), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5
	cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "Opcode:                                         ")
	fmt.Fprintf(cpuMessage, dbg_show_message)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}

func dbg_draw_opcode(x, y float64) {
	var (
		grade_pos_X float64
		grade_pos_Y float64
	)

	x -= 8
	y -= 6

	grade_pos_X = x + 220
	grade_pos_Y = y

	imd.Color = colornames.Black
	imd.Push(pixel.V ( x	    	, y + 2)       )
	imd.Push(pixel.V ( x + 300	, y - 194 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( x + 1 		    , y + 2- 1 )       )
	imd.Push(pixel.V ( x + 300 - 1	, y - 194 + 1 ) )
	imd.Rectangle(0)

	// Current Opcode
	imd.Color = colornames.Lightyellow
	if len(dbg_opc_messages) >= 0 && len(dbg_opc_messages) <= 9 && counter_F_Cycle > 0 {
		imd.Push(pixel.V ( x + 1 		    , y - float64(1.0 * (19.3 * float64(len(dbg_opc_messages))) + 19.3) )  )
		imd.Push(pixel.V ( x + 300 - 1	, y - float64(19.6 * float64(len(dbg_opc_messages))) + 1.0) )
		imd.Rectangle(0)
	}
	if len(dbg_opc_messages) == 10 {
		imd.Push(pixel.V ( x + 1 		    , y - float64 (1.0 * (19.3 * float64(len(dbg_opc_messages)-1)) + 19.3) ) )
		imd.Push(pixel.V ( x + 300 - 1	, y - float64(19.6 * float64(len(dbg_opc_messages)-1)) + 1.0) )
		imd.Rectangle(0)
	}
	imd.Color = colornames.Gray

	// imd.Push(pixel.V(100, 200), pixel.V(100, 500))
	// Vertical Grade
	imd.Push(pixel.V ( grade_pos_X , y - 193   ) )
	imd.Push(pixel.V ( grade_pos_X , y + 1 ) )
	imd.Line(1)

	// Horizontal Grade
	imd.Color = colornames.Gray

	for i:= 0 ; i < 9 ; i++ {
		grade_pos_Y -= 19.4

		imd.Push(pixel.V ( x + 300, grade_pos_Y ) )
		imd.Push(pixel.V ( x + 1	, grade_pos_Y ) )
		imd.Line(1)

	}

}
