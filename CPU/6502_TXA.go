package CPU

import	"fmt"

// TXA  Transfer Index X to Accumulator
//
//      X -> A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TXA           8A    1     2
func opc_TXA() {
	A = X

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tTXA  Transfer Index X to Accumulator.\tA = X (%d)\n", Opcode, X)
	}

	flags_Z(A)
	flags_N(A)

	PC += 1
	Beam_index += 2
}
