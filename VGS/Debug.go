package VGS

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel/imdraw"
)

// Print Graphics on Console
func drawDebugScreen(imd *imdraw.IMDraw) {

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

	// ------------------------------- Memory ------------------------------- //

	// cpuMessage = text.New(pixel.V(350, 740), atlas)


	imd.Color = colornames.Black
	imd.Push(pixel.V ( 380 , 715 ) )
	imd.Push(pixel.V ( 824 , 558 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 381 , 714 ) )
	imd.Push(pixel.V ( 823 , 559 ) )
	imd.Rectangle(0)
	imd.Color = colornames.Gray

	// imd.Push(pixel.V(100, 200), pixel.V(100, 500))
	// Vertical Grade
	var posX float64 = 405
	for i:= 0 ; i < 15 ; i++ {
		imd.Push(pixel.V ( posX, 715) )
		imd.Push(pixel.V ( posX, 558) )
		imd.Line(1)
		posX+=28
	}
	// Horizontal Grade
	var posY float64 = 698
	for i:= 0 ; i < 7 ; i++ {
		imd.Push(pixel.V ( 380, posY) )
		imd.Push(pixel.V ( 823, posY) )
		imd.Line(1)
		posY-=20
	}


	// ----------------------------- Draw Boxes ----------------------------- //

	// Frames
	imd.Color = colornames.Black
	imd.Push(pixel.V (  85 , 444  ) )
	imd.Push(pixel.V ( 135 , 426 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V (  86 , 443  ) )
	imd.Push(pixel.V ( 134 , 427 ) )
	imd.Rectangle(0)

	// Frame Cycle
	imd.Color = colornames.Black
	imd.Push(pixel.V (  85 , 424 ) )
	imd.Push(pixel.V ( 135 , 406 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V (  86 , 423 ) )
	imd.Push(pixel.V ( 134 , 407 ) )
	imd.Rectangle(0)

	// --------------------------- //

	// Scan Line
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 250 , 444 ) )
	imd.Push(pixel.V ( 300 , 426 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 251 , 443 ) )
	imd.Push(pixel.V ( 299 , 427 ) )
	imd.Rectangle(0)

	// Scan Cycle
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 250 , 424 ) )
	imd.Push(pixel.V ( 300 , 406 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 251 , 423 ) )
	imd.Push(pixel.V ( 299 , 407 ) )
	imd.Rectangle(0)

	// Pixel Pos
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 250 , 384 ) )
	imd.Push(pixel.V ( 300 , 366 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 251 , 383 ) )
	imd.Push(pixel.V ( 299 , 367 ) )
	imd.Rectangle(0)

	// Color Clock
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 250 , 404 ) )
	imd.Push(pixel.V ( 300 , 386 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 251 , 403 ) )
	imd.Push(pixel.V ( 299 , 387 ) )
	imd.Rectangle(0)

	// --------------------------- //

	// PC
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 444 ) )
	imd.Push(pixel.V ( 410 , 426 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 443 ) )
	imd.Push(pixel.V ( 409 , 427 ) )
	imd.Rectangle(0)

	// SP
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 424 ) )
	imd.Push(pixel.V ( 410 , 406 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 423 ) )
	imd.Push(pixel.V ( 409 , 407 ) )
	imd.Rectangle(0)

	// A
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 404 ) )
	imd.Push(pixel.V ( 410 , 386 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 403 ) )
	imd.Push(pixel.V ( 409 , 387 ) )
	imd.Rectangle(0)

	// X
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 384 ) )
	imd.Push(pixel.V ( 410 , 366 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 383 ) )
	imd.Push(pixel.V ( 409 , 367 ) )
	imd.Rectangle(0)

	// Y
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 364 ) )
	imd.Push(pixel.V ( 410 , 346 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 363 ) )
	imd.Push(pixel.V ( 409 , 347 ) )
	imd.Rectangle(0)

	// P
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 358 , 345 ) )
	imd.Push(pixel.V ( 525 , 326 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 359 , 344 ) )
	imd.Push(pixel.V ( 524 , 327 ) )
	imd.Rectangle(0)
	imd.Color = colornames.Gray
	// Grade
	// imd.Push(pixel.V(100, 200), pixel.V(100, 500))
	posX = 379
	for i:= 0 ; i < 7 ; i++ {
		imd.Push(pixel.V ( posX, 345) )
		imd.Push(pixel.V ( posX, 326) )
		imd.Line(1)
		posX+=21
	}

	// --------------------------- //

	// PC #
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 475 , 444 ) )
	imd.Push(pixel.V ( 525 , 426 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 476 , 443 ) )
	imd.Push(pixel.V ( 524 , 427 ) )
	imd.Rectangle(0)

	// SP #
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 475 , 424 ) )
	imd.Push(pixel.V ( 525 , 406 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 476 , 423 ) )
	imd.Push(pixel.V ( 524 , 407 ) )
	imd.Rectangle(0)

	// A #
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 475 , 404 ) )
	imd.Push(pixel.V ( 525 , 386 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 476 , 403 ) )
	imd.Push(pixel.V ( 524 , 387 ) )
	imd.Rectangle(0)

	// X #
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 475 , 384 ) )
	imd.Push(pixel.V ( 525 , 366 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 476 , 383 ) )
	imd.Push(pixel.V ( 524 , 367 ) )
	imd.Rectangle(0)

	// Y #
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 475 , 364 ) )
	imd.Push(pixel.V ( 525 , 346 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 476 , 363 ) )
	imd.Push(pixel.V ( 524 , 347 ) )
	imd.Rectangle(0)


	imd.Draw(win)
}

func drawDebugInfo() {

	var (
		fontSize		float64 = 1
		txt			string
	)

	// ------------------------------- Memory ------------------------------- //

	cpuMessage = text.New(pixel.V(350, 740), atlas)
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


	// -------------------------- Draw Text -------------------------- //

	// Debug Text
	cpuMessage = text.New(pixel.V(20, 460), atlas)	// X, Y
	cpuMessage.Clear()
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "DEBUG")
	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize * 1.3))



	// ----------------------- 1st Text Column ----------------------- //

	cpuMessage = text.New(pixel.V(20, 430), atlas)
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

	// ----------------------- 2nd Text Column ----------------------- //

	cpuMessage = text.New(pixel.V(170, 430), atlas)
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

	// ----------------------- 3rd Text Column ----------------------- //

	cpuMessage = text.New(pixel.V(330, 430), atlas)
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

	// ----------------------- 4th Text Column ----------------------- //

	cpuMessage = text.New(pixel.V(437, 430), atlas)
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

	// ----------------------- Opcode Info ----------------------- //
	cpuMessage = text.New(pixel.V(200, 100), atlas)
	cpuMessage.Clear()
	cpuMessage.LineHeight = atlas.LineHeight() * 1.5

	// PC #
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Opcode:                                         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	// txt = debug_opc_text
	//
	// dbg_opc_PC			uint16
	// dbg_opc_mnm			string
	// dbg_opc_bytes		uint16
	// dbg_opc_opcode		byte
	// dbg_opc_payload1	byte
	// dbg_opc_payload2	byte

	if dbg_opc_bytes == 1 {
		txt += fmt.Sprintf("%04x     %s     %d     %x     %02x",dbg_opc_PC, dbg_opc_mnm, dbg_opc_bytes, dbg_opc_opcode, dbg_opc_payload1)
	}
	// } else if dbg_opc_bytes == 2 {
	// 	txt += fmt.Sprintf("%s %X %X",mnm, opcode, payload_B1 )
	// }  else if dbg_opc_bytes == 3 {
	// 	txt += fmt.Sprintf("%s %X %X %X",mnm, opcode, payload_B1, payload_B2 )
	// }

	// txt += fmt.Sprintf(" %X %X", opcode, Memory[PC+1])
	// txt = fmt.Sprintf("%02X  \n",opcode)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
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
