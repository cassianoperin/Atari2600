package VGS

import (
	"os"
	"fmt"
	// "golang.org/x/image/colornames"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)


func TIA(action int16, win_2nd_level *pixelgl.Window) {

	// Don't draw outside visible area
	if line > 40 && line <= 232 {

		// Don't draw in horizontal blank area
		if beamIndex * 3 > 68 {
			TIA_draw()
		}

	}


	switch action {
		// --------------------------------------- WSYNC --------------------------------------- //
		// Halt CPU until next scanline starts and skip to the next scanline
		case int16(WSYNC): //0x02
			if debugGraphics {
				fmt.Printf("\tLine: %d\tWSYNC SET (Beam index: %d)\n", line, beamIndex)
			}

			// Disable CPU
			CPU_Enabled = false

			// Increment beam index until the end of the line to re-enable CPU
			beamIndex ++


		// --------------------------------------- VBLANK --------------------------------------- //
		case int16(VBLANK): //0x01

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
		case int16(VSYNC): //0x00

			// Enable VSYNC
			if Memory[VSYNC] == 0x02 {
				if debugGraphics {
					fmt.Printf("\tVSYNC Enabled\n")
				}
			} else if Memory[VSYNC] == 0x00 {
				if debugGraphics {
					fmt.Printf("\tVSYNC Disabled\n")
				}
			// NOT MAPPED WHY 42 (2A) YET!!! surroundings game
			} else if Memory[VSYNC] == 0x2A {
				if debugGraphics {
					fmt.Printf("\tVSYNC Enabled (VSYNC = 42) !!!!\n")
				}
			} else {
				fmt.Printf("\tVSYNC VALUE !=0 !=2 !=42 exiting\t%d\n",Memory[VSYNC] )
				os.Exit(2)
			}

		case int16(COLUBK): //0x09
			if debugGraphics {
				fmt.Printf("\tCOLUBK SET! Beam index: %d\n", beamIndex)
			}

		case int16(GRP0): //0x1B
			if debugGraphics {
				fmt.Printf("\tCycle: %d\tGRP0 SET\tGRP0: %b\n", counter_F_Cycle, Memory[GRP0])
			}

		case int16(GRP1): //0x1C
			if debugGraphics {
				fmt.Printf("\tCycle: %d\tGRP1 SET\tGRP1: %b\n", counter_F_Cycle, Memory[GRP1])
			}

		case int16(RESP0): //0x1B
			if debugGraphics {
				fmt.Printf("\t%d - RESP0 SET - DRAW P0 SPRITE!\tBeam: %d\tGRP0: %b\n", counter_F_Cycle, beamIndex, Memory[GRP0])
			}
			XPositionP0 = beamIndex

		case int16(RESP1): //0x11
			if debugGraphics {
				fmt.Printf("\t%d - RESP1 SET - DRAW P1 SPRITE!\tBeam: %d\tGRP1: %b\n", counter_F_Cycle, beamIndex, Memory[GRP1])
			}
			XPositionP1 = beamIndex

		case int16(HMP0): //0x20
			if debugGraphics {
				fmt.Printf("\tHMP0 SET - Define P0 Fine Positioning\n")
			}
			// XFinePositionP0 = FinePositioning(Memory[HMP0])

		case int16(HMP1): //0x21
			if debugGraphics {
				fmt.Printf("\tHMP1 SET - Define P1 Fine Positioning\n")
			}
			// XFinePositionP1 = FinePositioning(Memory[HMP1])

		case int16(HMCLR): //0x2B
			if debugGraphics {
				fmt.Printf("\tHMCLR SET - Clear Horizontal Move Registers\n")
			}
			Memory[HMP0] = 0x00
			Memory[HMP1] = 0x00
			Memory[HMM0] = 0x00
			Memory[HMM1] = 0x00
			Memory[HMBL] = 0x00

		case int16(HMOVE): //0x2A
			if debugGraphics {
				fmt.Printf("\tHMOVE SET - Apply Horizontal Motion\n")
			}
			// Check if will be necessary to keep the HMP values in some cache
			XFinePositionP0 = FinePositioning(Memory[HMP0])
			XFinePositionP1 = FinePositioning(Memory[HMP1])

		case int16(CXCLR): //0x2C
			if debugGraphics {
				fmt.Printf("\tCXCLR SET - Clear Collision Latches\n")
			}
			Memory_TIA_RO[CXPPMM] = 0x00
			Memory_TIA_RO[CXP0FB] = 0x00
			Memory_TIA_RO[CXP1FB] = 0x00


		// --------------------------- RIOT --------------------------- //

		case int16(TIM1T): //0x294
			if debugRIOT {
				fmt.Printf("\tTIM1T (Write) - Set 1 Cycle Timer\n")
			}
			fmt.Printf("\tTIM1T (Write) - Proposital Exit to map usage!\n")
			os.Exit(2)

		case int16(TIM8T): //0x295
			if debugRIOT {
				fmt.Printf("\tTIM8T (Write) - Set 8 Cycle Timer\n")
			}
			// os.Exit(2)

		case int16(TIM64T): //0x296
			if debugRIOT {
				fmt.Printf("\tTIM64T (Write) - Set 64 Cycle Timer\n")
			}
			// os.Exit(2)

		case int16(T1024T): //0x297
			if debugRIOT {
				fmt.Printf("\tT1024T (Write) - Set 1024 Cycle Timer\n")
			}
			fmt.Printf("\tTIM1T (Write) - Proposital Exit to map usage!\n")
			os.Exit(2)
		default:

	}

	// When finished drawing the LINE, reset Beamer and start a new LINE
	if beamIndex > 76 {
		newLine(win_2nd_level)
	}

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
				// win_2nd_level.Clear(colornames.Black)
				// win_2nd_level.Update()

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
			}

		// VBLANK turned OFF, start drawing the 192 lines of visible Area
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


	// // Print Collision Player 0 Slice
	// for i := 0 ; i < len(collision_P0) ; i++ {
	// 	if collision_P0[i] == 1 {
	// 		fmt.Printf("#")
	// 	} else {
	// 		fmt.Printf(" ")
	// 	}
	// }
	// fmt.Printf("\n")

	// // Print Collision Player 1 Slice
	// for i := 0 ; i < len(collision_P1) ; i++ {
	// 	if collision_P1[i] == 1 {
	// 		fmt.Printf("#")
	// 	} else {
	// 		fmt.Printf(" ")
	// 	}
	// }
	// fmt.Printf("\n")

	// // Print Collision Playfield Slice
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
				// Memory_TIA_RO[CXP0FB] = 0x80
				update_Memory_TIA_RO(CXP0FB, 0x80)
			}
		}

		// CXPPMM (D7) - P0-P1
		if collision_P0[i] == 1 {
			if collision_P1[i] == 1 {
				// Memory_TIA_RO[CXPPMM] = 0x80
				update_Memory_TIA_RO(CXPPMM, 0x80)
			}
		}

		// CXP1FB (D7) - P1-PF
		if collision_P1[i] == 1 {
			if collision_PF[i] == 1 {
				// Memory_TIA_RO[CXP1FB] = 0x80
				update_Memory_TIA_RO(CXP1FB, 0x80)
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
		// Memory[INPT0] = 0xFF //1111 11111
		// Memory[INPT1] = 0xFF //1111 11111
		// Memory[INPT2] = 0xFF //1111 11111
		// Memory[INPT3] = 0xFF //1111 11111
		update_Memory_TIA_RO(INPT4, 0xFF) //1111 11111
		update_Memory_TIA_RO(INPT5, 0xFF) //1111 11111


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
