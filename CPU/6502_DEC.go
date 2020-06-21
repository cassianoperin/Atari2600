package CPU

import	"fmt"

// DEC  Decrement Memory by One
//
//      M - 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      DEC oper      C6    2     5
func opc_DEC(memAddr uint16, mode string) {

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tDEC  Decrement Memory by One.\tMemory[%02X] -= 1 (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] - 1 )
	}
	Memory[memAddr] -= 1

	flags_Z(Memory[memAddr])
	flags_N(Memory[memAddr])

	PC += 2
	Beam_index += 5
}
