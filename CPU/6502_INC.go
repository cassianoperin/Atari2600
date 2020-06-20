package CPU

import	"fmt"

// INC  Increment Memory by One
//
//      M + 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      INC oper      E6    2     5
func opc_INC(offset byte, memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tINC  Increment Memory[%02X](%d) by One (%d)\n", Opcode, Memory[PC+1], mode, Memory[PC+1], offset, offset + 1)
	}

	Memory[ Memory[PC+1] ] = offset + 1

	flags_Z(Memory[ Memory[PC+1] ])
	flags_N(Memory[ Memory[PC+1] ])
	PC += 2
	Beam_index += 5
}
