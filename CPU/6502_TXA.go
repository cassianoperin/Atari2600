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
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tTXA  Transfer Index X to Accumulator.\tA = X (%d)\n", Opcode, X)
	}
	PC += 1
	flags_Z(A)
	flags_N(A)
	Beam_index += 2
}
