package CPU

import	"fmt"

// CPY  Compare Memory and Index Y
//
//      Y - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPY #oper     C0    2     2
//      zeropage      CPY oper      C4    2     3
func opc_CPY(memAddr uint16, mode string) {

	tmp := Y - Memory[memAddr]

	if Debug {
		if tmp == 0 {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], mode, Y, PC+1, Memory[memAddr], tmp)
		} else {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], mode, Y, PC+1, Memory[memAddr], tmp)
		}
	}

	flags_Z(tmp)
	flags_N(tmp)
	flags_C(Y,Memory[memAddr])

	// If mode=immediate
	if Opcode == 0xC0 {
		PC += 2
		Beam_index += 2
	// If mode=zeropage
	} else if Opcode == 0xC4 {
		PC += 2
		Beam_index += 3
	}
}
