package VGS

import	"fmt"

// STY  Store Index Y in Memory (zeropage)
//
//      Y -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STY oper      84    2     3
func opc_STY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Change variable to a positive number to TIA detect the change
		if memAddr < 128 {
			TIA_Update = int8(memAddr)	// Change variable to a positive number to TIA detect the change
		}

		Memory[memAddr] = Y

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTY  Store Index Y in Memory.\tMemory[%02X] = Y (%d)\n", opcode, Memory[PC+1], mode, memAddr, Y)

			// Collect data for debug interface just on first cycle
			if opc_cycle_count == 1 {
				debug_opc_text		= fmt.Sprintf("%04x     STY      ;%d", PC, opc_cycles)
				dbg_opc_bytes		= bytes
				dbg_opc_opcode		= opcode
				dbg_opc_payload1	= Memory[PC+1]
			}
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
