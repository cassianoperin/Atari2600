package CPU

import	"fmt"

// AND  AND Memory with Accumulator
//
//      A AND M -> A                     N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immidiate     AND #oper     29    2     2
func opc_AND(memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", Opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A & Memory[memAddr] )
	}

	A = A & Memory[memAddr]

	flags_Z(A)
	flags_N(A)

	PC += 2
	Beam_index += 2
}
