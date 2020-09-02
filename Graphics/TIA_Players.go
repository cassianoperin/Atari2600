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

func drawPlayer(player byte) {
	var (
		bit				byte = 0
		inverted		byte = 0
		// P0 and P1 registers
		register_REFP	byte
		register_GRP	byte
		register_COLUP	byte
		register_NUSIZ	byte
		XPosition		byte
		XFinePosition	int8
	)

	// // TESTEEEEEE SPRITES DO MESMO TAMANHO
	// CPU.Memory[CPU.NUSIZ0] = 0x05
	// CPU.Memory[CPU.NUSIZ1] = 0x00

	// Configs for Drawing P0
	if player == 0 {
		// If a program doesnt use RESP0, initialize
		if XPositionP0 == 0 {
			XPositionP0 = 23
		}

		// if debug {
			fmt.Printf("Line: %d\tGRP0: %08b\tXPositionP0: %d\tHMP0: %d\n", line, CPU.Memory[CPU.GRP0], XPositionP0, CPU.Memory[CPU.HMP0])
		// }

		//
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

		// if debug {
			fmt.Printf("Line: %d\tGRP1: %08b\tXPositionP1: %d\tHMP1: %d\n", line, CPU.Memory[CPU.GRP1], XPositionP1, CPU.Memory[CPU.HMP1])
		// }

		register_REFP	= CPU.Memory[CPU.REFP1]
		register_GRP	= CPU.Memory[CPU.GRP1]
		register_COLUP	= CPU.Memory[CPU.COLUP1]
		register_NUSIZ	= CPU.Memory[CPU.NUSIZ1]
		XPosition		= XPositionP1
		XFinePosition	= XFinePositionP1
	}

	// ----- Collision Detection Variables ----- //
	if register_NUSIZ == 0x00 {

		// Sprite size and copy of the sprite
		if player == 0 {
			CD_P0_sprite_size = 8						// P0 Sprite size
			CD_P0_sprite_copy = uint32(register_GRP)	// P0 sprite after NUSIZ
		} else {
			CD_P1_sprite_size = 8						// P1 sprite size
			CD_P1_sprite_copy = uint32(register_GRP)	// P1 sprite after NUSIZ
		}

	} else if register_NUSIZ == 0x01 {
		// Implement
	} else if register_NUSIZ == 0x02 {
		// Implement
	} else if register_NUSIZ == 0x03 {
		// Implement
	} else if register_NUSIZ == 0x04 {
		// Implement
	} else if register_NUSIZ == 0x05 {

		// Sprite size
		if player == 0 {
			CD_P0_sprite_size = 16
			CD_P0_sprite_copy = 0	// clean
		} else {
			CD_P1_sprite_size = 16
			CD_P1_sprite_copy = 0	// clean
		}

		// Copy of the sprite
		var mask byte = 128
		for j:=8 ; j > 0 ; j-- {
			if player == 0 {
				CD_P0_sprite_copy += ( uint32(register_GRP & (mask >> (8-j) ) ) <<  j    )
				CD_P0_sprite_copy += ( uint32(register_GRP & (mask >> (8-j) ) ) << (j-1) )
			} else {
				CD_P1_sprite_copy += ( uint32(register_GRP & (mask >> (8-j) ) ) <<  j    )
				CD_P1_sprite_copy += ( uint32(register_GRP & (mask >> (8-j) ) ) << (j-1) )
			}
		}
		// fmt.Printf("COPY: %08b\n%016b\n", uint16(CPU.Memory[CPU.GRP1]), CD_P1_sprite_copy)

	} else if register_NUSIZ == 0x06 {
		// Implement
	} else if register_NUSIZ == 0x07 {
		// Implement
	}


	// ----- Draw Player Line ----- //
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

			if register_NUSIZ == 0x00 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x01 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(16) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(16) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x02 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x03 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) + float64(16) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) + float64(16) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) +float64(XFinePosition) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x04 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x05 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i*2)) + float64(XFinePosition) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i*2)) + float64(XFinePosition) ) * width + (width*2)			, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x06 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width						, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) ) * width + width				, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(32) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(32) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(64) ) * width			, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i)) + float64(XFinePosition) + float64(64) ) * width + width	, float64(232-line) * height + height))
				imd.Rectangle(0)
			} else if register_NUSIZ == 0x07 {
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i*4)) + float64(XFinePosition) ) * width					, float64(232-line) * height ))
				imd.Push(pixel.V( (float64( (XPosition*3) - 68 + byte(i*4)) + float64(XFinePosition) ) * width + (width*4)			, float64(232-line) * height + height))
				imd.Rectangle(0)
			}
		}
	}

	// Set variables used in Collision Detection
	// Collision Detection - Sprite size
	if player == 0 {
		CD_drawP0_currentline	= true
		CD_P0_StartPos			= uint8( ( (int8(XPosition) * 3) - 68 ) + XFinePosition )
		CD_P0_EndPos			= uint8( ( (int8(XPosition) * 3) - 68 ) + XFinePosition + int8(CD_P0_sprite_size) )
	} else {
		CD_drawP1_currentline	= true
		CD_P1_StartPos			= uint8( ( (int8(XPosition) * 3) - 68 ) + XFinePosition )
		CD_P1_EndPos			= uint8( ( (int8(XPosition) * 3) - 68 ) + XFinePosition + int8(CD_P1_sprite_size) )
	}

	imd.Draw(win)
	// Count draw operations number per second
	draws ++
}
