package Graphics

import (
	"os"
	"fmt"
	"Atari2600/CPU"
)

var (

	// Line draw control
	line				int = 1
	line_max			int = 262

	// FPS count
	frames			= 0
	draws			= 0

	// Workaround to avoid  WSYNC before VSYNC
	VSYNC_passed		bool = false
)


func TIA(action int8) {

	// TODO
	// Just draw in visible Area
	// if visibleArea {

		drawBackground()

		// if line ==40 {
		// 	CPU.Pause = true
		//
		// }
	// }


	switch action {
		// --------------------------------------- WSYNC --------------------------------------- //
		// Halt CPU until next scanline starts
		// Skip to the next scanline
		case int8(CPU.WSYNC): //0x02
			if debug {
				fmt.Printf("\tCRT - WSYNC SET\n")
			}

			// Test if in Vertical Blank (do not draw anything)
			if CPU.Memory[CPU.VBLANK] == 2 {
				// os.Exit(2)


				// During Vertical Blank, if vsync is set
				if  CPU.Memory[CPU.VSYNC] == 2  {

					VSYNC_passed = true	// Used to control WSYNCS before VSYNC

					// When VSYNC is set, CPU inform CRT to start a new frame
					// 3 lines VSYNC

					// ENABLE VSYNC
					if CPU.Memory[CPU.VSYNC] == 0x02 {

						if CPU.Memory[CPU.VBLANK] == 2 {
							if debug {
								fmt.Printf("\tLine: %d\tCRT - VSYNC\n\n", line)
							}
						} else {
							if debug {
								fmt.Printf("\tLine: %d\tCRT - VSYNC without VBLANK - Not correct!!!\n\n", line)
							}
						}

					// DISABLE VSYNC
					} else if CPU.Memory[CPU.VSYNC] == 0x00 {
						if debug {
							fmt.Printf("\tCRT - VSYNC DISABLED\n")
						}

					} else {
						fmt.Printf("\tCRT - VSYNC VALUE NOT 0 or 2! Exiting!\n")
						os.Exit(2)
					}

				// 37 lines VBLANK
				} else if CPU.Memory[CPU.VBLANK] == 2 {
					if debug {
						fmt.Printf("\tLine: %d\tVBLANK\t\t(vblank: %02X\tvsync: %02X)\n\n", line,CPU.Memory[CPU.VBLANK], CPU.Memory[CPU.VSYNC])
					}
					visibleArea = false // Inform that finished visible lines

				}

			// VBLANK turned OFF, start drawing the 192 lines of visible Area
			} else {
				visibleArea = true // Inform that reached visible lines

				// Finish drawing line (X=228) 76x3
				CPU.Beam_index = 76
				if debug {
					fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, CPU.Beam_index)
				}
				drawBackground()



				if debug {
					fmt.Printf("\tLine: %d\tVisible Area: %d\n\n", line, line-40)
				}

				// Draw the entire line of Playfield
				draw_playfield()


				// // DRAW PLAYER 0
				if CPU.Memory[CPU.GRP0] != 0 {
					// fmt.Printf("Cycle: %d - DRAW P0\n", CPU.Cycle)
					drawPlayer0()
				}

				// // DRAW PLAYER 1
				if CPU.Memory[CPU.GRP1] != 0 {
					drawPlayer1()
				}

				// COLLISION DETECTION
				// P0 - P1
				CollisionDetectionP0_P1()
			}

			// Reset the beam index
			CPU.Beam_index = 0
			old_BeamIndex = 0
			line ++



		// --------------------------------------- VBLANK --------------------------------------- //
		case int8(CPU.VBLANK): //0x01

			// Enable VBLANK
			if CPU.Memory[CPU.VBLANK] == 0x02 {
				if debug {
					fmt.Printf("\tVBLANK Enabled\n")
				}
			} else if CPU.Memory[CPU.VBLANK] == 0x00 {
				if debug {
					fmt.Printf("\tVBLANK Disabled\n")
				}
			} else {
				if debug {
					fmt.Printf("\tVBLANK VALUE !=0 !=2 exiting\t%d\n", CPU.Memory[CPU.VBLANK])
				}
			}

		// --------------------------------------- VSYNC --------------------------------------- //
		case int8(CPU.VSYNC): //0x00

			// Enable VSYNC
			if CPU.Memory[CPU.VSYNC] == 0x02 {
				if debug {
					fmt.Printf("\tVSYNC Enabled\n")
				}
			} else if CPU.Memory[CPU.VSYNC] == 0x00 {
				if debug {
					fmt.Printf("\tVSYNC Disabled\n")
				}
			} else {
				if debug {
					fmt.Printf("\tVSYNC VALUE !=0 !=2 exiting\t%d\n",CPU.Memory[CPU.VSYNC] )
				}
				os.Exit(0)
			}

		case int8(CPU.COLUBK): //0x09
			if debug {
				fmt.Printf("\tCOLUBK SET! Beam index: %d\n", CPU.Beam_index)
			}

		case int8(CPU.GRP0): //0x1B
			if debug {
				fmt.Printf("\tGRP0 SET\n")
			}

		case int8(CPU.GRP1): //0x1C
			if debug {
				fmt.Printf("\tGRP1 SET\n")
			}

		case int8(CPU.RESP0): //0x1B
			if debug {
				fmt.Printf("\tRESP0 SET - DRAW P0 SPRITE!\tBeam: %d\n", CPU.Beam_index)
			}
			XPositionP0 = CPU.Beam_index
			// drawPlayer0()


		case int8(CPU.RESP1): //0x1C
			if debug {
				fmt.Printf("\tRESP1 SET - DRAW P1 SPRITE!\tBeam: %d\n", CPU.Beam_index)
			}
			XPositionP1 = CPU.Beam_index
			// drawPlayer1()

		case int8(CPU.HMP0): //0x20
			if debug {
				fmt.Printf("\tHMP0 SET - Define P0 Fine Positioning\n")
			}
			XFinePositionP0 = Fine(CPU.Memory[CPU.HMP0])

		case int8(CPU.HMP1): //0x21
			if debug {
				fmt.Printf("\tHMP1 SET - Define P1 Fine Positioning\n")
			}
			XFinePositionP1 = Fine(CPU.Memory[CPU.HMP1])

		case int8(CPU.CXCLR): //0x2C
			if debug {
				fmt.Printf("\tCXCLR SET - Clear Collision Latches\n")
			}
			CPU.MemTIAWrite[CPU.CXPPMM] = 0x00

		default:

	}


	// When finished drawing the screen, reset and start a new frame
	if line == line_max + 1 {
		if debug {
			fmt.Printf("\nFinished the screen height, start a new frame.\n")
		}
		// Reset line counter
		line = 1
		// Workaround for WSYNC before VSYNC
		VSYNC_passed = false

		// Update Collision Detection Flags
		CD_P0_P1_status = false				// Informm TIA to start looking for collisions again

		// Increment frames
		frames ++
	}

	// When finished drawing the LINE, reset Beamer and start a new LINE
	// Needed for colorbg demo
	// DISABLED because its causing empty lines in the begin
	// if CPU.Beam_index > 76 {
	// 	// if debug {
	// 		fmt.Printf("\nFinished the line, starting a new one.\n")
	// 	// }
	// 	CPU.Beam_index = 0
	// 	old_BeamIndex = 0
	// 	line ++
	// }

	// Reset to default value
	CPU.TIA_Update = -1
}
