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
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tBRK  Force Break.\tPC = %04X\n", Opcode, uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE]))
	}
	// IRQ Enabled
	P[2] = 1

	// Read the Opcode from PC+1 and PC bytes (Little Endian)
	PC = uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE])

	Beam_index += 7
}
