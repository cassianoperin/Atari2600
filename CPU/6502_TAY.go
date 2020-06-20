package CPU

import	"fmt"

// TAY  Transfer Accumulator to Index Y
//
//      A -> Y                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TAY           A8    1     2
func opc_TAY() {
	Y = A
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tTAY  Transfer Accumulator to Index Y.\tA = X (%d)\n", Opcode, X)
	}
	PC += 1
	flags_Z(Y)
	flags_N(Y)
	Beam_index += 2
}
