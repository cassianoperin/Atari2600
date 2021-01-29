package VGS

import	"fmt"

// EOR  Exclusive-OR Memory with Accumulator (immidiate)
//
//      A EOR M -> A                     N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immidiate     EOR #oper     49    2     2
//      zeropage      EOR oper      45    2     3
func opc_EOR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

	// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tEOR  Exclusive-OR Memory with Accumulator.\tA = A(%d) XOR Memory[%02X](%d)\t(%d)\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A ^ Memory[memAddr] )
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("EOR", bytes, opc_cycle_count + opc_cycle_extra)
		}

		A = A ^ Memory[memAddr]

		flags_Z(A)
		flags_N(A)


		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}
}
