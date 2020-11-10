package VGS

import	"fmt"

// TAX  Transfer Accumulator to Index X
//
//      A -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TAX           AA    1     2
func opc_TAX(bytes uint16, opc_cycles byte) {

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

		X = A

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tTAX  Transfer Accumulator to Index X.\tX = A (%d)\n", opcode, A)
		}

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
