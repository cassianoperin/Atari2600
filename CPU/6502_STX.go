package CPU

import	"fmt"

// STX  Store Index X in Memory (zeropage)
//
//      X -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STX oper      86    2     3
func opc_STX(offset byte, memAddr uint16, mode string) {
	Memory[Memory[PC+1]] = X

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[%02X] = X (%d)\n", Opcode, Memory[PC+1], mode, memAddr, X)
	}

	PC += 2
	Beam_index += 3


	// Memory[Memory[PC+1]] = X
	// if Debug {
	// 	fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[%02X] = X (%d)\n", Opcode, Memory[PC+1], mode, memAddr, X)
	// }
	//
	// PC += 2
	// Beam_index += 3

}
