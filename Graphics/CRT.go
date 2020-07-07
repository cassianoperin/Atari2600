package Graphics

import (
	"os"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"Atari2600/Palettes"
	"Atari2600/CPU"
	"image/color"

)

var (

	// Line draw control
	line				int = 1
	line_max			int = 262

	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]byte			//Improve to binary
	pixelSize			float64 = 4.0		// 80 lines (half screen) / 20 PF0, PF1 and PF2 bits

	// FPS count
	frames			= 0
	draws			= 0

	// Workaround to avoid  WSYNC before VSYNC
	VSYNC_passed		bool = false

	// TIA
	old_BeamIndex	byte = 0	// Used to draw the beam updates every cycle on the CRT
	visibleArea		bool		// Not used yet, but will be used to just draw in visible area
)



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

func PF_Reflect_Duplicate() {
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
}


func invert_bits(value byte) byte {
	// fmt.Printf("\n\n%08b", value)
	value = (value & 0xF0) >> 4 | (value & 0x0F) << 4;
	value = (value & 0xCC) >> 2 | (value & 0x33) << 2;
	value = (value & 0xAA) >> 1 | (value & 0x55) << 1;
	// fmt.Printf("\n\n%08b", value)

	return value;
}


func drawPlayer0() {
	var (
		bit		byte = 0
		inverted	byte = 0
	)

	if CPU.DrawP0 {

		// If a program doesnt use RESP0, initialize
		if CPU.XPositionP0 == 0 {
			CPU.XPositionP0 = 23
		}

		if debug {
			fmt.Printf("Line: %d\tGRP0: %08b\tXPositionP0: %d\tXFinePositionP0: %d\n", line, CPU.Memory[CPU.GRP0], CPU.XPositionP0, CPU.XFinePositionP0)
		}

		// For each bit in GRPn, draw if == 1
		for i:=0 ; i <=7 ; i++ {

			// handle the order of the bits (normal or inverted)
			if CPU.Memory[CPU.REFP0] == 0x00 {
				bit = CPU.Memory[CPU.GRP0] >> (7-byte(i)) & 0x01
			} else {
				// If Reflect Player Enabled (REFP0), invert the order of GRPn
				inverted = invert_bits(CPU.Memory[CPU.GRP0])
				bit = inverted >> (7-byte(i)) & 0x01
			}

			if bit == 1 {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP0])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				if CPU.Memory[CPU.NUSIZ0] == 0x00 {
					imd.Push(pixel.V( (float64( ((CPU.XPositionP0)*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( ((CPU.XPositionP0)*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x01 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) +float64(CPU.XFinePositionP0) + float64(16) )*width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) +float64(CPU.XFinePositionP0) + float64(16) )*width + width		, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x02 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x03 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(16) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(16) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x04 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(64) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(64) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x05 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i*2)) + float64(CPU.XFinePositionP0) ) * width					, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i*2)) + float64(CPU.XFinePositionP0) ) * width + (width*2)			, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x06 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(64) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP0*3) - 68 + byte(i)) + float64(CPU.XFinePositionP0) + float64(64) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ0] == 0x07 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*4)) + float64(CPU.XFinePositionP1) ) * width					, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*4)) + float64(CPU.XFinePositionP1) ) * width + (width*4)			, float64(232-line) * height + height))
					imd.Rectangle(0)
				}

				// imd.Draw(win)
				// // Count draw operations number per second
				// draws ++
			}
		}

		imd.Draw(win)
		// Count draw operations number per second
		draws ++

		CPU.DrawP0 = false
	}
}


func drawPlayer1() {
	var (
		bit		byte = 0
		inverted	byte = 0
	)

	if CPU.DrawP1 {

		// If a program doesnt use RESP0, initialize
		if CPU.XPositionP1 == 0 {
			CPU.XPositionP1 = 30
		}

		if debug {
			fmt.Printf("Line: %d\tGRP1: %08b\tXPositionP1: %d\tHMP1: %d\n", line, CPU.Memory[CPU.GRP1], CPU.XPositionP1, CPU.Memory[CPU.HMP1])
		}

		// For each bit in GRPn, draw if == 1
		for i:=0 ; i <=7 ; i++{

			// handle the order of the bits (normal or inverted)
			if CPU.Memory[CPU.REFP1] == 0x00 {
				bit = CPU.Memory[CPU.GRP1] >> (7-byte(i)) & 0x01
			} else {
				// If Reflect Player Enabled (REFP1), invert the order of GRPn
				inverted = invert_bits(CPU.Memory[CPU.GRP1])
				bit = inverted >> (7-byte(i)) & 0x01
			}

			if bit == 1 {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP1])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				if CPU.Memory[CPU.NUSIZ1] == 0x00 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x01 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(16) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(16) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x02 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x03 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) + float64(16) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) + float64(16) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) +float64(CPU.XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x04 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(64) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(64) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x05 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*2)) + float64(CPU.XFinePositionP1) ) * width					, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*2)) + float64(CPU.XFinePositionP1) ) * width + (width*2)			, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x06 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width						, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(64) ) * width			, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i)) + float64(CPU.XFinePositionP1) + float64(64) ) * width + width	, float64(232-line) * height + height))
					imd.Rectangle(0)
				} else if CPU.Memory[CPU.NUSIZ1] == 0x07 {
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*4)) + float64(CPU.XFinePositionP1) ) * width					, float64(232-line) * height ))
					imd.Push(pixel.V( (float64( (CPU.XPositionP1*3) - 68 + byte(i*4)) + float64(CPU.XFinePositionP1) ) * width + (width*4)			, float64(232-line) * height + height))
					imd.Rectangle(0)
				}
				// imd.Draw(win)
				// // Count draw operations number per second
				// draws ++
			}
		}

		imd.Draw(win)
		// Count draw operations number per second
		draws ++

		CPU.DrawP1 = false
	}
}



func CRT(action int8) {

	// Just draw in visible Area
	// if visibleArea {

		drawBackground()

		// if line ==40 {
		// 	CPU.Pause = true
		//
		// }
	// }


	switch action {
		// --------------------------------------- WSYNC --------------------------------------- //
		// Halt CPU until next scanline starts
		// Skip to the next scanline
		case int8(CPU.WSYNC): //0x02
			if debug {
				fmt.Printf("\tCRT - WSYNC SET\n")
			}

			// Test if in Vertical Blank (do not draw anything)
			if CPU.Memory[CPU.VBLANK] == 2 {
				// os.Exit(2)


				// During Vertical Blank, if vsync is set
				if  CPU.Memory[CPU.VSYNC] == 2  {

					VSYNC_passed = true	// Used to control WSYNCS before VSYNC

					// When VSYNC is set, CPU inform CRT to start a new frame
					// 3 lines VSYNC

					// ENABLE VSYNC
					if CPU.Memory[CPU.VSYNC] == 0x02 {

						if CPU.Memory[CPU.VBLANK] == 2 {
							if debug {
								fmt.Printf("\tLine: %d\tCRT - VSYNC\n\n", line)
							}
						} else {
							if debug {
								fmt.Printf("\tLine: %d\tCRT - VSYNC without VBLANK - Not correct!!!\n\n", line)
							}
						}

					// DISABLE VSYNC
					} else if CPU.Memory[CPU.VSYNC] == 0x00 {
						if debug {
							fmt.Printf("\tCRT - VSYNC DISABLED\n")
						}

					} else {
						fmt.Printf("\tCRT - VSYNC VALUE NOT 0 or 2! Exiting!\n")
						os.Exit(2)
					}

				// 37 lines VBLANK
				} else if CPU.Memory[CPU.VBLANK] == 2 {
					if debug {
						fmt.Printf("\tLine: %d\tVBLANK\t\t(vblank: %02X\tvsync: %02X)\n\n", line,CPU.Memory[CPU.VBLANK], CPU.Memory[CPU.VSYNC])
					}
					visibleArea = false // Inform that finished visible lines

				}

			// VBLANK turned OFF, start drawing the 192 lines of visible Area
			} else {
				visibleArea = true // Inform that reached visible lines

				// Finish drawing line (X=228) 76x3
				CPU.Beam_index = 76
				if debug {
					fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, CPU.Beam_index)
				}
				drawBackground()



				if debug {
					fmt.Printf("\tLine: %d\tVisible Area: %d\n\n", line, line-40)
				}

				// TODO UNIFY in ONE
				readPF0()
				readPF1()
				readPF2()
				PF_Reflect_Duplicate()
				// fmt.Println(playfield)

				// DRAW PLAYFIELD ENTIRE LINE

				for i := 0 ; i < len(playfield) ; i++ {

					if playfield[i] == 1 {

						// READ COLUPF (Memory[0x08]) - Set the Playfield Color
						R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUPF])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

						// Draw
						imd.Push(pixel.V( float64( i*4 )     * width	, float64(232-line) * height ))
						imd.Push(pixel.V( float64( i*4 + 4 ) * width 	, float64(232-line) * height + height))
						imd.Rectangle(0)

						imd.Draw(win)
						draws ++

					}
				}



				// drawVisibleModeLine()

				// // DRAW PLAYER 0
				// if CPU.DrawP0 {
				// 	drawPlayer0()
				//
				// 	CPU.DrawP0 = false
				// }
				//
				// // DRAW PLAYER 1
				// if CPU.DrawP1 {
				// 	drawPlayer1()
				//
				// 	CPU.DrawP1 = false
				// }

			}

			// Reset the beam index
			CPU.Beam_index = 0
			old_BeamIndex = 0
			line ++


		// --------------------------------------- VBLANK --------------------------------------- //
		case int8(CPU.VBLANK): //0x01

			// Enable VBLANK
			if CPU.Memory[CPU.VBLANK] == 0x02 {
				if debug {
					fmt.Printf("\n\tVBLANK Enabled")
				}
			} else if CPU.Memory[CPU.VBLANK] == 0x00 {
				if debug {
					fmt.Printf("\n\tVBLANK Disabled")
				}
			} else {
				if debug {
					fmt.Printf("\n\tVBLANK VALUE !=0 !=2 exiting\t%d", CPU.Memory[CPU.VBLANK])
				}
				// os.Exit(0)
			}

		// --------------------------------------- VSYNC --------------------------------------- //
		case int8(CPU.VSYNC): //0x00

			// Enable VSYNC
			if CPU.Memory[CPU.VSYNC] == 0x02 {
				if debug {
					fmt.Printf("\n\tVSYNC Enabled\n")
				}
			} else if CPU.Memory[CPU.VSYNC] == 0x00 {
				if debug {
					fmt.Printf("\n\tVSYNC Disabled\n")
				}
			} else {
				if debug {
					fmt.Printf("\n\tVSYNC VALUE !=0 !=2 exiting\t%d",CPU.Memory[CPU.VSYNC] )
				}
				os.Exit(0)
			}

		case int8(CPU.COLUBK): //0x09
			if debug {
				fmt.Printf("\n\tCOLUBK SET! Beam index: %d\n", CPU.Beam_index)
			}
			// CPU.Pause = true

		default:

		// case 32, 16, 33, 17, 42, 13, 14, 15, 27, 28, 8, 10, 6, 5, 7, 44:
		// 	//fmt.Printf("\n\tDO NOTHING\n")

		// case -1:
		//
		// default:
		// 	fmt.Printf("\n\tInvalid CRT action %d!\n\n", action)
		// 	os.Exit(0)
	}








}


func drawBackground() {

	// Dont draw in horizontal blank
	if CPU.Beam_index*3 >= 68 {
		imd	= imdraw.New(nil)

		R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUBK])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

		// Draw
		imd.Push(pixel.V( (float64(old_BeamIndex  * 3) -68 ) * width	, float64(232-line) * height ))
		imd.Push(pixel.V( (float64(CPU.Beam_index * 3) -68 ) * width 	, float64(232-line) * height + height))
		imd.Rectangle(0)

		if debug {
			fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, CPU.Beam_index)
		}

		old_BeamIndex = CPU.Beam_index

		imd.Draw(win)

		// Count draw operations number per second
		draws ++
	}

}












func drawVisibleModeLine() {

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
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width																			, float64(232-line)*height ))
			imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height + height))
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

	imd.Push(pixel.V(  (float64(index) *pixelSize)*width								, float64(232-line)*height ))
	imd.Push(pixel.V(  (float64(index) *pixelSize)*width +float64(count*int(pixelSize))*width	, float64(232-line)*height + height))
	imd.Rectangle(0)

	imd.Draw(win)

	// Count draw operations number per second
	draws ++
}
