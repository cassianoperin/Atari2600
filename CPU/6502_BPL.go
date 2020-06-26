package CPU

import	"fmt"

// BPL  Branch on Result Plus
//
//      branch on N = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BPL oper      10    2     2**
func opc_BPL(value int8) {
	// If Positive
	if P[7] == 0 {

		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBranch on Result POSITIVE.\tCarry EQUAL 1, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+uint16(value) )
		}
		// Current PC (To detect page bounday cross)
		tmp := PC
		// fmt.Printf("\ntmp: %02X\n",tmp)

		// PC + the number of bytes to jump on carry clear
		PC += uint16(value)

		// Increment PC
		PC += 2

		// Add 1 to cycles if branch occurs on same page
		Beam_index += 1

		// // Add one extra cycle if branch occurs in a differente memory page
		if MemPageBoundary(uint16(tmp), PC) {
			Beam_index += 1
		}

	// If not negative
	} else {
		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes]\tBranch on Result POSITIVE.\t\tNEGATIVE Flag enabled, PC+=2\n", Opcode, Memory[PC+1])
		}
		PC += 2
	}

	Beam_index += 2
}
