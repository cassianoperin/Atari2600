package CPU

import	"fmt"

// ORA  OR Memory with Accumulator
//
//      A OR M -> A                      N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      ORA oper      05    2     3
func opc_ORA(memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tORA  OR Memory with Accumulator.\tA = A(%d) | Memory[%02X](%d)\t(%d)\n", Opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A | Memory[memAddr] )
	}

	A = A | Memory[memAddr]

	flags_Z(A)
	flags_N(A)

	PC += 2
	Beam_index += 3
}
