package Graphics

import (
	"os"
	"fmt"
	"Atari2600/CPU"
	"Atari2600/Global"
	"image/color"
	"github.com/faiface/pixel"
	"Atari2600/Palettes"
)

var (
	// Players Vertical Positioning
	XPositionP0			byte
	XFinePositionP0			int8
	XPositionP1			byte
	XFinePositionP1			int8

	// Collision Detection
	CD_debug			bool	= true	// Debug
	CD_P0_P1			[160]byte
	CD_P0_P1_collision_detected			bool	= false	// Set when collision is detected
	CD_P0_PF			[160]byte
	CD_P0_PF_collision_detected			bool	= false	// Set when collision is detected
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

func drawLineScroll(currentPosition int16) int16 {
	// Left horizontal scroll
	if currentPosition < 0 {
		currentPosition = 160 + currentPosition
	// Right horizontal scroll
	} else if currentPosition > 159 {
		currentPosition = currentPosition % 160
	}

	return currentPosition
}

func drawPlayer(player byte) {
	var (
		bit				byte = 0
		inverted			byte = 0
		// P0 and P1 registers
		register_REFP		byte
		register_GRP		byte
		register_COLUP		byte
		register_NUSIZ		byte
		XPosition		byte
		XFinePosition		int8
		drawLine		float64
		drawLinePosition	int16
		drawLinePosition2	int16
		drawLinePosition3	int16
		drawLinePosition4	int16
	)

	// Tests
	// CPU.Memory[CPU.NUSIZ0] = 0x07
	// CPU.Memory[CPU.NUSIZ1] = 0x03
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
		XPosition	= XPositionP0
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
		XPosition	= XPositionP1
		XFinePosition	= XFinePositionP1
	}


	// ----- Draw Player Line ----- //
	//Global.ScreenHeight * (1 - Global.SizeYused))  used to add the draw to the top of screen

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

			// ----------------------------------------------- NUSIZx = 0x00 ----------------------------------------------- //
			if register_NUSIZ == 0x00 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width				, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
								// fmt.Println(CD_P0_PF)
							}

						} else {
							CD_P0_PF[ drawLinePosition ] = 1
						}
					}
				}


			// ----------------------------------------------- NUSIZx = 0x01 ----------------------------------------------- //
			} else if register_NUSIZ == 0x01 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 16

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
						}
					}
				}


			// ----------------------------------------------- NUSIZx = 0x02 ----------------------------------------------- //
			} else if register_NUSIZ == 0x02 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 32

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	,(Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
						}
					}
				}

			// ----------------------------------------------- NUSIZx = 0x03 ----------------------------------------------- //
			} else if register_NUSIZ == 0x03 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 16
				drawLinePosition3 = drawLinePosition + 32

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)
				drawLinePosition3 = drawLineScroll(drawLinePosition3)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition3 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition3 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 || CD_P0_P1[ drawLinePosition3 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
						CD_P0_P1[ drawLinePosition3 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 || CD_P0_PF[ drawLinePosition3 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
							CD_P0_PF[ drawLinePosition3 ] = 1
						}
					}
				}


			// ----------------------------------------------- NUSIZx = 0x04 ----------------------------------------------- //
			} else if register_NUSIZ == 0x04 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 64

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) +  drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
						}
					}
				}

			// ----------------------------------------------- NUSIZx = 0x05 ----------------------------------------------- //
			} else if register_NUSIZ == 0x05 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i*2) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 1

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						// Fill the 2 bytes drawed
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 {

							fmt.Println(drawLinePosition)
							fmt.Println(CD_P0_PF[ drawLinePosition ])
							fmt.Println(drawLinePosition2)
							fmt.Println(CD_P0_PF[ drawLinePosition2 ])
							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							// Fill the 2 bytes drawed
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
						}
					}
				}


			// ----------------------------------------------- NUSIZx = 0x06 ----------------------------------------------- //
			} else if register_NUSIZ == 0x06 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 32
				drawLinePosition3 = drawLinePosition + 64

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)
				drawLinePosition3 = drawLineScroll(drawLinePosition3)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition3) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition3) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)


				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 || CD_P0_P1[ drawLinePosition3 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
						CD_P0_P1[ drawLinePosition3 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 || CD_P0_PF[ drawLinePosition3 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
							CD_P0_PF[ drawLinePosition3 ] = 1
						}
					}
				}


			// ----------------------------------------------- NUSIZx = 0x07 ----------------------------------------------- //
			} else if register_NUSIZ == 0x07 {

				// --------------------------- Draw ---------------------------- //
				drawLinePosition  = int16(XPosition*3) - 68 + int16(i*4) + int16(XFinePosition)
				drawLinePosition2 = drawLinePosition + 1
				drawLinePosition3 = drawLinePosition + 2
				drawLinePosition4 = drawLinePosition + 3

				// If value < 0 or > 159, scroll draw position
				drawLinePosition  = drawLineScroll(drawLinePosition)
				drawLinePosition2 = drawLineScroll(drawLinePosition2)
				drawLinePosition3 = drawLineScroll(drawLinePosition3)
				drawLinePosition4 = drawLineScroll(drawLinePosition4)

				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition2 ) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition3) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition3) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( drawLinePosition4) ) * Global.Width		, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height ))
				imd.Push(pixel.V( (float64( drawLinePosition4) ) * Global.Width + Global.Width	, (Global.ScreenHeight * (1 - Global.SizeYused)) + drawLine * Global.Height + Global.Height))
				imd.Rectangle(0)

				// -------------------- Collision Detection -------------------- //

				// CXPPMM (D7) - P0-P1
				if !CD_P0_P1_collision_detected {
					if CD_P0_P1[ drawLinePosition ] ==  1 || CD_P0_P1[ drawLinePosition2 ] ==  1 || CD_P0_P1[ drawLinePosition3 ] ==  1 || CD_P0_P1[ drawLinePosition4 ] ==  1 {

						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Inform TIA that does not need to check collisions anymore in this frame
						CD_P0_P1_collision_detected = true

						if CD_debug {
							fmt.Println("Collision Detection: P0-P1 Detected!")
						}

					} else {
						// Fill the 4 bytes drawed
						CD_P0_P1[ drawLinePosition ]  = 1
						CD_P0_P1[ drawLinePosition2 ] = 1
						CD_P0_P1[ drawLinePosition3 ] = 1
						CD_P0_P1[ drawLinePosition4 ] = 1
					}
				}

				// CXP0FB (D7) - P0-PF
				if player == 0 {
					if !CD_P0_PF_collision_detected {
						if CD_P0_PF[ drawLinePosition ] ==  1 || CD_P0_PF[ drawLinePosition2 ] ==  1 || CD_P0_PF[ drawLinePosition3 ] ==  1 || CD_P0_PF[ drawLinePosition4 ] ==  1 {

							CPU.MemTIAWrite[CPU.CXP0FB] = 0x80

							// Inform TIA that does not need to check collisions anymore in this frame
							CD_P0_PF_collision_detected = true

							if CD_debug {
								fmt.Println("Collision Detection: P0-PF Detected!")
							}

						} else {
							// Fill the 4 bytes drawed
							CD_P0_PF[ drawLinePosition ]  = 1
							CD_P0_PF[ drawLinePosition2 ] = 1
							CD_P0_PF[ drawLinePosition3 ] = 1
							CD_P0_PF[ drawLinePosition4 ] = 1
						}
					}
				}

			}
		}
	}

	// Debug Collision Detection arrays
	// fmt.Println(CD_P0_P1)
	// fmt.Println(CD_P0_PF)

	imd.Draw(Global.Win)
	// Count draw operations number per second
	draws ++
}
