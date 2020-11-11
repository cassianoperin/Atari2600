package VGS

import (
	"os"
	"fmt"
	"time"
	"github.com/faiface/pixel"
)


func TIA(action int8) {

	// Time measurement - TIA Cycle
	if debugTiming {
		debugTiming_StartTIA = time.Now()
	}


	// TODO
	// Just draw in visible Area
	// if visibleArea {

		drawBackground()

		// if line ==40 {
		// 	Pause = true
		//
		// }
	// }


	switch action {
		// --------------------------------------- WSYNC --------------------------------------- //
		// Halt CPU until next scanline starts
		// Skip to the next scanline
		case int8(WSYNC): //0x02
			if debugGraphics {
				fmt.Printf("\tCRT - WSYNC SET\n")
			}

			// Test if in Vertical Blank (do not draw anything)
			if Memory[VBLANK] == 2 {
				// os.Exit(2)


				// During Vertical Blank, if vsync is set
				if  Memory[VSYNC] == 2  {
					newFrame()
					VSYNC_passed = true	// Used to control WSYNCS before VSYNC

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
					visibleArea = false // Inform that finished visible lines

				}

			// VBLANK turned OFF, start drawing the 192 lines of visible Area
			} else {
				visibleArea = true // Inform that reached visible lines

				// Finish drawing line (X=228) 76x3
				beamIndex = 76
				if debugGraphics {
					// fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, Beam_index)
				}
				drawBackground()



				if debugGraphics {
					fmt.Printf("\tLine: %d\tVisible Area: %d\n\n", line, line-40)
				}

				// Draw the entire line of Playfield
				draw_playfield()


				// // DRAW PLAYER 0
				if Memory[GRP0] != 0 {
					// fmt.Printf("Cycle: %d - DRAW P0\n", Cycle)
					drawPlayer(0)
				}

				// // DRAW PLAYER 1
				if Memory[GRP1] != 0 {
					drawPlayer(1)
				}

				// COLLISION DETECTION
				// P0 - P1
				// CollisionDetectionP0_P1()
			}

			// Reset the beam index
			beamIndex = 0
			old_beamIndex = 0
			// Reset Collision Detection Line Array
			CD_P0_P1 = [160]byte{}
			CD_P0_PF = [160]byte{}

			// Increment Line
			// Pause = true
			line ++



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
					newFrame()
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
				fmt.Printf("\tGRP0 SET\n")
			}

		case int8(GRP1): //0x1C
			if debugGraphics {
				fmt.Printf("\tGRP1 SET\n")
			}

		case int8(RESP0): //0x1B
			if debugGraphics {
				fmt.Printf("\tRESP0 SET - DRAW P0 SPRITE!\tBeam: %d\n", beamIndex)
			}
			XPositionP0 = beamIndex
			// drawPlayer0()


		case int8(RESP1): //0x1C
			if debugGraphics {
				fmt.Printf("\tRESP1 SET - DRAW P1 SPRITE!\tBeam: %d\n", beamIndex)
			}
			XPositionP1 = beamIndex
			// drawPlayer1()

		case int8(HMP0): //0x20
			if debugGraphics {
				fmt.Printf("\tHMP0 SET - Define P0 Fine Positioning\n")
			}
			XFinePositionP0 = Fine(Memory[HMP0])

		case int8(HMP1): //0x21
			if debugGraphics {
				fmt.Printf("\tHMP1 SET - Define P1 Fine Positioning\n")
			}
			XFinePositionP1 = Fine(Memory[HMP1])

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
	// DISABLED because its causing empty lines in the begin
	if beamIndex > 76 {
		if debugGraphics {
			fmt.Printf("\nFinished the line, starting a new one.\n")
		}
		beamIndex = beamIndex - 76
		old_beamIndex = 0
		line ++
	}

	// Reset to default value
	TIA_Update = -1

	// Time measurement - TIA Cycle
	if debugTiming {
		elapsedTIA := time.Since(debugTiming_StartTIA)
		if elapsedTIA.Seconds() > debugTiming_Limit {
			fmt.Printf("\tOpcode: %X\tEntire TIA Cycle took %f seconds\n", opcode, elapsedTIA.Seconds())
			// Pause = true
		}
	}

	// Draw messages into the screen
	if ShowMessage {
		textMessage.Clear()
		fmt.Fprintf(textMessage, TextMessageStr)
		textMessage.Draw(win, pixel.IM.Scaled(textMessage.Orig, 1))
	}


}

// When finished drawing the screen, reset and start a new frame
func newFrame() {
	if debugGraphics {
		fmt.Printf("\nFinished the screen height, start a new frame.\n")
	}
	// Reset line counter
	line = 1
	// Workaround for WSYNC before VSYNC
	VSYNC_passed = false

	// Update Collision Detection Flags
	CD_P0_P1_collision_detected = false		// Informm TIA to start looking for collisions again
	CD_P0_PF_collision_detected = false		// Informm TIA to start looking for collisions again

	// Increment frames
	counter_FPS ++
	// Reset Frame Cycle counter
	counter_F_Cycle = 0
	// Increment Frame Counter
	counter_Frame ++
}
