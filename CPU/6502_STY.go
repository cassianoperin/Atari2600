package CPU

import	"fmt"

// STY  Store Index Y in Memory (zeropage)
//
//      Y -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STY oper      84    2     3
func opc_STY(memAddr uint16, mode string) {

	TIA_Update = int8(memAddr)	// Change variable to a positive number to TIA detect the change

	Memory[memAddr] = Y

	if Debug {
		fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTY  Store Index Y in Memory.\tMemory[%02X] = Y (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Y)
	}

	PC += 2
	Beam_index += 3

}
