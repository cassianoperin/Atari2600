package Graphics

import (
	"fmt"
	"Atari2600/CPU"
)

var (
	// Collision Detection
	// Player 0 - Player 1
	CD_debug				bool = true
	CD_drawP0_currentline	bool = false	// Used to map which line P0 is being drawned // Negative when not used
	CD_drawP1_currentline	bool = false	// Used to map which line P1 is being drawned // Negative when not used
	CD_P0_StartPos			uint8	= 0		// P0 first bit location in the line
	CD_P0_EndPos			uint8	= 0		// P0 last bit location in the line
	CD_P1_StartPos			uint8	= 0		// P1 first bit location in the line
	CD_P1_EndPos			uint8	= 0		// P1 last bit location in the line
)

func CollisionDetectionP0_P1() {


	// STEP 1 - Detect if both players were drawed in the same line
	if CD_drawP0_currentline  {
		if CD_drawP1_currentline {

			if CD_debug {
				fmt.Printf("\nCollision Detection P0-P1 - STEP1 - Both Players Drawed in the same line\n")
			}

				// STEP 2 - Detect which player is in the left
				// P0 in the left
				if CD_P0_StartPos <= CD_P1_StartPos {
					if CD_debug {
						fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 0 is in the left:\tP0 pixel start: %d\t\tP1 pixel start: %d\n", CD_P0_StartPos, CD_P1_StartPos)
					}

					// STEP 3 - Detect if the pixels of both players are inside a possible match based on their positions and sizes
					if CD_P0_EndPos >= CD_P1_StartPos {
						if CD_debug {
							fmt.Printf("Collision Detection P0-P1 - STEP3 - Possible colision range, need to analyze the pixel bits\n")
						}

						// STEP 4 - Shift P0 bits to the left to match screen position
						collision_range := CD_P1_EndPos - CD_P0_StartPos	// Range analyzed from the start of first object until end of the other one
						first_obj_shift := collision_range - 8				// Shift left the first object some bits to match with objects position on the screen

						// Perform an AND between players pixels to detect if they had contact
						collision_test := uint16(CPU.Memory[CPU.GRP0]) << first_obj_shift & uint16(CPU.Memory[CPU.GRP1])	// AND between shifted first object and second object

						// If not zero, there are pixels colliding between Player0 and Player1
						if collision_test != 0 {
							if CD_debug {
								fmt.Printf("Collision Detection P0-P1 - STEP4 - Collision Detected:\n\tRange:\t\t%d\n\tShifted:\t%d\n\tOriginal P0:\t%020b\n\tP0:\t%020b\n\tP1:\t%020b\n\tResult:\t%020b\n", collision_range, first_obj_shift, CPU.Memory[CPU.GRP0], uint16(CPU.Memory[CPU.GRP0]) << first_obj_shift, CPU.Memory[CPU.GRP1], collision_test)
							}

							CPU.Pause = true
						}
					}


					// SET A VARIABLE TO DO NOT CHECK ANYMORE in THIS FRAAAAAME
					//REMOVER PAUSE
					// NUSIZ




				// P1 in the left
				} else {
					if CD_debug {
						fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 1 is in the left:\tP0 pixel start: %d\t\tP1 pixel start: %d\n", CD_P0_StartPos, CD_P1_StartPos)
					}

					// STEP 3 - Detect if the pixels of both players are inside a possible match based on their positions and sizes
					if CD_P1_EndPos >= CD_P0_StartPos {
						if CD_debug {
							fmt.Printf("Collision Detection P0-P1 - STEP3 - Possible colision range, need to analyze the pixel bits\n")
						}

						// STEP 4 - Shift P1 bits to the left to match screen position
						collision_range := CD_P0_EndPos - CD_P1_StartPos	// Range analyzed from the start of first object until end of the other one
						first_obj_shift := collision_range - 8				// Shift left the first object some bits to match with objects position on the screen

						// Perform an AND between players pixels to detect if they had contact
						collision_test := uint16(CPU.Memory[CPU.GRP1]) << first_obj_shift & uint16(CPU.Memory[CPU.GRP0])	// AND between shifted first object and second object

						// If not zero, there are pixels colliding between Player0 and Player1
						if collision_test != 0 {
							fmt.Printf("Collision Detection P0-P1 - STEP4 - Collision Detected:\n\tRange:\t\t%d\n\tShifted:\t%d\n\tOriginal P1:\t%020b\n\tP0:\t%020b\n\tP1:\t%020b\n\tResult:\t%020b\n", collision_range, first_obj_shift, CPU.Memory[CPU.GRP1], CPU.Memory[CPU.GRP0], uint16(CPU.Memory[CPU.GRP1]) << first_obj_shift,  collision_test)

							CPU.Pause = true
						}

					}
				}

				// Set P0-P1 Collision (bit 7 - 10000000)
				// CONFLICTING WITH P1 COLOR 0x07, NEED TO IMPLEMENT TIA READ ONLY REGISTERS
				// CPU.Memory[CPU.CXPPMM] = 0x80

		}

	}

	// CLEAN COLLISION DETECTION VARIABLES
	CD_drawP0_currentline = false
	CD_drawP1_currentline = false

}
