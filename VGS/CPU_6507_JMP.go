package VGS

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
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJMP  Jump to New Location.\t\tPC = 0x%04X\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr)
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("JMP", bytes, opc_cycle_count + opc_cycle_extra)
		}

		// Update PC
		PC = memAddr

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
