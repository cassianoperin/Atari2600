package CPU

import	"fmt"

// LDY  Load Index Y with Memory (immediate)
//
//      M -> Y                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDY #oper     A0    2     2
//      zeropage      LDY oper      A4    2     3
func opc_LDY(memAddr uint16, mode string) {
	Y = Memory[memAddr]
	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDY  Load Index y with Memory.\tY = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], mode, PC+1, Y)
	}

	flags_Z(Y)
	flags_N(Y)

	// if mode == "Immediate"
	if Opcode == 0xA0 {
		PC += 2
		Beam_index += 2
	// if mode == "Zeropage"
	} else if Opcode == 0xA4 {
		PC += 2
		Beam_index += 3
	}
}
