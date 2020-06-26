package CPU

import	"fmt"

// RTS  Return from Subroutine
//
//      pull PC, PC+1 -> PC              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTS           60    1     6
func opc_RTS() {
	PC = uint16(Memory[SP+2])<<8 | uint16(Memory[SP+1])

	// Clear the addresses retrieved from the stack
	Memory[SP+1] = 0
	Memory[SP+2] = 0
	// Update the Stack Pointer (Increase the two values retrieved)
	SP += 2

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tRTS  Return from Subroutine.\tPC = %04X.\n", Opcode, PC)
	}

	PC += 1
	Beam_index += 6
}
