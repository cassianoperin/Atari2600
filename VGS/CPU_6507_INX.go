package VGS

import	"fmt"

// INX  Increment Index X by One
//
//      X + 1 -> X                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       INX           E8    1     2
func opc_INX(bytes uint16, opc_cycles byte) {

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

		X ++

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [%d byte] [Mode: Implied]\tINX  Increment Index X by One (%02X)\n", opcode, bytes, X)
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("INX", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
