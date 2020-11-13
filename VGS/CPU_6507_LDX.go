package VGS

import	"fmt"

// LDX  Load Index X with Memory
//
//      M -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDX #oper     A2    2     2
func opc_LDX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Check for extra cycles (*) in the first opcode cycle
	// if opc_cycle_count == 1 {
	// 	if Opcode == 0xBE {
	// 		// Add 1 to cycles if page boundery is crossed
	// 		if MemPageBoundary(memAddr, PC) {
	// 			opc_cycle_extra = 1
	// 		}
	// 	}
	// }

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra, opc_cycles, opc_cycle_extra)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     LDX      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
			dbg_opc_payload1	= Memory[PC+1]
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles +  opc_cycle_extra {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		X = Memory[memAddr]

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[%02X] (%d)\n", opcode, Memory[PC+1], mode, PC+1, X)
		}

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
