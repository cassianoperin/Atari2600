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

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     BRK      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tBRK  Force Break.\tPC = %04X\n", opcode, uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE]))
		}
		// IRQ Enabled
		P[2] = 1

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		PC = uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE])

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
