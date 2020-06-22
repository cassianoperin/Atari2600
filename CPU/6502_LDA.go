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
//      zeropage      LDA oper      A5    2     3
//      absolute,Y    LDA oper,Y    B9    3     4*
//      (indirect),Y  LDA (oper),Y  B1    2     5*
func opc_LDA(memAddr uint16, mode string) {
	A = Memory[memAddr]

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDA  Load Accumulator with Memory.\tA = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], mode, memAddr, A)
	}

	flags_Z(A)
	flags_N(A)

	// if mode == "Immediate"
	if Opcode == 0xA9 {
		PC += 2
		Beam_index += 2

	// if mode == "Zeropage"
	} else if Opcode == 0xA5 {
		PC += 2
		Beam_index += 3

	// if mode == "Absolute,Y"
	} else if Opcode == 0xB9 {
		PC += 3
		// Add 1 to cycles if page boundery is crossed
		if MemPageBoundary(memAddr, PC) {
			Beam_index += 1
		}
		Beam_index += 4
	} else if Opcode == 0xB1 {
		PC += 2
		// Add 1 to cycles if page boundery is crossed
		if MemPageBoundary(memAddr, PC) {
			Beam_index += 1
		}
		Beam_index += 5
	}
}
