package Graphics

import (
	"fmt"
	"Atari2600/CPU"
)

var (
	// Collision Detection
	// Player 0 - Player 1
	CD_debug				bool	= true
	CD_drawP0_currentline	bool 	= false	// Used to map which line P0 is being drawned
	CD_drawP1_currentline	bool 	= false	// Used to map which line P1 is being drawned
	CD_P0_StartPos			uint8	= 0		// P0 first bit location in the line
	CD_P0_EndPos			uint8	= 0		// P0 last bit location in the line
	CD_P1_StartPos			uint8	= 0		// P1 first bit location in the line
	CD_P1_EndPos			uint8	= 0		// P1 last bit location in the line
	CD_P0_P1_status			bool 	= false	// Set when collision is detected
	CD_P0_sprite_size		uint8	= 0		// P0 size of the sprite after NUSIZ transformation
	CD_P0_sprite_copy		uint32	= 0		// P0 copy of the sprite after NUSIZ transformation
	CD_P1_sprite_size		uint8	= 0		// P1 size of the sprite after NUSIZ transformation
	CD_P1_sprite_copy		uint32	= 0		// P1 copy of the sprite after NUSIZ transformation
	//
	newP0sprite				uint32	= 0		// P0 copy of the sprite after shift
	newP1sprite				uint32	= 0		// P1 copy of the sprite after shift
	//
	collision_test			uint32	= 0		// P0 and P1 AND result
)

func CollisionDetectionP0_P1() {

	// Stop verifying if already detected a collision in this frame
	if !CD_P0_P1_status {

		// STEP 1 - Detect if both players were drawed in the same line
		if CD_drawP0_currentline  {
			if CD_drawP1_currentline {

				if CD_debug {
					fmt.Printf("\nCollision Detection P0-P1 - STEP1 - Both Players Drawed in the same line: %d line\n", line)
				}

				// STEP 2 - Detect witch player is positioned most to the right of the screen
				// P0
				if CD_P0_EndPos >= CD_P1_EndPos {
					fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 0 is on the right:\tP0 start: %d\tP0 end: %d\tP1 pixel start: %d\tP1 end: %d\n", CD_P0_StartPos, CD_P0_EndPos, CD_P1_StartPos, CD_P1_EndPos)

					// STEP 3 - Shift P0 bits to the left to match screen position
					P1_shift := CD_P0_EndPos - CD_P1_EndPos		// Range analyzed from the start of first object until end of the other one

					newP1sprite = CD_P1_sprite_copy << P1_shift

					// Perform an AND between players pixels to detect if they had contact
					collision_test = newP1sprite & CD_P0_sprite_copy	// AND between shifted first object and second object

					// If not zero, there are pixels colliding between Player0 and Player1
					if collision_test != 0 {
						if CD_debug {
							fmt.Printf("Collision Detection P0-P1 - STEP3 - Collision Detected:\n\tP0 size: %d\n\tP1 size: %d\n\tShifted:\t%d\n\tOriginal P1:\t%032b\n\tP0:\t%032b\n\tP1:\t%032b\n\tResult:\t%032b\n",  CD_P0_sprite_size, CD_P1_sprite_size, P1_shift, CD_P1_sprite_copy, CD_P0_sprite_copy, newP1sprite, collision_test)
						}

						// Inform TIA that does not need to check collisions anymore in this frame
						// CD_P0_P1_status = true

						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Temporary Pause do debug
						// CPU.Pause = true
					}
				// P1
				} else {
					fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 1 is on the right:\tP0 start: %d\tP0 end: %d\tP1 pixel start: %d\tP1 end: %d\n", CD_P0_StartPos, CD_P0_EndPos, CD_P1_StartPos, CD_P1_EndPos)

					// STEP 3 - Shift P1 bits to the left to match screen position
					P0_shift := CD_P1_EndPos - CD_P0_EndPos		// Range analyzed from the start of first object until end of the other one

					newP0sprite = CD_P0_sprite_copy << P0_shift

					// Perform an AND between players pixels to detect if they had contact
					collision_test = newP0sprite & CD_P1_sprite_copy	// AND between shifted first object and second object

					// If not zero, there are pixels colliding between Player0 and Player1
					if collision_test != 0 {
						if CD_debug {
							fmt.Printf("Collision Detection P0-P1 - STEP3 - Collision Detected:\n\tP0 size: %d\n\tP1 size: %d\n\tShifted:\t%d\n\tOriginal P0:\t%032b\n\tP0:\t%032b\n\tP1:\t%032b\n\tResult:\t%032b\n",  CD_P0_sprite_size, CD_P1_sprite_size, P0_shift, CD_P0_sprite_copy, newP0sprite, CD_P1_sprite_copy, collision_test)
						}

						// Inform TIA that does not need to check collisions anymore in this frame
						// CD_P0_P1_status = true

						// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
						CPU.MemTIAWrite[CPU.CXPPMM] = 0x80

						// Temporary Pause do debug
						// CPU.Pause = true
					}
				}









				// // STEP 2 - Detect which player is in the left
				// // P0 in the left
				// if CD_P0_StartPos <= CD_P1_StartPos {
				// 	if CD_debug {
				// 		fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 0 is in the left:\tP0 start: %d\tP0 end: %d\tP1 pixel start: %d\tP1 end: %d\n", CD_P0_StartPos, CD_P0_EndPos, CD_P1_StartPos, CD_P1_EndPos)
				// 	}
				//
				// 	fmt.Println()
				//
				// 	// STEP 3 - Detect if the pixels of both players are inside a possible match based on their positions and sizes
				// 	if CD_P0_EndPos >= CD_P1_EndPos {
				// 		if CD_debug {
				// 			fmt.Printf("Collision Detection P0-P1 - STEP3 - Possible colision range, need to analyze the pixel bits\n")
				// 		}
				//
				// 		// STEP 4 - Shift P0 bits to the left to match screen position
				// 		collision_range := CD_P1_EndPos - CD_P0_StartPos		// Range analyzed from the start of first object until end of the other one
				// 		first_obj_shift := collision_range - CD_P1_sprite_size	// Shift left the first object some bits to match with objects position on the screen
				//
				// 		// Perform an AND between players pixels to detect if they had contact
				// 		collision_test = (CD_P0_sprite_copy << first_obj_shift) & CD_P1_sprite_copy	// AND between shifted first object and second object
				//
				// 		// If not zero, there are pixels colliding between Player0 and Player1
				// 		if collision_test != 0 {
				// 			if CD_debug {
				// 				fmt.Printf("Collision Detection P0-P1 - STEP4 - Collision Detected:\n\tP0 size: %d\n\tP1 size: %d\n\tRange:\t\t%d (P1 End: %d - P0 Start: %d)\n\tShifted:\t%d\n\tOriginal P0:\t%020b\n\tP0:\t%020b\n\tP1:\t%020b\n\tResult:\t%020b\n",  CD_P0_sprite_size, CD_P1_sprite_size, collision_range, CD_P1_EndPos, CD_P0_StartPos, first_obj_shift, CD_P0_sprite_copy, CD_P0_sprite_copy << first_obj_shift, CD_P1_sprite_copy, collision_test)
				// 			}
				//
				// 			// Inform TIA that does not need to check collisions anymore in this frame
				// 			CD_P0_P1_status = true
				//
				// 			// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
				// 			CPU.MemTIAWrite[CPU.CXPPMM] = 0x80
				//
				// 			// Temporary Pause do debug
				// 			// CPU.Pause = true
				// 		}
				// 	}
				//
				// // P1 in the left
				// } else {
				// 	if CD_debug {
				// 		fmt.Printf("Collision Detection P0-P1 - STEP2 - Player 1 is in the left:\tP0 start: %d\tP0 end: %d\tP1 pixel start: %d\tP1 end: %d\n", CD_P0_StartPos, CD_P0_EndPos, CD_P1_StartPos, CD_P1_EndPos)
				// 	}
				//
				// 	// STEP 3 - Detect if the pixels of both players are inside a possible match based on their positions and sizes
				// 	if CD_P1_EndPos >= CD_P0_StartPos {
				// 		if CD_debug {
				// 			fmt.Printf("Collision Detection P0-P1 - STEP3 - Possible colision range, need to analyze the pixel bits\n")
				// 		}
				//
				// 		// STEP 4 - Shift P1 bits to the left to match screen position
				// 		collision_range := CD_P0_EndPos - CD_P1_StartPos		// Range analyzed from the start of first object until end of the other one
				// 		first_obj_shift := collision_range - CD_P0_sprite_size	// Shift left the first object some bits to match with objects position on the screen
				//
				// 		// Perform an AND between players pixels to detect if they had contact
				// 		collision_test = (CD_P0_sprite_copy << first_obj_shift) & CD_P1_sprite_copy	// AND between shifted first object and second object
				//
				// 		// If not zero, there are pixels colliding between Player0 and Player1
				// 		if collision_test != 0 {
				// 			// if CD_debug {
				// 				fmt.Printf("Collision Detection P0-P1 - STEP4 - Collision Detected:\n\tRange:\t\t%d (P0 End: %d - P1 Start: %d)\n\tShifted:\t%d\n\tOriginal P1:\t%020b\n\tP0:\t%020b\n\tP1:\t%020b\n\tResult:\t%020b\n", collision_range, CD_P0_EndPos, CD_P1_StartPos, first_obj_shift, CD_P1_sprite_copy, CD_P0_sprite_copy, CD_P1_sprite_copy << first_obj_shift,  collision_test)
				// 			// }
				//
				// 			// Inform TIA that does not need to check collisions anymore in this frame
				// 			CD_P0_P1_status = true
				//
				// 			// Set P0-P1 Collision (TIA READ-ONLY REGISTER CXPPMM: bit 7 - 10000000)
				// 			CPU.MemTIAWrite[CPU.CXPPMM] = 0x80
				//
				// 			// Temporary Pause
				// 			// CPU.Pause = true
				// 		}
				// 	}
				// }
			}
		}
	}

	// Reset variables for the next verification
	CD_drawP0_currentline = false
	CD_drawP1_currentline = false
}
