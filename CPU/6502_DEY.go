package CPU

import	"fmt"

// DEY  Decrement Index Y by One
//
//      Y - 1 -> Y                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       DEC           88    1     2
func opc_DEY(bytes uint16, opc_cycles byte) {

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

		Y --

		if Debug {
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tDEY  Decrement Index Y by One.\tY-- (%d)\n", Opcode, Y)
		}

		flags_Z(Y)
		flags_N(Y)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
