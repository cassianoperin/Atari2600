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

					// Check D1 status to use color of players in the score
					if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
						// Playfield color
						R, G, B := NTSC(Memory[COLUPF])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					} else {
						// Player 0 Color (Score)
						R, G, B := NTSC(Memory[COLUP0])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					}

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

					// Check D1 status to use color of players in the score
					if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
						// Playfield color
						R, G, B := NTSC(Memory[COLUPF])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					} else {
						// Player 0 Color (Score)
						R, G, B := NTSC(Memory[COLUP0])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					}

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

					// Check D1 status to use color of players in the score
					if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
						// Playfield color
						R, G, B := NTSC(Memory[COLUPF])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					} else {
						// Player 0 Color (Score)
						R, G, B := NTSC(Memory[COLUP0])
						imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
					}

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


			// --------------------------------- Playfield Reflection -------------------------------- //
			if pixel_position > 80 {

				// --------------------------------- PF0 Reflected Normal -------------------------------- //
				if (Memory[CTRLPF] & 0x01) == 0 {

					if pixel_position <= 96 {

						// fmt.Printf("%d\tPF0: %b\t%b\tPF0_BIT: %d\n", pixel_position, Memory[PF0], ( Memory[PF0] >> byte(pf0_bit) ) & 0x01, pf0_bit)

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF0] >> byte(pf0_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

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

					// --------------------------------- PF1 Reflected Normal -------------------------------- //
					} else if pixel_position <= 128 {

						// Memory[PF1] = 161

						// fmt.Printf("%d\tPF1: %b\t%b\tPF1_BIT: %d\n", pixel_position, Memory[PF1], ( Memory[PF1] >> byte(pf1_bit) ) & 0x01, pf1_bit)

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF1] >> byte(pf1_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

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

					// --------------------------------- PF2 Reflected Normal -------------------------------- //
					} else if pixel_position <= 160 {

						// Memory[PF2] = 161

						// fmt.Printf("%d\tPF2: %b\t%b\tPF2_BIT: %d\n", pixel_position, Memory[PF2], ( Memory[PF2] >> byte(pf2_bit) ) & 0x01, pf2_bit)

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF2] >> byte(pf2_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

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

				// -------------------------------- Mirrored Playfield -------------------------------- //

				} else {

					// ------------------------------------ PF2 Mirrored ------------------------------------ //
					if pixel_position <= 112 {

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF2] >> byte(pf2_mirror_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							// fmt.Println("ENTROU")
							pf2_mirror_bit --
						}

						// Reset PF1 bit index for the next line
						if pixel_position == 112 {
							pf2_mirror_bit = 7
						}

					// ------------------------------------ PF1 Mirrored ------------------------------------ //
					} else if pixel_position <= 144 {

						// If the bit is 1, set the color of the playfield
						if ( Memory[PF1] >> byte(pf1_mirror_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							// fmt.Println("ENTROU")
							pf1_mirror_bit ++
						}

						// Reset PF1 bit index for the next line
						if pixel_position == 144 {
							pf1_mirror_bit = 0
						}

					// ------------------------------------ PF0 Mirrored ------------------------------------ //
					} else if pixel_position <= 160 {


						// If the bit is 1, set the color of the playfield
						if ( Memory[PF0] >> byte(pf0_mirror_bit) ) & 0x01 == 1 {

							// Check D1 status to use color of players in the score
							if (Memory[CTRLPF] & 0x02) >> 1 == 0  {
								// Playfield color
								R, G, B := NTSC(Memory[COLUPF])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							} else {
								// Player 1 Color (Score)
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
							}

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							// fmt.Println("ENTROU")
							pf0_mirror_bit --
						}

						// Reset PF0 bit index for the next line
						if pixel_position == 160 {
							pf0_mirror_bit = 7
						}

					}

				}

			}





			// Memory[NUSIZ0] = 0x05




			// ---------------------------------- Draw Player 0 ----------------------------------- //

			// Check if GRP0 was set and draw the sprite
			if Memory[GRP0] != 0 {

				// ----------------------------------------------- NUSIZ0 = 0x00 ----------------------------------------------- //
				if Memory[NUSIZ0] == 0x00 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x01 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x01 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 16 {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x02 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x02 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 32 {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x03 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x03 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 16 || pixel_position == P0_base_position + int(P0_bit) + 32 {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x04 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x04 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 64 {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x05 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x05 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						if Memory[GRP0] >> (7 - (P0_bit/2) ) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 16 {
							P0_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ0 = 0x06 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x06 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 32 || pixel_position == P0_base_position + int(P0_bit) + 64 {

						if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 8 {
							P0_bit = 0
						}
					}
				// ----------------------------------------------- NUSIZ0 = 0x07 ----------------------------------------------- //
				} else if Memory[NUSIZ0] == 0x07 {

					// Determine the initial position of the player on the line
					P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						if Memory[GRP0] >> (7 - (P0_bit/4) ) & 0x01 == 1 {
							// READ COLUP0 - Set Player color
							R, G, B := NTSC(Memory[COLUP0])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P0_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P0_bit == 32 {
							P0_bit = 0
						}
					}
				}


			}



			// ---------------------------------- Draw Player 1 ----------------------------------- //

			// Check if GRP1 was set and draw the sprite
			if Memory[GRP1] != 0 {

				// ----------------------------------------------- NUSIZ1 = 0x00 ----------------------------------------------- //
				if Memory[NUSIZ1] == 0x00 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x01 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x01 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 16 {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x02 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x02 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 32 {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x03 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x03 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 16 || pixel_position == P1_base_position + int(P1_bit) + 32 {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x04 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x04 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 64 {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x05 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x05 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						if Memory[GRP1] >> (7 - (P1_bit/2) ) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 16 {
							P1_bit = 0
						}
					}


				// ----------------------------------------------- NUSIZ1 = 0x06 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x06 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 32 || pixel_position == P1_base_position + int(P1_bit) + 64 {

						if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 8 {
							P1_bit = 0
						}
					}
				// ----------------------------------------------- NUSIZ1 = 0x07 ----------------------------------------------- //
				} else if Memory[NUSIZ1] == 0x07 {

					// Determine the initial position of the player on the line
					P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						if Memory[GRP1] >> (7 - (P1_bit/4) ) & 0x01 == 1 {
							// READ COLUP1 - Set Player color
							R, G, B := NTSC(Memory[COLUP1])
							imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
						}

						// Incremente the bit of the image
						P1_bit ++

						// When finished all sprites (0-7), reset P0 index
						if P1_bit == 32 {
							P1_bit = 0
						}
					}
				}


			}









			// ------------------------------------- DRAW SPRITE ------------------------------------ //

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
