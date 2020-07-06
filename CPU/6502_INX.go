package CPU

import	"fmt"

// INX  Increment Index X by One
//
//      X + 1 -> X                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       INX           E8    1     2
func opc_INX() {

	X ++

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tINX  Increment Index X by One (%02X)\n", Opcode, X)
	}

	flags_Z(X)
	flags_N(X)

	PC++
	Beam_index += 2
}
