package VGS

import	"os"
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

		var tmp byte

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			// FIRST ATTEMPT TO DETECT ACCESS TO A TIA READ ONLY REGISTER (0x00-0x0D)
			if memAddr < 14 {
				tmp = A - Memory_TIA_RO[memAddr]
			// Read from other reserved TIA registers
			} else if memAddr < 128 {
				fmt.Printf("CMP - Controlled Exit to map access to TIA Write Addresses. COULD BE MIRRORS!!!!!.\t EXITING\n")
				os.Exit(2)
			// Read from RIOT Memory Map (> 0x280)
			} else {
				tmp = A - Memory[memAddr]
			}
		// 6507 interpreter mode
		} else {
			tmp = A - Memory[memAddr]
		}

		if Debug {
			// Access to TIA RO Memory
			if memAddr < 14 {
				if tmp == 0 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - TIA_RO_Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory_TIA_RO[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - TIA_RO_Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory_TIA_RO[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}
			// Access to the rest of memory
			} else {
				if tmp == 0 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("CMP", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(tmp)
		flags_N(tmp)
		// Access to TIA RO Memory
		if memAddr < 14 {
			flags_C_Subtraction(A, Memory_TIA_RO[memAddr])
		// Access to the rest of memory
		} else {
			flags_C_Subtraction(A ,Memory[memAddr])
		}
		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
