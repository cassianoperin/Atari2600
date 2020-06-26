package CPU

import	"fmt"

// PHA  Push Accumulator on Stack
//
//      push A                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PHA           48    1     3
func opc_PHA() {
	Memory[SP] = A

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tPHA  Push Accumulator on Stack.\tMemory[%02X] = A (%d) | SP--\n", Opcode, SP, Memory[SP])
	}

	PC += 1
	SP--
	Beam_index += 3
}
