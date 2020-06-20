package CPU

import	"fmt"

// DEY  Decrement Index Y by One
//
//      Y - 1 -> Y                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       DEC           88    1     2
func opc_DEY() {
	Y --
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tDEY  Decrement Index Y by One.\tY-- (%d)\n", Opcode, Y)
	}
	PC += 1
	flags_Z(Y)
	flags_N(Y)
	Beam_index += 2
}
