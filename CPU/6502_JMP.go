package CPU

import	"fmt"

// JMP  Jump to New Location (absolute)
//
//      (PC+1) -> PCL                    N Z C I D V
//      (PC+2) -> PCH                    - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JMP oper      4C    3     3
func opc_JMP(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	Beam_index ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJMP  Jump to New Location.\t\tPC = 0x%04X\n", Opcode, Memory[PC+2], Memory[PC+1], mode, memAddr)
		}

		// Update PC
		PC = memAddr

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
