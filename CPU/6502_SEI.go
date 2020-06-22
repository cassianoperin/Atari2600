package CPU

import	"fmt"

// SEI  Set Interrupt Disable Status
//
//      1 -> I                           N Z C I D V
//                                       - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEI           78    1     2
func opc_SEI() {
	P[2]	=  1

	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tSEI  Set Interrupt Disable Status.\tP[2]=%d\n", Opcode, P[2])
	}

	PC += 1
	Beam_index += 2
}
