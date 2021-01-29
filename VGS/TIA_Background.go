package VGS

import (
	"math/rand"
	"image/color"
	"github.com/faiface/pixel"

)


func drawBackground() {

	// Avoid to draw if already drawed in the first STA, STY or STX cycle
	if old_beamIndex != beamIndex {

		// KEY TO SPEED
		// imd	= imdraw.New(nil)

		R, G, B := NTSC(Memory[COLUBK])
		imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

		// TEMPORARY - Random background colors
		imd.Color = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

		// Draw
		imd.Push(pixel.V( (float64(old_beamIndex  * 3) -68 ) * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
		imd.Push(pixel.V( (float64(beamIndex * 3) -68 ) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
		imd.Rectangle(0)

		// Count draw operations number per second
		counter_DPS ++

		// Update the Old Beam index
		old_beamIndex = beamIndex
	}

}


// func drawBackground(janela_2nd_level *pixelgl.Window) {
//
//
//
// 	// Time measurement - TIA Background Draw
// 	if debugTiming {
// 		debugTiming_StartTIA_BG = time.Now()
// 	}
//
// 	// Dont draw in horizontal blank
// 	if beamIndex * 3 > 68 {
// 		// Avoid to draw if already drawed in the first STA, STY or STX cycle
// 		if old_beamIndex != beamIndex {
//
// 			// Pause = true
//
// 			// KEY TO SPEED
// 			imd	= imdraw.New(nil)
//
// 			R, G, B := NTSC(Memory[COLUBK])
// 			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
//
// 			if ipressed == true {
// 				imd.Color = color.RGBA{uint8(100), uint8(100), uint8(100), 255}
//
// 			}
// 			// fmt.Printf("\n%d\n", beamIndex - old_beamIndex)
//
// 			// Draw
// 			imd.Push(pixel.V( (float64(old_beamIndex  * 3) -68 ) * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
// 			imd.Push(pixel.V( (float64(beamIndex * 3) -68 ) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
// 			imd.Rectangle(0)
//
// 			// if debug {
// 			// 	fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_beamIndex, Beam_index)
// 			// }
//
// 			imd.Draw(janela_2nd_level)
// 			fmt.Printf("\n\n\nENTROU\n\n\n")
// 			// Count draw operations number per second
// 			counter_DPS ++
// 		}
// 	}
//
// 	old_beamIndex = beamIndex
//
// 	// Time measurement - TIA Background Draw
// 	if debugTiming {
// 		elapsedBG := time.Since(debugTiming_StartTIA_BG)
// 		if elapsedBG.Seconds() > debugTiming_Limit {
// 			fmt.Printf("\tOpcode: %X\tBackground Draw took %f seconds\n", opcode, elapsedBG.Seconds())
// 			// Pause = true
// 		}
// 	}
// }
