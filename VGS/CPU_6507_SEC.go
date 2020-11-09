package VGS

import	"fmt"

// SEC  Set Carry Flag
//
//      1 -> C                           N Z C I D V
//                                       - - 1 - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEC           38    1     2
func opc_SEC(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		P[0] = 1

		if Debug {
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tSEC  Set Carry Flag.\tP[0]=1\n", opcode)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
