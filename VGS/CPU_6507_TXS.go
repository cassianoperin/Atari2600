package VGS

import	"fmt"

// TXS  Transfer Index X to Stack Register
//
//      X -> SP                          N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TXS           9A    1     2
func opc_TXS(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     TXS      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		SP = X

		if Debug {
			fmt.Printf("\tOpcode %02X [1 bytes] [Mode: Implied]\tTXS  Transfer Index X to Stack Pointer.\tSP = X (%d)\n", opcode, SP)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
