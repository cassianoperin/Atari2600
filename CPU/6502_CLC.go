package CPU

import	"fmt"

// CLC  Clear Carry Flag
//
//      0 -> C                           N Z C I D V
//                                       - - 0 - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLC           18    1     2
func opc_CLC() {

	P[0]	= 0

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tCLC  Clear Carry Flag.\tP[0]=0\n", Opcode)
	}

	PC += 1
	Beam_index += 2
}
