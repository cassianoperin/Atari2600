package VGS

import	"fmt"

// CLC  Clear Carry Flag
//
//      0 -> C                           N Z C I D V
//                                       - - 0 - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLC           18    1     2
func opc_CLC(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

		// Reset to default value
		TIA_Update = -1

	// After spending the cycles needed, execute the opcode
	} else {

		P[0] = 0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tCLC  Clear Carry Flag.\tP[0]=0\n", opcode)
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("CLC", bytes, opc_cycle_count + opc_cycle_extra)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
