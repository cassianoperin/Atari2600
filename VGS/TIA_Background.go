package VGS

import (
	// "fmt"
	"math/rand"
	"image/color"
	"github.com/faiface/pixel"

)


func drawBackground() {

	// fmt.Printf("Beamer: %d\n", beamIndex)

	R, G, B := NTSC(Memory[COLUBK])
	imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}



	// // Draw
	// imd.Push(pixel.V( (float64((beamIndex-1)  * 3) -68 ) * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
	// imd.Push(pixel.V( (float64(beamIndex * 3) -68 ) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
	// imd.Rectangle(0)


	// Draw the 3 sprites of CPU cycle
	for i := 0 ; i < 3 ; i ++ {
		// TEMPORARY - Random background colors
		imd.Color = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

		// Draw
		imd.Push(pixel.V( (float64((beamIndex-1) * 3) - 68 + float64(i) ) * width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
		imd.Push(pixel.V( (float64((beamIndex-1) * 3) - 68 + float64(i) + 1) * width 	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
		imd.Rectangle(0)

		// Count draw operations number per second
		counter_DPS ++
	}






}
