package VGS

import	"fmt"

// TAY  Transfer Accumulator to Index Y
//
//      A -> Y                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TAY           A8    1     2
func opc_TAY(bytes uint16, opc_cycles byte) {

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

		Y = A

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tTAY  Transfer Accumulator to Index Y.\tY = A (%d)\n", opcode, A)
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("TAY", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(Y)
		flags_N(Y)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
