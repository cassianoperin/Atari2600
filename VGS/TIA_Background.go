package VGS

import (
	"fmt"
	"time"
	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)


func drawBackground() {

	// Time measurement - TIA Background Draw
	if debugTiming {
		debugTiming_StartTIA_BG = time.Now()
	}

	// Dont draw in horizontal blank
	if beamIndex * 3 > 68 {
		// Avoid to draw if already drawed in the first STA, STY or STX cycle
		if old_BeamIndex != beamIndex {

			imd	= imdraw.New(nil)

			R, G, B := NTSC(Memory[COLUBK])
			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

			// fmt.Printf("\n%d\n", Beam_index - old_BeamIndex)

			// Draw
			imd.Push(pixel.V( (float64(old_BeamIndex  * 3) -68 ) * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
			imd.Push(pixel.V( (float64(beamIndex * 3) -68 ) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
			imd.Rectangle(0)

			// if debug {
			// 	fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, Beam_index)
			// }

			imd.Draw(win)

			// Count draw operations number per second
			counter_DPS ++
		}
	}

	old_BeamIndex = beamIndex

	// Time measurement - TIA Background Draw
	if debugTiming {
		elapsedBG := time.Since(debugTiming_StartTIA_BG)
		if elapsedBG.Seconds() > debugTiming_Limit {
			fmt.Printf("\tOpcode: %X\tBackground Draw took %f seconds\n", opcode, elapsedBG.Seconds())
			// Pause = true
		}
	}
}
