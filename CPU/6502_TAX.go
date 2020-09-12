package CPU

import	"fmt"

// TAX  Transfer Accumulator to Index X
//
//      A -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TAX           AA    1     2
func opc_TAX() {
	X = A

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tTAX  Transfer Accumulator to Index X.\tX = A (%d)\n", Opcode, A)
	}

	flags_Z(X)
	flags_N(X)

	PC += 1
	Beam_index += 2
}
