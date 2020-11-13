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

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     PLA      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		A = Memory[SP+1]

		// Not documented, clean the value on the stack after pull it to accumulator
		Memory[SP+1] = 0

		if Debug {
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[%02X] (%d) | SP++\n", opcode, SP, A )
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
