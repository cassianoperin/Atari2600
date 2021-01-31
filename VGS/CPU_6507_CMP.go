package VGS

import	"fmt"

// CMP  Compare Memory with Accumulator
//
//      A - M                          N Z C I D V
//                                     + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      CMP oper      C5    2     3
//      immediate     CMP #oper     C9    2     2
func opc_CMP(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// // Check for extra cycles (*) in the first opcode cycle
	// if opc_cycle_count == 1 {
	// 	if Opcode == 0xB9 || Opcode == 0xBD || Opcode == 0xB1 {
	// 		// Add 1 to cycles if page boundery is crossed
	// 		if MemPageBoundary(memAddr, PC) {
	// 			opc_cycle_extra = 1
	// 		}
	// 	}
	// }

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles +  opc_cycle_extra {
		opc_cycle_count ++

		// Reset to default value
		TIA_Update = -1

	// After spending the cycles needed, execute the opcode
	} else {

		tmp := A - Memory[memAddr]

		if Debug {
			if tmp == 0 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
				fmt.Println(dbg_show_message)
			} else {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
				fmt.Println(dbg_show_message)
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("CMP", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(tmp)
		flags_N(tmp)
		flags_C_Subtraction(A,Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
