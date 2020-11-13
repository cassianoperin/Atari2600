package VGS

import	"fmt"

// RTS  Return from Subroutine
//
//      pull PC, PC+1 -> PC              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTS           60    1     6
func opc_RTS(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     RTS      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		PC = uint16(Memory[SP+2])<<8 | uint16(Memory[SP+1])

		// Clear the addresses retrieved from the stack
		Memory[SP+1] = 0
		Memory[SP+2] = 0
		// Update the Stack Pointer (Increase the two values retrieved)
		SP += 2

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tRTS  Return from Subroutine.\tPC = %04X.\n", opcode, PC)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
