package CPU

import	"fmt"

// LDA  Load Accumulator with Memory (immidiate)
//
//      M -> A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDA #oper     A9    2     2
func opc_LDA(offset byte, memAddr uint16, mode string) {
	A = offset
	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDA  Load Accumulator with Memory.\tA = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], mode, memAddr, A)
	}

	flags_Z(A)
	flags_N(A)
	PC += 2
	Beam_index += 2
}
