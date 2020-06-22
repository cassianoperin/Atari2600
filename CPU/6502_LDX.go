package CPU

import	"fmt"

// LDX  Load Index X with Memory
//
//      M -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDX #oper     A2    2     2
func opc_LDX(memAddr uint16, mode string) {
	X = Memory[memAddr]

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], mode, PC+1, X)
	}

	flags_Z(X)
	flags_N(X)

	PC += 2
	Beam_index += 2
}
