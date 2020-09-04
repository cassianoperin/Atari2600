package Graphics

import (
	"os"
	"fmt"
	"github.com/faiface/pixel"
	"Atari2600/Palettes"
	"Atari2600/CPU"
	"image/color"
)

var (
	// Players Vertical Positioning
	XPositionP0				byte
	XFinePositionP0			int8
	XPositionP1				byte
	XFinePositionP1			int8

	// Collision Detection
	CD_debug				bool	= true	// Debug
	CD_P0_P1				[160]byte		// Line Array // TODO FIX TO 160 DUE TO PLAYER POSITION ON SCREEN START IN 1
	CD_P0_P1_status			bool 	= false	// Set when collision is detected
)


func invert_bits(value byte) byte {
	// fmt.Printf("\n\n%08b", value)
	value = (value & 0xF0) >> 4 | (value & 0x0F) << 4;
	value = (value & 0xCC) >> 2 | (value & 0x33) << 2;
	value = (value & 0xAA) >> 1 | (value & 0x55) << 1;
	// fmt.Printf("\n\n%08b", value)

	return value;
}


func Fine(HMPX byte) int8 {

	var value int8

	switch HMPX {
		case 0x70:
			value = -7
		case 0x60:
			value = -6
		case 0x50:
			value = -5
		case 0x40:
			value = -4
		case 0x30:
			value = -3
		case 0x20:
			value = -2
		case 0x10:
			value = -1
		case 0x00:
			value =  0
		case 0xF0:
			value =  1
		case 0xE0:
			value =  2
		case 0xD0:
			value =  3
		case 0xC0:
			value =  4
		case 0xB0:
			value =  5
		case 0xA0:
			value =  6
		case 0x90:
			value =  7
		case 0x80:
			value =  8
		default:
			fmt.Printf("\n\tInvalid HMP0 %02X!\n\n", CPU.HMP0)
			os.Exit(0)
	}

	return value

}

func drawPlayer(player byte) {
	var (
		bit					byte = 0
		inverted			byte = 0
		// P0 and P1 registers
		register_REFP		byte
		register_GRP		byte
		register_COLUP		byte
		register_NUSIZ		byte
		XPosition			byte
		XFinePosition		int8
		drawLine			float64
		drawLinePosition	byte


	)

	// // TESTEEEEEE SPRITES DO MESMO TAMANHO
	CPU.Memory[CPU.NUSIZ0] = 0x06
	CPU.Memory[CPU.NUSIZ1] = 0x03
	// CPU.Memory[CPU.NUSIZ1] = 0x07

	// Configs for Drawing P0
	if player == 0 {
		// If a program doesnt use RESP0, initialize
		if XPositionP0 == 0 {
			XPositionP0 = 23
		}

		if debug {
			fmt.Printf("Line: %d\tGRP0: %08b\tXPositionP0: %d\tHMP0: %d\n", line, CPU.Memory[CPU.GRP0], XPositionP0, CPU.Memory[CPU.HMP0])
		}

		register_REFP	= CPU.Memory[CPU.REFP0]
		register_GRP	= CPU.Memory[CPU.GRP0]
		register_COLUP	= CPU.Memory[CPU.COLUP0]
		register_NUSIZ	= CPU.Memory[CPU.NUSIZ0]
		XPosition		= XPositionP0
		XFinePosition	= XFinePositionP0

	// Configs for Drawing P1
	} else {
		// If a program doesnt use RESP0, initialize (Initial Player Position)
		if XPositionP1 == 0 {
			XPositionP1 = 30
		}

		if debug {
			fmt.Printf("Line: %d\tGRP1: %08b\tXPositionP1: %d\tHMP1: %d\n", line, CPU.Memory[CPU.GRP1], XPositionP1, CPU.Memory[CPU.HMP1])
		}

		register_REFP	= CPU.Memory[CPU.REFP1]
		register_GRP	= CPU.Memory[CPU.GRP1]
		register_COLUP	= CPU.Memory[CPU.COLUP1]
		register_NUSIZ	= CPU.Memory[CPU.NUSIZ1]
		XPosition		= XPositionP1
		XFinePosition	= XFinePositionP1
	}


	// ----- Draw Player Line ----- //
	// Set the line
	drawLine = 232 - float64(line)

	// For each bit in GRPn, draw if == 1
	for i:=0 ; i <=7 ; i++ {

		// handle the order of the bits (normal or inverted)
		if register_REFP == 0x00 {
			bit = register_GRP >> (7-byte(i)) & 0x01
		} else {
			// If Reflect Player Enabled (REFP1), invert the order of GRPn
			inverted = invert_bits(register_GRP)
			bit = inverted >> (7-byte(i)) & 0x01
		}

		if bit == 1 {
			// READ COLUPF (Memory[0x08]) - Set the Playfield Color
			R, G, B := Palettes.NTSC(register_COLUP)
			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

			// NUSIZx = 0x00
			if register_NUSIZ == 0x00 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ] = 1
						}
					}
				}

			// NUSIZx = 0x01
			} else if register_NUSIZ == 0x01 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition + 16) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition + 16) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 16 ] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +16 >= 0 && drawLinePosition +16 < 160 {
							CD_P0_P1[ drawLinePosition + 16 ] = 1
						}
					}
				}


			// NUSIZx = 0x02
			} else if register_NUSIZ == 0x02 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition + 32 ) ) * width			,	 drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition + 32 ) ) * width + width	,	 drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 32 ] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +32 >= 0 && drawLinePosition +32 < 160 {
							CD_P0_P1[ drawLinePosition + 32 ] = 1
						}
					}
				}

			// NUSIZx = 0x03
			} else if register_NUSIZ == 0x03 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition + 16 ) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition + 16 ) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition + 32 ) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition + 32 ) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 16 ] ==  1 || CD_P0_P1[ drawLinePosition + 32 ] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +16 >= 0 && drawLinePosition +16 < 160 {
							CD_P0_P1[ drawLinePosition + 16 ] = 1
						}
						if drawLinePosition +32 >= 0 && drawLinePosition +32 < 160 {
							CD_P0_P1[ drawLinePosition + 32 ] = 1
						}
					}
				}

			// NUSIZx = 0x04
			} else if register_NUSIZ == 0x04 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition + 64 ) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition + 64 ) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 64 ] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +64 >= 0 && drawLinePosition +64 < 160 {
							CD_P0_P1[ drawLinePosition + 64 ] = 1
						}
					}
				}

			// NUSIZx = 0x05
			} else if register_NUSIZ == 0x05 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i*2) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width					, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + (width*2)		, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 1] ==  1{
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_status = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						// Fill the 2 bytes drawed
						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +1 >= 0 && drawLinePosition +1 < 160 {
							CD_P0_P1[ drawLinePosition + 1 ] = 1
						}
					}
				}

			// NUSIZx = 0x06
			} else if register_NUSIZ == 0x06 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition % 160 ) ) * width				, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + width		, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (drawLinePosition + 32)  % 160  ) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( (drawLinePosition + 32)  % 160 ) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (drawLinePosition + 64)  % 160 ) ) * width			, drawLine * height ))
				imd.Push(pixel.V( (float64( (drawLinePosition + 64)  % 160 ) ) * width + width	, drawLine * height + height))
				imd.Rectangle(0)
				fmt.Println(drawLinePosition)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				// if !CD_P0_P1_status {
				// 	if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 32 ] ==  1 || CD_P0_P1[ drawLinePosition + 64 ] ==  1 {
				// 		// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
				// 		CPU.MemTIAWrite[CPU.CXPPMM] = 0x80
				//
				// 		if CD_debug {
				// 			fmt.Println("Collision Detection: P0-P1 Detected!")
				// 			// fmt.Println(CD_P0_P1)
				// 		}
				//
				// 	} else {
				// 		if drawLinePosition >= 0 && drawLinePosition < 160 {
				// 			CD_P0_P1[ drawLinePosition ]      = 1
				// 		}
				// 		if drawLinePosition +32 >= 0 && drawLinePosition +32 < 160 {
				// 			CD_P0_P1[ drawLinePosition + 32 ] = 1
				// 		}
				// 		if drawLinePosition +64 >= 0 && drawLinePosition +64 < 160 {
				// 			CD_P0_P1[ drawLinePosition + 64 ] = 1
				// 		}
				// 	}
				// }


			// NUSIZx = 0x07
			} else if register_NUSIZ == 0x07 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = (XPosition*3) - 68 + byte(i*4) + byte(XFinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width					, drawLine * height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * width + (width*4)		, drawLine * height + height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //
				// If collition not detected yet in this frame, check for collisions
				if !CD_P0_P1_status {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition + 1 ] ==  1 || CD_P0_P1[ drawLinePosition + 2 ] ==  1 || CD_P0_P1[ drawLinePosition + 3] ==  1 {
						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_status = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
							// fmt.Println(CD_P0_P1)
						}

					} else {
						// Fill the 4 bytes drawed
						CD_P0_P1[ drawLinePosition ]     = 1
						CD_P0_P1[ drawLinePosition + 1 ] = 1
						CD_P0_P1[ drawLinePosition + 2 ] = 1
						CD_P0_P1[ drawLinePosition + 3 ] = 1

						if drawLinePosition >= 0 && drawLinePosition < 160 {
							CD_P0_P1[ drawLinePosition ]      = 1
						}
						if drawLinePosition +1 >= 0 && drawLinePosition +1 < 160 {
							CD_P0_P1[ drawLinePosition + 1 ] = 1
						}
						if drawLinePosition +2 >= 0 && drawLinePosition +2 < 160 {
							CD_P0_P1[ drawLinePosition + 2 ] = 1
						}
						if drawLinePosition +3 >= 0 && drawLinePosition +3 < 160 {
							CD_P0_P1[ drawLinePosition + 3 ] = 1
						}
					}
				}


			}
		}
	}

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}
