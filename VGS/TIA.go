package VGS

import (
	"os"
	"fmt"
	// "golang.org/x/image/colornames"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)


func TIA(action int8, win_2nd_level *pixelgl.Window) {

	// Don't draw outside visible area
	if line > 40 && line <= 232 {
	// if line > 40 && line <= 150 {

		// Don't draw in horizontal blank area
		if beamIndex * 3 > 68 {
			drawBackground()
		}

	}


	switch action {
		// --------------------------------------- WSYNC --------------------------------------- //
		// Halt CPU until next scanline starts
		// Skip to the next scanline
		case int8(WSYNC): //0x02
			if debugGraphics {
				fmt.Printf("\tLine: %d\tWSYNC SET (Beam index: %d)\n", line, beamIndex)
			}

			// Disable CPU
			CPU_Enabled = false

			// Increment beam index until the end of the line to re-enable CPU
			beamIndex ++


		// --------------------------------------- VBLANK --------------------------------------- //
		case int8(VBLANK): //0x01

			// Enable VBLANK
			if Memory[VBLANK] == 0x02 {
				if debugGraphics {
					fmt.Printf("\tVBLANK Enabled\n")
				}
			} else if Memory[VBLANK] == 0x00 {
				if debugGraphics {
					fmt.Printf("\tVBLANK Disabled\n")
				}
			} else {
				if debugGraphics {
					fmt.Printf("\tVBLANK VALUE !=0 !=2 exiting\t%d\n", Memory[VBLANK])
				}
			}

		// --------------------------------------- VSYNC --------------------------------------- //
		case int8(VSYNC): //0x00

			// Enable VSYNC
			if Memory[VSYNC] == 0x02 {
				if debugGraphics {
					fmt.Printf("\tVSYNC Enabled\n")
				}
			} else if Memory[VSYNC] == 0x00 {
				if debugGraphics {
					fmt.Printf("\tVSYNC Disabled\n")
				}
			} else {
				if debugGraphics {
					fmt.Printf("\tVSYNC VALUE !=0 !=2 exiting\t%d\n",Memory[VSYNC] )
				}
				os.Exit(0)
			}

		case int8(COLUBK): //0x09
			if debugGraphics {
				fmt.Printf("\tCOLUBK SET! Beam index: %d\n", beamIndex)
			}

		case int8(GRP0): //0x1B
			if debugGraphics {
				fmt.Printf("\tCycle: %d\tGRP0 SET\t%b\n", counter_F_Cycle, Memory[GRP0])
			}

		case int8(GRP1): //0x1C
			if debugGraphics {
				fmt.Printf("\tCycle: %d\tGRP1 SET\t%b\n", counter_F_Cycle, Memory[GRP1])
			}

		case int8(RESP0): //0x1B
			if debugGraphics {
				fmt.Printf("\t%d - RESP0 SET - DRAW P0 SPRITE!\tBeam: %d\tP0: %b\n", counter_F_Cycle, beamIndex, Memory[GRP0])
			}
			XPositionP0 = beamIndex

		case int8(RESP1): //0x1C
			if debugGraphics {
				fmt.Printf("\tRESP1 SET - DRAW P1 SPRITE!\tBeam: %d\n", beamIndex)
			}
			XPositionP1 = beamIndex

		case int8(HMP0): //0x20
			if debugGraphics {
				fmt.Printf("\tHMP0 SET - Define P0 Fine Positioning\n")
			}
			XFinePositionP0 = FinePositioning(Memory[HMP0])

		case int8(HMP1): //0x21
			if debugGraphics {
				fmt.Printf("\tHMP1 SET - Define P1 Fine Positioning\n")
			}
			XFinePositionP1 = FinePositioning(Memory[HMP1])

		case int8(CXCLR): //0x2C
			if debugGraphics {
				fmt.Printf("\tCXCLR SET - Clear Collision Latches\n")
			}
			MemTIAWrite[CXPPMM] = 0x00
			MemTIAWrite[CXP0FB] = 0x00

		default:

	}

	// When finished drawing the LINE, reset Beamer and start a new LINE
	// Needed for colorbg demo
	if beamIndex > 76 {
		newLine(win_2nd_level)
		// Pause = true
	}
	//
	// // Draw messages into the screen
	// if ShowMessage {
	// 	// textMessage.Clear()
	// 	// fmt.Fprintf(textMessage, TextMessageStr)
	// 	// textMessage.Draw(win, pixel.IM.Scaled(textMessage.Orig, 1))
	// }


}


// Every end of line check vor VSYNC and VBLANK to sync with CRT
func check_VSYNC_VBLANK(win_2nd_level *pixelgl.Window) {

	// Applications that doesn't handle correctly VSYNC, if line > 262, force newFrame
	if line > 262 {
		newFrame(win_2nd_level)

	// If line <= 262, handle normally
	} else {

			// Test if in Vertical Blank (do not draw anything)
			if Memory[VBLANK] == 2 {

				// During Vertical Blank, if vsync is set
				if  Memory[VSYNC] == 2  {
					newFrame(win_2nd_level)
					// win_2nd_level.Clear(colornames.Red)
					// win_2nd_level.Update()

					// When VSYNC is set, CPU inform CRT to start a new frame
					// 3 lines VSYNC

					// ENABLE VSYNC
					if Memory[VSYNC] == 0x02 {

						if Memory[VBLANK] == 2 {
							if debugGraphics {
								fmt.Printf("\tLine: %d\tCRT - VSYNC\n\n", line)
							}
						} else {
							if debugGraphics {
								fmt.Printf("\tLine: %d\tCRT - VSYNC without VBLANK - Not correct!!!\n\n", line)
							}
						}

					// DISABLE VSYNC
					} else if Memory[VSYNC] == 0x00 {
						if debugGraphics {
							fmt.Printf("\tCRT - VSYNC DISABLED\n")
						}

					} else {
						fmt.Printf("\tCRT - VSYNC VALUE NOT 0 or 2! Exiting!\n")
						os.Exit(2)
					}

				// 37 lines VBLANK
				} else if Memory[VBLANK] == 2 {
					if debugGraphics {
						fmt.Printf("\tLine: %d\tVBLANK\t\t(vblank: %02X\tvsync: %02X)\n\n", line,Memory[VBLANK], Memory[VSYNC])
					}
					// visibleArea = false // Inform that finished visible lines

				}

			// VBLANK turned OFF, start drawing the 192 lines of visible Area
			} else {
				// visibleArea = true // Inform that reached visible lines






				// // DRAW PLAYER 0
				if Memory[GRP0] != 0 {
					// fmt.Printf("Cycle: %d - DRAW P0\n", Cycle)
					// drawPlayer(0, win_2nd_level)
					// P0_draw_line = 232 - line
					// fmt.Println(P0_draw_line)

				}

				// // DRAW PLAYER 1
				if Memory[GRP1] != 0 {
					// drawPlayer(1, win_2nd_level)
				}

			}

		}


}


func newLine(win_2nd_level *pixelgl.Window) {

	if debugGraphics {
		fmt.Printf("Finished the line %d, starting a new one. Beam: %d\n", line, beamIndex)
	}
	// beamIndex = beamIndex - 76
	beamIndex = 0
	line ++

	CPU_Enabled = true
	// // Reset to default value
	// TIA_Update = -1
	check_VSYNC_VBLANK(win_2nd_level)

	// Reset Collision Detection Line Array


	// // Print Player 0
	// for i := 0 ; i < len(collision_P0) ; i++ {
	// 	if collision_P0[i] == 1 {
	// 		fmt.Printf("#")
	// 	} else {
	// 		fmt.Printf(" ")
	// 	}
	// }
	// fmt.Printf("\n")

	// // Print Playfield
	// for i := 0 ; i < len(collision_PF) ; i++ {
	// 	if collision_PF[i] == 1 {
	// 		fmt.Printf("#")
	// 	} else {
	// 		fmt.Printf(" ")
	// 	}
	// }
	// fmt.Printf("\n")


	// COLLISION DETECTION
	for i := 1 ; i <= 160 ; i++ {

		// CXP0FB (D7) - P0-PF
		if collision_P0[i] == 1 {
			if collision_PF[i] == 1 {
				MemTIAWrite[CXP0FB] = 0x80
			}
		}

		// CXPPMM (D7) - P0-P1
		if collision_P0[i] == 1 {
			if collision_P1[i] == 1 {
				MemTIAWrite[CXPPMM] = 0x80
			}
		}

		// CXP1FB (D7) - P1-PF
		if collision_P1[i] == 1 {
			if collision_PF[i] == 1 {
				MemTIAWrite[CXP1FB] = 0x80
			}
		}



	}



	// After processing, clean collision detection slices
	collision_PF = [161]byte{}
	collision_P0 = [161]byte{}
	collision_P1 = [161]byte{}

}






// When finished drawing the screen, reset and start a new frame
func newFrame(win_3nd_level *pixelgl.Window) {

	// Start a new frame on first VSYNC
	if counter_VSYNC == 1 {

		// Reset line counter
		line = 1

		// Increment frames
		counter_FPS ++
		// Reset Frame Cycle counter
		counter_F_Cycle = 0
		// Increment Frame Counter
		counter_Frame ++

		// Reset Controllers Buttons to 1 (not pressed)
		Memory[SWCHA] = 0xFF //1111 11111



		if debugGraphics {
			fmt.Printf("\nFinished the screen height, start a new frame (%d).\n", counter_Frame)
		}

		// Update control to just do it on first occurence
		counter_VSYNC ++

		// After finishing a frame, draw it to screen and refresh
		imd.Draw(win_3nd_level)
		win_3nd_level.Update()

		// Clean the current draws for next frame
		imd	= imdraw.New(nil)

	// Reset counter for next frame
	} else if counter_VSYNC == 2 {
		counter_VSYNC ++
	} else if counter_VSYNC == 3 {
		counter_VSYNC = 1
	}


}
