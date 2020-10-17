package Graphics

import (
	"os"
	"fmt"
	"time"
	"Atari2600/CPU"
	"Atari2600/Global"
	"github.com/faiface/pixel"
)

var (

	// Line draw control
	line				int = 1
	line_max			int = 262

	// FPS count
	frames			= 0
	draws			= 0
)


func TIA(action int8) {

	// Time measurement - TIA Cycle
	if CPU.DebugTiming {
		CPU.StartTIA = time.Now()
	}


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
					line = 1
					Global.VSYNC_passed = true	// Used to control WSYNCS before VSYNC

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
					// fmt.Printf("Old BeamIndex: %d\t New BeamIndex: %d\n", old_BeamIndex, CPU.Beam_index)
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
					drawPlayer(0)
				}

				// // DRAW PLAYER 1
				if CPU.Memory[CPU.GRP1] != 0 {
					drawPlayer(1)
				}

				// COLLISION DETECTION
				// P0 - P1
				// CollisionDetectionP0_P1()
			}

			// Reset the beam index
			CPU.Beam_index = 0
			old_BeamIndex = 0
			// Reset Collision Detection Line Array
			CD_P0_P1 = [160]byte{}
			CD_P0_PF = [160]byte{}

			// Increment Line
			// CPU.Pause = true
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
					line = 1
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
			CPU.MemTIAWrite[CPU.CXP0FB] = 0x00

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
		Global.VSYNC_passed = false

		// Update Collision Detection Flags
		CD_P0_P1_collision_detected = false		// Informm TIA to start looking for collisions again
		CD_P0_PF_collision_detected = false		// Informm TIA to start looking for collisions again

		// Increment frames
		frames ++
	}

	// When finished drawing the LINE, reset Beamer and start a new LINE
	// Needed for colorbg demo
	// DISABLED because its causing empty lines in the begin
	// if CPU.Beam_index > 76 {
	// 	// if debug {
	// 		fmt.Printf("\nFinished the line, starting a new one.\n")
	// 		// CPU.Pause = true
	// 	// }
	// 	CPU.Beam_index = 0
	// 	old_BeamIndex = 0
	// 	line ++
	// }

	// Reset to default value
	CPU.TIA_Update = -1

	// Time measurement - TIA Cycle
	if CPU.DebugTiming {
		elapsedTIA := time.Since(CPU.StartTIA)
		if elapsedTIA.Seconds() > CPU.DebugTimingLimit {
			fmt.Printf("\tOpcode: %X\tEntire TIA Cycle took %f seconds\n", CPU.Opcode, elapsedTIA.Seconds())
			// CPU.Pause = true
		}
	}

	// Draw messages into the screen
	if Global.ShowMessage {
		textMessage.Clear()
		fmt.Fprintf(textMessage, Global.TextMessageStr)
		textMessage.Draw(Global.Win, pixel.IM.Scaled(textMessage.Orig, 1))
	}


}
