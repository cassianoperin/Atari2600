package VGS

import (
	// "fmt"
	// "math/rand"
	"image/color"
	"github.com/faiface/pixel"
)


func drawBackground() {

	// --------------------------- Draw the 3 sprites of CPU cycle --------------------------- //
	for i := 0 ; i < 3 ; i ++ {

		// Define the pixel position
		pixel_position := ( (int(beamIndex)-1) * 3 ) - 68 + i + 1

		// Dont draw first two sprites outside screen (-2 and -1 X position)
		if pixel_position > 0 {

			// Set the background color as default
			R, G, B := NTSC(Memory[COLUBK])
			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

			// Memory[PF0] = 80

			// ----------------------------------------- PF0 ----------------------------------------- //
			if pixel_position <= 16 {

				// fmt.Printf("%d\tPF0: %b\t%b\tPF0_BIT: %d\n", pixel_position, Memory[PF0], ( Memory[PF0] >> byte(pf0_bit) ) & 0x01, pf0_bit)

				// If the bit is 1, set the color of the playfield
				if ( Memory[PF0] >> byte(pf0_bit) ) & 0x01 == 1 {
					R, G, B := NTSC(Memory[COLUPF])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
					// fmt.Println("ENTROU")
					pf0_bit ++
				}

				// Reset PF0 bit index for the next line
				if pixel_position == 16 {
					pf0_bit = 4
				}

			// ----------------------------------------- PF1 ----------------------------------------- //
			} else if pixel_position <= 48 {

				// Memory[PF1] = 161

				// fmt.Printf("%d\tPF1: %b\t%b\tPF1_BIT: %d\n", pixel_position, Memory[PF1], ( Memory[PF1] >> byte(pf1_bit) ) & 0x01, pf1_bit)

				// If the bit is 1, set the color of the playfield
				if ( Memory[PF1] >> byte(pf1_bit) ) & 0x01 == 1 {
					R, G, B := NTSC(Memory[COLUPF])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
					// fmt.Println("ENTROU")
					pf1_bit --
				}

				// Reset PF1 bit index for the next line
				if pixel_position == 48 {
					pf1_bit = 7
				}

			// ----------------------------------------- PF2 ----------------------------------------- //
			} else if pixel_position <= 80 {

				// Memory[PF2] = 161

				// fmt.Printf("%d\tPF2: %b\t%b\tPF2_BIT: %d\n", pixel_position, Memory[PF2], ( Memory[PF2] >> byte(pf2_bit) ) & 0x01, pf2_bit)

				// If the bit is 1, set the color of the playfield
				if ( Memory[PF2] >> byte(pf2_bit) ) & 0x01 == 1 {
					R, G, B := NTSC(Memory[COLUPF])
					imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
					// fmt.Println("ENTROU")
					pf2_bit ++
				}

				// Reset PF1 bit index for the next line
				if pixel_position == 80 {
					pf2_bit = 0
				}
			}


			// ---------------------------------- Reflect Playfield ---------------------------------- //
			if (Memory[CTRLPF] & 0x01) == 1 {


				if pixel_position > 80 {

					// ------------------------------------ PF0 Reflected ------------------------------------ //
					if pixel_position <= 96 {

						// fmt.Printf("%d\tPF0: %b\t%b\tPF0_BIT: %d\n", pixel_position, Memory[PF0], ( Memory[PF0] >> byte(pf0_bit) ) & 0x01, pf0_bit)

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF0] >> byte(pf0_bit) ) & 0x01 == 1 {
							R, G, B := NTSC(Memory[COLUPF])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							// fmt.Println("ENTROU")
							pf0_bit ++
						}

						// Reset PF0 bit index for the next line
						if pixel_position == 96 {
							pf0_bit = 4
						}

					// ------------------------------------ PF1 Reflected ------------------------------------ //
					} else if pixel_position <= 128 {

						// Memory[PF1] = 161

						// fmt.Printf("%d\tPF1: %b\t%b\tPF1_BIT: %d\n", pixel_position, Memory[PF1], ( Memory[PF1] >> byte(pf1_bit) ) & 0x01, pf1_bit)

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF1] >> byte(pf1_bit) ) & 0x01 == 1 {
							R, G, B := NTSC(Memory[COLUPF])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							// fmt.Println("ENTROU")
							pf1_bit --
						}

						// Reset PF1 bit index for the next line
						if pixel_position == 128 {
							pf1_bit = 7
						}

					// ------------------------------------ PF2 Reflected ------------------------------------ //
					} else if pixel_position <= 160 {

					// Memory[PF2] = 161

					// fmt.Printf("%d\tPF2: %b\t%b\tPF2_BIT: %d\n", pixel_position, Memory[PF2], ( Memory[PF2] >> byte(pf2_bit) ) & 0x01, pf2_bit)

					// If the bit is 1, set the color of the playfield
					if ( Memory[PF2] >> byte(pf2_bit) ) & 0x01 == 1 {
						R, G, B := NTSC(Memory[COLUPF])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					}

					// Each 4 sprites increase the index (playfield bit)
					if pixel_position % 4 == 0 {
						// fmt.Println("ENTROU")
						pf2_bit ++
					}

					// Reset PF1 bit index for the next line
					if pixel_position == 160 {
						pf2_bit = 0
					}
				}

				}


			}















			// TEMPORARY - Random background colors
			// imd.Color = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

			// Draw
			imd.Push(pixel.V( float64(pixel_position - 1)		* width	, (screenHeight * (1 - sizeYused)) + float64(232-line) * height ))
			imd.Push(pixel.V( float64(pixel_position)				* width , (screenHeight * (1 - sizeYused)) + float64(232-line) * height + height))
			imd.Rectangle(0)

			// Count draw operations number per second
			counter_DPS ++
		}

	}

}
