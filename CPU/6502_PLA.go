package CPU

import	"fmt"

// PLA  Pull Accumulator from Stack
//
//      pull A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLA           68    1     4
func opc_PLA() {
	A = Memory[SP+1]

	// Not documented, clean the value on the stack after pull it to accumulator
	Memory[SP+1] = 0

	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[%02X] (%d) | SP++\n", Opcode, SP, A )
	}
	PC += 1
	SP++
	Beam_index += 4

	flags_N(A)
	flags_Z(A)
}
