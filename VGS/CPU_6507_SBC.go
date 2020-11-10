package VGS

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

	// After spending the cycles needed, execute the opcode
	} else {

		tmp := A

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+1], mode, A, PC+1, Memory[memAddr], borrow , A - Memory[memAddr] - borrow )
		}

		// Result
		A = A - Memory[memAddr] - borrow

		// For the flags:
		// The subtraction is VALUE1 (A) - VALUE2 (Memory[PC+1] - (P[0] ^ 1)
		value2 := Memory[PC+1] - borrow

		// First V because it need the original carry flag value
		Flags_V_SBC(tmp, value2)
		// After, update the carry flag value
		flags_C_Subtraction(tmp, value2)

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
	}

}
