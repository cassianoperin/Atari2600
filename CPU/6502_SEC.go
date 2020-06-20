package CPU

import	"fmt"

// SEC  Set Carry Flag
//
//      1 -> C                           N Z C I D V
//                                       - - 1 - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEC           38    1     2
func opc_SEC() {
	P[0]	= 1
	PC += 1
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tSEC  Set Carry Flag.\tP[0]=1\n", Opcode)
	}
	Beam_index += 2
}
