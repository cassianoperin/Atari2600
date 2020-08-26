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
	XPositionP0		byte
	XFinePositionP0	int8
	XPositionP1		byte
	XFinePositionP1	int8
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


func drawPlayer0() {
	var (
		bit			byte = 0
		inverted	byte = 0
	)

	// If a program doesnt use RESP0, initialize
	if XPositionP0 == 0 {
		XPositionP0 = 23
	}

	if debug {
		fmt.Printf("Line: %d\tGRP0: %08b\tXPositionP0: %d\tXFinePositionP0: %d\n", line, CPU.Memory[CPU.GRP0], XPositionP0, XFinePositionP0)
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
				imd.Push(pixel.V( (float64( ((XPositionP0)*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( ((XPositionP0)*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				// Collision Detection - Sprite size
				CD_P0_sprite_size = 8
				// Copy of the Sprite
				CD_P0_sprite_copy = uint16(CPU.Memory[CPU.GRP0])
			} else if CPU.Memory[CPU.NUSIZ0] == 0x01 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(16) )*width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(16) )*width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ0] == 0x02 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ0] == 0x03 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(16) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(16) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ0] == 0x04 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ0] == 0x05 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i*2)) + float64(XFinePositionP0) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i*2)) + float64(XFinePositionP0) ) * width + (width*2)			, float64(232-line) * height + height))
				imd.Rectangle(0)
				// Collision Detection - Sprite size
				CD_P0_sprite_size = 16

				// Copy of the sprite
				// var mask byte = 128
				// CD_P1_sprite_copy = 0
				// for i:=8 ; i > 0 ; i-- {
				// 	// fmt.Println(i)
				// 	CD_P1_sprite_copy += ( uint16(CPU.Memory[CPU.GRP1] & (mask >> (8-i) ) ) << i )
				// 	CD_P1_sprite_copy += ( uint16(CPU.Memory[CPU.GRP1] & (mask >> (8-i) ) ) << (i-1) )
				// }
				// fmt.Printf("COPY: %08b\n%016b\n", uint16(CPU.Memory[CPU.GRP1]), CD_P1_sprite_copy)

			} else if CPU.Memory[CPU.NUSIZ0] == 0x06 {
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP0*3) - 68 + byte(i)) + float64(XFinePositionP0) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ0] == 0x07 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*4)) + float64(XFinePositionP0) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*4)) + float64(XFinePositionP0) ) * width + (width*4)			, float64(232-line) * height + height))
				imd.Rectangle(0)
			}

		}
	}

	// Set variables used in Collision Detection
	CD_drawP0_currentline	= true
	CD_P0_StartPos			= uint8( ( (int8(XPositionP0) * 3) - 68 ) + XFinePositionP0 )
	CD_P0_EndPos			= uint8( ( (int8(XPositionP0) * 3) - 68 ) + XFinePositionP0 + int8(CD_P0_sprite_size) )

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}


func drawPlayer1() {
	var (
		bit			byte = 0
		inverted	byte = 0
	)

	// TESTEEEEEE SPRITES DO MESMO TAMANHO
	CPU.Memory[CPU.NUSIZ1] = 0x00

	// If a program doesnt use RESP0, initialize (Initial Player Position)
	if XPositionP1 == 0 {
		XPositionP1 = 30
	}

	if debug {
		fmt.Printf("Line: %d\tGRP1: %08b\tXPositionP1: %d\tHMP1: %d\n", line, CPU.Memory[CPU.GRP1], XPositionP1, CPU.Memory[CPU.HMP1])
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
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				// Collision Detection - Sprite size
				CD_P1_sprite_size = 8
				// Copy of the sprite
				CD_P1_sprite_copy = uint16(CPU.Memory[CPU.GRP1])
			} else if CPU.Memory[CPU.NUSIZ1] == 0x01 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(16) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(16) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ1] == 0x02 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ1] == 0x03 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) + float64(16) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) + float64(16) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) +float64(XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ1] == 0x04 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ1] == 0x05 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*2)) + float64(XFinePositionP1) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*2)) + float64(XFinePositionP1) ) * width + (width*2)			, float64(232-line) * height + height))
				imd.Rectangle(0)
				// Collision Detection - Sprite size
				CD_P1_sprite_size = 16

				// Copy of the sprite
				// var mask byte = 128
				// CD_P1_sprite_copy = 0
				// for i:=8 ; i > 0 ; i-- {
				// 	// fmt.Println(i)
				// 	CD_P1_sprite_copy += ( uint16(CPU.Memory[CPU.GRP1] & (mask >> (8-i) ) ) << i )
				// 	CD_P1_sprite_copy += ( uint16(CPU.Memory[CPU.GRP1] & (mask >> (8-i) ) ) << (i-1) )
				// }
				// fmt.Printf("COPY: %08b\n%016b\n", uint16(CPU.Memory[CPU.GRP1]), CD_P1_sprite_copy)

			} else if CPU.Memory[CPU.NUSIZ1] == 0x06 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i)) + float64(XFinePositionP1) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if CPU.Memory[CPU.NUSIZ1] == 0x07 {
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*4)) + float64(XFinePositionP1) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPositionP1*3) - 68 + byte(i*4)) + float64(XFinePositionP1) ) * width + (width*4)			, float64(232-line) * height + height))
				imd.Rectangle(0)
			}
		}
	}

	// Set variables used in Collision Detection
	CD_drawP1_currentline	= true
	CD_P1_StartPos			= uint8( ( (int8(XPositionP1) * 3) - 68 ) + XFinePositionP1 )
	CD_P1_EndPos			= uint8( ( (int8(XPositionP1) * 3) - 68 ) + XFinePositionP1 + int8(CD_P1_sprite_size) )

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}
