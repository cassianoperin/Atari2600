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

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     CLC      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		P[0] = 0

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tCLC  Clear Carry Flag.\tP[0]=0\n", opcode)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
