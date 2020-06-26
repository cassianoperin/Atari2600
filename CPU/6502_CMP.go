package CPU

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
func opc_CMP(memAddr uint16, mode string) {

	tmp := A - Memory[memAddr]

	if Debug {
		if tmp == 0 {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
		} else {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
		}
	}
	flags_Z(tmp)
	flags_N(tmp)
	flags_C_Subtraction(A,Memory[memAddr])

	// If mode=zeropage
	if Opcode == 0xC5 {
		Beam_index += 3
	// If mode=immediate
	} else if Opcode == 0xC9 {
		Beam_index += 2
	}

	PC += 2
}
