package Graphics

import (
	// "os"
	// "fmt"
	// "Atari2600/CPU"
)







// TIA monitor changes in Memory addresses 0 to $7F (0..127)
// func TIA() {
	// var i byte
	//
	// for i = 0 ; i < byte(len(TIA_Memory)) ; i++ {
	// 	if TIA_Memory[i] != CPU.Memory[i] {
	// 		// fmt.Printf("\tTIA[%d] = %d\tMemory[%d] = %d\t DIFFERENT\n", i, TIA_Memory[i], i, CPU.Memory[i])
	//
	// 		// Update TIA
	// 		TIA_Memory[i] = CPU.Memory[i]
	//
	// 		// fmt.Println(CRT_action)
	// 		// CRT(i)
	// 		// finish the loop
	// 		// break
	// 	}
	// }


	//
	// // --------------------------------------- COLUBK --------------------------------------- //
	// // Halt CPU until next scanline starts
	// // Skip to the next scanline
	// if memAddr == uint16(COLUBK) {
	// 	if Debug {
	// 		fmt.Printf("\nTIA - COLUBK SET\n")
	// 	}
	//
	// 	DrawLine = true		// Tell DrawGraphics to draw a line
	// 	Beam_index = 0		// Reset the beam index
	// 	WSYNC_flag = true	// Tells CRT that WSYNC is set
	//
	// }



// 	if memAddr == uint16(WSYNC) {
// 		if Debug {
// 			fmt.Printf("\nWSYNC SET\n")
// 		}
// 		DrawLine = true
// 		Beam_index = 0
//
// 		if Memory[GRP0] != 0 {
// 			if Debug {
// 				fmt.Printf("\nGRP0 SET\n")
// 			}
// 			DrawP0 = true
// 		}
//
// 		if Memory[GRP1] != 0 {
// 			if Debug {
// 				fmt.Printf("\nGRP1 SET\n")
// 			}
// 			DrawP1 = true
// 		}
//
// 	}
//
// 	if memAddr == uint16(RESP0) {
// 		if Memory[RESP0] != 0 {
// 			XPositionP0 = Beam_index
// 			if Debug {
// 				fmt.Printf("\nRESP0 SET\tXPositionP0: %d\n", XPositionP0)
// 			}
// 		}
// 	}
//
// 	if memAddr == uint16(RESP1) {
// 		if Memory[RESP1] != 0 {
// 			XPositionP1 = Beam_index
// 			if Debug {
// 				fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
// 			}
//
// 		}
// 	}
//
//
// 	if memAddr == uint16(HMP0) {
// 		XFinePositionP0 = Fine(Memory[HMP0])
//
// 		if Debug {
// 			fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
// 		}
//
// 	}
//
// 	if memAddr == uint16(HMP1) {
// 		XFinePositionP1 = Fine(Memory[HMP1])
// 		if Debug {
// 			fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
// 		}
// 	}
// }
