package Graphics

import (
	"fmt"
	"time"
	"Atari2600/CPU"
	"Atari2600/Global"
	"Atari2600/Palettes"
	"image/color"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	// Control every cycle draw
	old_BeamIndex	byte = 0	// Used to draw the beam updates every cycle on the CRT
	visibleArea		bool		// Not used yet, but will be used to just draw in visible area
)

func drawBackground() {

	// Time measurement - TIA Background Draw
	if CPU.DebugTiming {
		CPU.StartTIA_BG = time.Now()
	}

	// Dont draw in horizontal blank
	if CPU.Beam_index*3 > 68 {
		// Avoid to draw if already drawed in the first STA, STY or STX cycle
		if old_BeamIndex != CPU.Beam_index {

			imd	= imdraw.New(nil)

			R, G, B := Palettes.NTSC(CPU.Memory[CPU.COLUBK])
			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

			// fmt.Printf("\n%d\n", CPU.Beam_index - old_BeamIndex)

			// Draw
			imd.Push(pixel.V( (float64(old_BeamIndex  * 3) -68 ) * width	, (screenHeight * (1 - Global.SizeYused)) + float64(232-line) * height ))
			imd.Push(pixel.V( (float64(CPU.Beam_index * 3) -68 ) * width 	, (screenHeight * (1 - Global.SizeYused)) + float64(232-line) * height + height))
			imd.Rectangle(0)

			// if debug {
			// 	fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, CPU.Beam_index)
			// }

			imd.Draw(Global.Win)

			// Count draw operations number per second
			draws ++
		}
	}

	old_BeamIndex = CPU.Beam_index

	// Time measurement - TIA Background Draw
	if CPU.DebugTiming {
		elapsedBG := time.Since(CPU.StartTIA_BG)
		if elapsedBG.Seconds() > CPU.DebugTimingLimit {
			fmt.Printf("\tOpcode: %X\tBackground Draw took %f seconds\n", CPU.Opcode, elapsedBG.Seconds())
			// CPU.Pause = true
		}
	}
}
