package VGS

import (
	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


func readPF0() {
	for i := 4 ; i < 8 ; i++ {
		playfield[i-4] = ( Memory[PF0] >> byte(i) ) & 0x01
	}
}


func readPF1() {
	for i := 0 ; i < 8 ; i++ {
		playfield[11-i] = ( Memory[PF1] >> byte(i) ) & 0x01
	}
}


func readPF2() {
	for i := 0 ; i < 8 ; i++ {
		playfield[12+i] = ( Memory[PF2] >> byte(i) ) & 0x01
	}
}


func PF_Reflect_Duplicate() {
	// D0 = 1 = Reflect first 20 sprites of the PF to the last 20
	if (Memory[CTRLPF] & 0x01) == 1 {
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


func draw_playfield(janela_2nd_level *pixelgl.Window) {
	readPF0()
	readPF1()
	readPF2()
	PF_Reflect_Duplicate()

	// DRAW PLAYFIELD ENTIRE LINE
	for i := 0 ; i < len(playfield) ; i++ {

		if playfield[i] == 1 {


			// Check D1 status to use color of players in the score
			if (Memory[CTRLPF] & 0x02) >> 1 == 1  {
				R, G, B := NTSC(Memory[COLUP0])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

				// Set P1 Color
				if i < 20 {
					// READ COLUP1 (Memory[0x07]) - Set the Player 1 Color (On Score)
					R, G, B := NTSC(Memory[COLUP1])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				}
			// Set the Color Playfield
			} else {
				// READ COLUPF (Memory[0x08]) - Set the Playfield Color
				R, G, B := NTSC(Memory[COLUPF])
				imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
			}

			// Draw
			imd.Push(pixel.V( float64( i*4 )     * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
			imd.Push(pixel.V( float64( i*4 + 4 ) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
			imd.Rectangle(0)

			// -------- Collision Detection -------- //
			CD_P0_PF[  i*4 ]    = 1
			CD_P0_PF[ (i*4)+1 ] = 1
			CD_P0_PF[ (i*4)+2 ] = 1
			CD_P0_PF[ (i*4)+3 ] = 1

			imd.Draw(janela_2nd_level)
			counter_DPS ++

		}
	}
}
