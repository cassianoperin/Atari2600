package VGS

import	"fmt"

// CLD  Clear Decimal Mode
//
//      0 -> D                           N Z C I D V
//                                       - - - - 0 -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLD           D8    1     2
func opc_CLD(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		P[3]	=  0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tCLD  Clear Decimal Mode.\tP[3]=%d\n", opcode, P[3])
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("CLD", bytes, opc_cycle_count + opc_cycle_extra)

						dbg_running_opc = false
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
