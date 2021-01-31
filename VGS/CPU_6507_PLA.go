package VGS

import	"fmt"

// PLA  Pull Accumulator from Stack
//
//      pull A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLA           68    1     4
func opc_PLA(bytes uint16, opc_cycles byte) {

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

		A = Memory[SP+1]

		// Not documented, clean the value on the stack after pull it to accumulator
		Memory[SP+1] = 0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[%02X] (%d) | SP++\n", opcode, SP, A )
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("PLA", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_N(A)
		flags_Z(A)

		SP++

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
