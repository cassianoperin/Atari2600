package VGS

import (
	"os"
	"fmt"
	// "math/rand"
	"image/color"
	"github.com/faiface/pixel"
)


func TIA_draw() {

	// --------------------------- Draw the 3 sprites of CPU cycle --------------------------- //
	for i := 0 ; i < 3 ; i ++ {

		// Define the pixel position
		pixel_position := ( (int(beamIndex)-1) * 3 ) - 68 + i + 1

		// Dont draw first two sprites outside screen (-2 and -1 X position)
		if pixel_position > 0 {

			// Set the background color as default
			R, G, B := NTSC(Memory[COLUBK])
			imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

			// ----------------------------------------- PF0 ----------------------------------------- //
			if pixel_position <= 16 {

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

					// Update the slice with playfield pixel positions for collision detection
					collision_PF[pixel_position] = 1

				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
					pf0_bit ++
				}

				// Reset PF0 bit index for the next line
				if pixel_position == 16 {
					pf0_bit = 4
				}

			// ----------------------------------------- PF1 ----------------------------------------- //
			} else if pixel_position <= 48 {

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

					// Update the slice with playfield pixel positions for collision detection
					collision_PF[pixel_position] = 1

				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
					pf1_bit --
				}

				// Reset PF1 bit index for the next line
				if pixel_position == 48 {
					pf1_bit = 7
				}

			// ----------------------------------------- PF2 ----------------------------------------- //
			} else if pixel_position <= 80 {

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

					// Update the slice with playfield pixel positions for collision detection
					collision_PF[pixel_position] = 1

				}

				// Each 4 sprites increase the index (playfield bit)
				if pixel_position % 4 == 0 {
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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							pf0_bit ++
						}

						// Reset PF0 bit index for the next line
						if pixel_position == 96 {
							pf0_bit = 4
						}

					// --------------------------------- PF1 Reflected Normal -------------------------------- //
					} else if pixel_position <= 128 {

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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							pf1_bit --
						}

						// Reset PF1 bit index for the next line
						if pixel_position == 128 {
							pf1_bit = 7
						}

					// --------------------------------- PF2 Reflected Normal -------------------------------- //
					} else if pixel_position <= 160 {

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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
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

							// Update the slice with playfield pixel positions for collision detection
							collision_PF[pixel_position] = 1

						}

						// Each 4 sprites increase the index (playfield bit)
						if pixel_position % 4 == 0 {
							pf0_mirror_bit --
						}

						// Reset PF0 bit index for the next line
						if pixel_position == 160 {
							pf0_mirror_bit = 7
						}

					}

				}

			}



			// // NUSIZ and REFP Player draw tests
			// Memory[NUSIZ0] = 0x07
			// Memory[NUSIZ1] = 0x04
			// Memory[REFP1] = 1
			// Memory[GRP1] = 8


			// ---------------------------------- Draw Player 0 ----------------------------------- //

			// Check if GRP0 was set and draw the sprite
			if Memory[GRP0] != 0 {

				// Determine the initial position of the player on the line
				P0_base_position := (int(XPositionP0) * 3) - 68 + int(XFinePositionP0)

				// ----------------------------------------------- NUSIZ0 = 0x00 ----------------------------------------------- //
				if Memory[NUSIZ0] == 0x00 {

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 16 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 32 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 16 || pixel_position == P0_base_position + int(P0_bit) + 32 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 64 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - (P0_bit/2) ) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit/2) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) || pixel_position == P0_base_position + int(P0_bit) + 32 || pixel_position == P0_base_position + int(P0_bit) + 64 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P0_base_position + int(P0_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP0] == 0x00 {
							if Memory[GRP0] >> (7 - (P0_bit/4) ) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP0] >> (P0_bit/4) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP0])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P0[pixel_position] = 1
							}
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

				// Determine the initial position of the player on the line
				P1_base_position := (int(XPositionP1) * 3) - 68 + int(XFinePositionP1)

				// ----------------------------------------------- NUSIZ1 = 0x00 ----------------------------------------------- //
				if Memory[NUSIZ1] == 0x00 {

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 16 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 32 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 16 || pixel_position == P1_base_position + int(P1_bit) + 32 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 64 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit/2) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit/2) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) || pixel_position == P1_base_position + int(P1_bit) + 32 || pixel_position == P1_base_position + int(P1_bit) + 64 {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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

					// Check the initial draw position (set by RESP1)
					if pixel_position == P1_base_position + int(P1_bit) {

						// handle the order of the bits (normal or inverted)
						if Memory[REFP1] == 0x00 {
							if Memory[GRP1] >> (7 - P1_bit/4) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
						// If Reflect Player Enabled (REFP1), invert the order of GRPn
						} else {
							if Memory[GRP1] >> (P1_bit/4) & 0x01 == 1 {
								// READ COLUP0 - Set Player color
								R, G, B := NTSC(Memory[COLUP1])
								imd.Color = color.RGBA{uint8(R), uint8(G), uint8(B), 255}

								// Update the slice with Player0 pixel positions for collision detection
								collision_P1[pixel_position] = 1
							}
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


func FinePositioning(HMPX byte) int8 {

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
			fmt.Printf("\n\tInvalid HMP0 %02X!\n\n", HMP0)
			os.Exit(0)
	}

	return value

}


// func invert_bits(value byte) byte {
// 	// fmt.Printf("\n\n%08b", value)
// 	value = (value & 0xF0) >> 4 | (value & 0x0F) << 4;
// 	value = (value & 0xCC) >> 2 | (value & 0x33) << 2;
// 	value = (value & 0xAA) >> 1 | (value & 0x55) << 1;
// 	// fmt.Printf("\n\n%08b", value)
//
// 	return value;
// }
