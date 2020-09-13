package CPU

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
	Beam_index ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		P[3]	=  0

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tCLD  Clear Decimal Mode.\tP[3]=%d\n", Opcode, P[3])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
