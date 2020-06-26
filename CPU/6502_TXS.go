package CPU

import	"fmt"

// TXS  Transfer Index X to Stack Register
//
//      X -> SP                          N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TXS           9A    1     2
func opc_TXS() {
	SP = X

	if Debug {
		fmt.Printf("\tOpcode %02X [1 bytes] [Mode: Implied]\tTXS  Transfer Index X to Stack Pointer.\tSP = X (%d)\n", Opcode, SP)
	}

	PC += 1
	Beam_index += 2
}
