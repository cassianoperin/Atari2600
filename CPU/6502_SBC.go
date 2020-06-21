package CPU

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
func opc_SBC(memAddr uint16, mode string) {

	// Inverted Carry
	var borrow byte = P[0] ^ 1

	tmp := A

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow (zeropage).\tA = A(%d) - Memory[Memory[%02X]](%d) - (Borrow(Inverted Carry)(%d)) = %d\n", Opcode, Memory[PC+1], mode, A, PC+1, Memory[memAddr], borrow , A - Memory[memAddr] - borrow )
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

	// If mode=zeropage
	if Opcode == 0xE5 {
		PC += 2
		Beam_index += 3
	// If mode=immediate
	} else if Opcode == 0xE9 {
		PC += 2
		Beam_index += 2
	}


}
