package VGS

import	"fmt"

// BRK  Force Break
//
//      interrupt,                       N Z C I D V
//      push PC+2, push SR               - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       BRK           00    1     7
func opc_BRK(bytes uint16, opc_cycles byte) {

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

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tBRK  Force Break.\tPC = %04X\n", opcode, uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE]))
			println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("BRK", bytes, opc_cycle_count + opc_cycle_extra)

		}
		// IRQ Enabled
		P[2] = 1

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		PC = uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE])

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
