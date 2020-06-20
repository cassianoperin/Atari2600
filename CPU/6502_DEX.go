package CPU

import	"fmt"

// DEX  Decrement Index X by One
//
//      X - 1 -> X                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       DEC           CA    1     2
func opc_DEX() {
	X --
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tDEX  Decrement Index X by One.\tX-- (%d)\n", Opcode, X)
	}
	PC += 1
	flags_Z(X)
	flags_N(X)
	Beam_index += 2
}
