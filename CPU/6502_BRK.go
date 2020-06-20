package CPU

import	"fmt"

// BRK  Force Break
//
//      interrupt,                       N Z C I D V
//      push PC+2, push SR               - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       BRK           00    1     7
func opc_BRK() {
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tBRK  Force Break.\tBREAK!\n", Opcode)
	}
	// IRQ Enabled
	P[2] = 1
	Break()
	Beam_index += 7
}
