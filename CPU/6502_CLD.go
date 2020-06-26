package CPU

import	"fmt"

// CLD  Clear Decimal Mode
//
//      0 -> D                           N Z C I D V
//                                       - - - - 0 -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLD           D8    1     2
func opc_CLD() {
	P[3]	=  0

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tCLD  Clear Decimal Mode.\tP[3]=%d\n", Opcode, P[3])
	}

	PC += 1
	Beam_index += 2
}
