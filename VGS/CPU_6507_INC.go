package VGS

import	"fmt"

// INC  Increment Memory by One
//
//      M + 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      INC oper      E6    2     5
func opc_INC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     INC      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
			dbg_opc_payload1	= Memory[PC+1]
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tINC  Increment Memory[%02X](%d) by One (%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr] + 1)
		}

		Memory[ memAddr ] += 1

		flags_Z(Memory[ memAddr ])
		flags_N(Memory[ memAddr ])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
