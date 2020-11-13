package VGS

import	"fmt"

// TXA  Transfer Index X to Accumulator
//
//      X -> A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TXA           8A    1     2
func opc_TXA(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     TXA      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		A = X

		if Debug {
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tTXA  Transfer Index X to Accumulator.\tA = X (%d)\n", opcode, X)
		}

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
