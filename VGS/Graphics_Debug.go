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

	basePosition := screenHeight * (1 - sizeYused)	// Value reserved for debug on screen

	// -------------------------- Draw Debug Rectangle -------------------------- //
	// Background
	imd.Color = colornames.Bisque
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( screenWidth , basePosition ) )
	imd.Rectangle(0)

	// ------------------------------- Draw Boxes ------------------------------- //

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
	imd.Push(pixel.V ( 360 , 444 ) )
	imd.Push(pixel.V ( 410 , 426 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 361 , 443 ) )
	imd.Push(pixel.V ( 409 , 427 ) )
	imd.Rectangle(0)

	// SP
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 360 , 424 ) )
	imd.Push(pixel.V ( 410 , 406 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 361 , 423 ) )
	imd.Push(pixel.V ( 409 , 407 ) )
	imd.Rectangle(0)

	// A
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 360 , 404 ) )
	imd.Push(pixel.V ( 410 , 386 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 361 , 403 ) )
	imd.Push(pixel.V ( 409 , 387 ) )
	imd.Rectangle(0)

	// X
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 360 , 384 ) )
	imd.Push(pixel.V ( 410 , 366 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 361 , 383 ) )
	imd.Push(pixel.V ( 409 , 367 ) )
	imd.Rectangle(0)

	// Y
	imd.Color = colornames.Black
	imd.Push(pixel.V ( 360 , 364 ) )
	imd.Push(pixel.V ( 410 , 346 ) )
	imd.Rectangle(0)
	imd.Color = colornames.White
	imd.Push(pixel.V ( 361 , 363 ) )
	imd.Push(pixel.V ( 409 , 347 ) )
	imd.Rectangle(0)

	// ------------------------------ Debug Borders ----------------------------- //

	imd.Color = colornames.White
	// Up bar
	imd.Push(pixel.V ( 0 , basePosition  ) )
	imd.Push(pixel.V ( screenWidth , basePosition -2 ) )
	imd.Rectangle(0)
	// Down bar
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( screenWidth , 2 ) )
	imd.Rectangle(0)
	// Left bar
	imd.Push(pixel.V ( 0 , 0  ) )
	imd.Push(pixel.V ( 2 , basePosition ) )
	imd.Rectangle(0)
	// Right bar
	imd.Push(pixel.V ( screenWidth , 0  ) )
	imd.Push(pixel.V ( screenWidth -2 , basePosition ) )
	imd.Rectangle(0)

	imd.Draw(win)
}

func drawDebugInfo() {

	var (
		fontSize		float64 = 1
		txt			string
	)

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
	if Cycle == 0 {
		txt = fmt.Sprintf("%d  \n",Cycle)
	} else {
		txt = fmt.Sprintf("%d  \n",Cycle - 1)
	}
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Frame Cycle
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "F. Cycle:         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	if Cycle == 0 {
		txt = fmt.Sprintf("%d  \n",Cycle)
	} else {
		txt = fmt.Sprintf("%d  \n",Cycle - 1)
	}
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
	txt = fmt.Sprintf("%d  \n",PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Scan Cycle
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Scan Cycle:         ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Pixel Pos
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Pixel Pos:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Color Clock
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Color Clk:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",PC)
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
	txt = fmt.Sprintf("%d  \n",PC)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// SP
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "SP:          ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",SP)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// A
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "A:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",A)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// X
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "X:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",X)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	// Y
	cpuMessage.Color = colornames.Black
	fmt.Fprintf(cpuMessage, "Y:           ")
	// cpuMessage.Color = colornames.White
	txt = ""
	txt = fmt.Sprintf("%d  \n",Y)
	cpuMessage.Dot.X -= cpuMessage.BoundsOf(txt).W()
	fmt.Fprintf(cpuMessage, txt)

	cpuMessage.Draw(win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))


	// // Opcode
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "Opcode:")
	// cpuMessage.Color = colornames.White
	// text = fmt.Sprintf(" %04X  ",Opcode)
	// fmt.Fprintf(cpuMessage, text)
	// // PC
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "PC:               ")
	// cpuMessage.Color = colornames.White
	// text = fmt.Sprintf("%d(0x%04X)  ",PC, PC)
	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// fmt.Fprintf(cpuMessage, text)
	// // I
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "I:       ")
	// cpuMessage.Color = colornames.White
	// // text = fmt.Sprintf("%d  ",I)
	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// fmt.Fprintf(cpuMessage, text)
	// // DT
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "DT:      ")
	// cpuMessage.Color = colornames.White
	// // text = fmt.Sprintf("%d  ",DelayTimer)
	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// fmt.Fprintf(cpuMessage, text)
	// // ST
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "ST:      ")
	// cpuMessage.Color = colornames.White
	// // text = fmt.Sprintf("%d  ",SoundTimer)
	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// fmt.Fprintf(cpuMessage, text)
	// // SP
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "SP:     ")
	// cpuMessage.Color = colornames.White
	// text = fmt.Sprintf("%d",SP)
	// cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// fmt.Fprintf(cpuMessage, text)
	// // Stack
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "\nStack: ")
	// cpuMessage.Color = colornames.White
	// fmt.Fprintf(cpuMessage, "[   ")
	// // for i:=0 ; i  <len(Stack) ; i++ {
	// 	// text = fmt.Sprintf("%d",Stack[i])
	// // 	cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// // 	fmt.Fprintf(cpuMessage, text)
	// // 	if i < 15 {
	// // 		fmt.Fprintf(cpuMessage, "     ")
	// // 	}
	// // }
	// fmt.Fprintf(cpuMessage, "]")
	// // V
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "\nV:     ")
	// cpuMessage.Color = colornames.White
	// fmt.Fprintf(cpuMessage, "[   ")
	// // for i:=0 ; i  <len(V) ; i++ {
	// // 	text = fmt.Sprintf("%d",V[i])
	// // 	cpuMessage.Dot.X -= cpuMessage.BoundsOf(text).W()
	// // 	fmt.Fprintf(cpuMessage, text)
	// // 	if i < 15 {
	// // 		fmt.Fprintf(cpuMessage, "     ")
	// // 	}
	// // }
	// fmt.Fprintf(cpuMessage, "]")
	// // Keys
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage, "\nKeys: ")
	// cpuMessage.Color = colornames.White
	// // text = fmt.Sprintf(" %d  ",Key)
	// fmt.Fprintf(cpuMessage, text)
	// //Opcode Message
	// cpuMessage.Color = colornames.Black
	// fmt.Fprintf(cpuMessage,"\nMsg:   ")
	// cpuMessage.Color = colornames.White
	// // text = fmt.Sprintf("%s ",OpcMessage)
	// fmt.Fprintf(cpuMessage, text)

	// Draw Text
	// cpuMessage.Draw(Win, pixel.IM.Scaled(cpuMessage.Orig, fontSize))
}
