package CPU

import	"fmt"

// INY  Increment Index Y by One
//
//      Y + 1 -> Y                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       INY           C8    1     2
func opc_INY() {
	Y ++
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tINY  Increment Index Y by One (%02X)\n", Opcode, Y)
	}
	flags_Z(Y)
	flags_N(Y)
	PC++
	Beam_index += 2
}
