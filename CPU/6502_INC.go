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
func opc_INC(memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tINC  Increment Memory[%02X](%d) by One (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr] + 1)
	}

	Memory[ memAddr ] += 1

	flags_Z(Memory[ Memory[PC+1] ])
	flags_N(Memory[ Memory[PC+1] ])
	PC += 2
	Beam_index += 5
}
