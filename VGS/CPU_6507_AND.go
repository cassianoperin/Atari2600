package VGS

import	"os"
import	"fmt"

// AND  AND Memory with Accumulator
//
//      A AND M -> A                     N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immidiate     AND #oper     29    2     2
//      zeropage	    AND oper    	25  	2   	3
func opc_AND(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
	if memAddr < 0x80 {
		fmt.Printf("AND - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
		os.Exit(2)
	}

	// Increment the beam
	beamIndex ++

	// // Check for extra cycles (*) in the first opcode cycle
	// if opc_cycle_count == 1 {
	// 	if Opcode == 0xB9 || Opcode == 0xBD || Opcode == 0xB1 {
	// 		// Add 1 to cycles if page boundery is crossed
	// 		if MemPageBoundary(memAddr, PC) {
	// 			opc_cycle_extra = 1
	// 		}
	// 	}
	// }

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles +  opc_cycle_extra {
		opc_cycle_count ++

		// Reset to default value
		TIA_Update = -1

	// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A & Memory[memAddr] )
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("AND", bytes, opc_cycle_count + opc_cycle_extra)
		}

		A = A & Memory[memAddr]

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
