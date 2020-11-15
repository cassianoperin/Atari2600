package VGS

import	"fmt"

// PHA  Push Accumulator on Stack
//
//      push A                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PHA           48    1     3
func opc_PHA(bytes uint16, opc_cycles byte) {

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

		Memory[SP] = A

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPHA  Push Accumulator on Stack.\tMemory[%02X] = A (%d) | SP--\n", opcode, SP, Memory[SP])
			println(dbg_show_message)

			dbg_opcode_message("PHA", bytes, opc_cycle_count + opc_cycle_extra)
		}

		SP--

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
