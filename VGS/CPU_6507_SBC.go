package VGS

import	"os"
import	"fmt"

// SBC  Subtract Memory from Accumulator with Borrow (zeropage)
//
//      A - M - C -> A                   N Z C I D V
//                                       + + + - - +
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      SBC oper      E5    2     3
//      immediate     SBC #oper     E9    2     2
func opc_SBC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Atari 2600 interpreter mode
	if CPU_MODE == 0 {
		// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
		if memAddr < 0x80 {
			fmt.Printf("SBC - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
			os.Exit(2)
		}

		// Some tests of instructions that tryes to read from RIOT addresses (640 - 671)
		if memAddr > 0x280 &&  memAddr <= 0x29F {
			fmt.Printf("SBC - Tryed to read from RIOT ADDRESS! Memory[%X]\tEXIT\n", memAddr)
			os.Exit(2)
		}
	}

	// Inverted Carry
	var borrow byte = P[0] ^ 1

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

		original_A := A

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			if Debug {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], borrow , A - Memory[memAddr] - borrow )
				fmt.Println(dbg_show_message)

				// Collect data for debug interface after finished running the opcode
				dbg_opcode_message("SBC", bytes, opc_cycle_count + opc_cycle_extra)
			}

			// Result
			A = A - Memory[memAddr] - borrow

			// For the flags:
			// The subtraction is VALUE1 (A) - VALUE2 (Memory[PC+1] - (P[0] ^ 1)
			// value2 := Memory[PC+1] - borrow

			// First V because it need the original carry flag value
			Flags_V_SBC(original_A, Memory[memAddr])
			// After, update the carry flag value
			flags_C_Subtraction(original_A, A)

			// // Clear Carry if overflow in bit 7 // NOT NECESSARY
			// if P[6] == 1 {
			// 	fmt.Printf("\n\tCarry cleared due to an overflow!")
			// 	P[0] = 0
			// }

			flags_Z(A)
			flags_N(A)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1

			// Reset Opcode Extra Cycle counter
			opc_cycle_extra = 0

		// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {
			fmt.Println("SBC DECIMAL NOT INPLEMENTED YET! EXITING")
			os.Exit(2)
		}
	}

}
