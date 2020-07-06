package CPU

import	"fmt"

// CPX  Compare Memory and Index X
//
//      X - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPX #oper     E0    2     2
func opc_CPX(memAddr uint16, mode string) {

	tmp := X - Memory[memAddr]

	if Debug {
		if tmp == 0 {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
		} else {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
		}
	}

	flags_Z(tmp)
	flags_N(tmp)
	flags_C(X,Memory[memAddr])

	// If mode=immediate
	if Opcode == 0xE0 {
		PC += 2
		Beam_index += 2
	}
}
