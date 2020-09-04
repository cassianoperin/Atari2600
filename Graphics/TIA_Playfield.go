package Graphics

import (
	"github.com/faiface/pixel"
	"Atari2600/Palettes"
	"Atari2600/CPU"
	"image/color"
)

var (
	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]byte			//Improve to binary
	pixelSize			float64 = 4.0		// 80 lines (half screen) / 20 PF0, PF1 and PF2 bits
)


func readPF0() {
	for i := 4 ; i < 8 ; i++ {
		playfield[i-4] = ( CPU.Memory[CPU.PF0] >> byte(i) ) & 0x01
	}
}


func readPF1() {
	for i := 0 ; i < 8 ; i++ {
		playfield[11-i] = ( CPU.Memory[CPU.PF1] >> byte(i) ) & 0x01
	}
}


func readPF2() {
	for i := 0 ; i < 8 ; i++ {
		playfield[12+i] = ( CPU.Memory[CPU.PF2] >> byte(i) ) & 0x01
	}
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


func draw_playfield() {
	readPF0()
	readPF1()
	readPF2()
	PF_Reflect_Duplicate()

	// DRAW PLAYFIELD ENTIRE LINE
	for i := 0 ; i < len(playfield) ; i++ {

		if playfield[i] == 1 {


			// Check D1 status to use color of players in the score
			if (CPU.Memory[CPU.CTRLPF] & 0x02) >> 1 == 1  {
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP0])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				// Set P1 Color
				if i < 20 {
					// READ COLUP1 (Memory[0x07]) - Set the Player 1 Color (On Score)
					R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUP1])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				}
			// Set the Color Playfield
			} else {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUPF])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			}

			// Draw
			imd.Push(pixel.V( float64( i*4 )     * width	, float64(232-line) * height ))
			imd.Push(pixel.V( float64( i*4 + 4 ) * width 	, float64(232-line) * height + height))
			imd.Rectangle(0)

			// -------- Collision Detection -------- //
			CD_P0_PF[  i*4 ]    = 1
			CD_P0_PF[ (i*4)+1 ] = 1
			CD_P0_PF[ (i*4)+2 ] = 1
			CD_P0_PF[ (i*4)+3 ] = 1

			imd.Draw(win)
			draws ++

		}
	}
}
